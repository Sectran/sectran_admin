package ssh

import (
	"context"
	"io"
	"net"
	"sectran/common/constants"
	"strconv"
	"unsafe"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

type SSHConnRequest struct {
	ReqProtocolType constants.REQ_PROTOCOL_TYPE // user terminal protocol
	Config          *SSHConfig                  // terget server config
	Conn            io.ReadWriteCloser          // net connection
	ConnType        uint8                       // pty/sftp
}

type SSHConnResponse struct {
	Err error
}

type SSHModuleMessage struct {
	RequestChan  chan *SSHConnRequest
	ResponseChan chan *SSHConnResponse
	Cancle       context.CancelFunc
}

// global ssh common config
var SSHModuleConfig SSHConfig

type handlePostRead func(data []byte, termianl unsafe.Pointer) bool

// chan SSHConnRequest: request a proxy
// chan SSHConnResponse: response for request channle
// CancelFunc: if want to stop ssh common gracefully
// error: some error
func StartSSHModule(config *SSHConfig) (*SSHModuleMessage, error) {
	if err := CheckSSHConfig(config); err != nil {
		return nil, err
	}

	ctx, cancle := context.WithCancel(context.Background())

	//reference to globle
	SSHModuleConfig = *config
	reqChan := make(chan *SSHConnRequest, 1)
	resChan := make(chan *SSHConnResponse, 1)

	//start tcp common
	go startSSHTcpService(config, reqChan, net.JoinHostPort(config.Host, strconv.Itoa(int(config.Port))))

	message := &SSHModuleMessage{
		Cancle:       cancle,
		RequestChan:  reqChan,
		ResponseChan: resChan,
	}
	go handleConnection(ctx, message)

	return message, nil
}

func handleConnection(ctx context.Context, message *SSHModuleMessage) {
	var (
		err       error
		req       *SSHConnRequest
		rwcc      io.ReadWriteCloser
		rwcs      io.ReadWriteCloser
		sshClient *SSHClient
	)

	for {
		select {
		case req = <-message.RequestChan:
			rwcs = req.Conn
			switch req.ConnType {
			case 0: //SSH
				sshClient, err = NewSSHClient(req.Config)
				if err != nil {
					message.ResponseChan <- &SSHConnResponse{
						Err: err,
					}
					req.Conn.Close()
					break
				}

				rwcc = sshClient.PtyChannel
				var terminal unsafe.Pointer
				terminal = XtermStart(int(req.Config.PtyRequestMsg.Columns), int(req.Config.PtyRequestMsg.Rows))

				peerPostReadCb := func(data []byte, termianl unsafe.Pointer) bool {
					// logrus.Infof("%q", data)
					if len(data) == 1 {
						switch data[0] {
						case '\r':
							command := XtermGetCommand(termianl)
							if len(command) > 0 {
								logrus.Infof("get current command :\n%s", command)
							}
						case 0x03:
							//just flush buffer
							_ = XtermGetCommand(termianl)
						default:
							XtermMarkStdin(termianl, data)
						}
					} else {
						XtermMarkStdin(termianl, data)
					}
					return true
				}

				clientPostReadCb := func(data []byte, termianl unsafe.Pointer) bool {
					XtermWrite(termianl, data)
					XtermDumpToFile(termianl)
					return true
				}

				go RevereConnection(rwcs, rwcc, peerPostReadCb, clientPostReadCb, terminal)
			case 1: //SFTP
				logrus.Errorf("sftp is not surpported now")
				req.Conn.Close()
			default:
				logrus.Errorf("unkone request type of %d in ssh channel", req.ConnType)
			}
		case <-ctx.Done():
			logrus.Infof("a connection is done")
		}

	}
}

func RevereConnection(peer, client io.ReadWriteCloser, peerPostRead handlePostRead,
	clientPostRead handlePostRead, termianl unsafe.Pointer) {
	errChan := make(chan error, 1)

	proxy := func(from io.ReadWriteCloser, to io.ReadWriteCloser, postRead handlePostRead,
		errChan chan error, termianl unsafe.Pointer) {
		var buffer []byte = make([]byte, 4096)
		for {
			if n, err := from.Read(buffer); err != nil {
				errChan <- err
				return
			} else if n > 0 {
				if postRead(buffer[:n], termianl) {
					if _, err := to.Write(buffer[:n]); err != nil {
						errChan <- err
						return
					}
				}
			}
		}
	}

	defer peer.Close()
	defer client.Close()

	go proxy(peer, client, peerPostRead, errChan, termianl)
	go proxy(client, peer, clientPostRead, errChan, termianl)

	err := <-errChan
	logrus.Errorf("proxy error due to :%s", err)
}

func startSSHTcpService(config *SSHConfig, netChan chan *SSHConnRequest, addr string) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatalf("error listen in addr %s.due to %s", addr, err)
	}
	defer l.Close()
	logrus.Infof("start ssh proxy tcp common with %s", addr)

	for {
		c, err := l.Accept()
		if err != nil {
			logrus.Errorf("error accept client due to %s", err)
		}

		go func(clientConfig *SSHConfig, netChan chan *SSHConnRequest) {
			channels, err := NewSSHServer(c, clientConfig)
			if err != nil {
				logrus.Errorf("error New SSH Server due to %s", err)
				return
			}

			for {
				select {
				case pty := <-channels.Pty:
					logrus.Infof("destnation number is:%s", clientConfig.Password)
					clientConfig.Host = pty.UserConf.Host
					clientConfig.Port = pty.UserConf.Port
					clientConfig.UserName = pty.UserConf.UserName
					clientConfig.Password = pty.UserConf.Password

					clientConfig.PasswordAuth = true

					netChan <- &SSHConnRequest{
						ConnType:        0,
						Conn:            pty.Data.(ssh.Channel),
						Config:          clientConfig,
						ReqProtocolType: constants.REQUEST_TCP,
					}
				case sftp := <-channels.Sftp:
					netChan <- &SSHConnRequest{
						ConnType:        1,
						Conn:            sftp.Data.(ssh.Channel),
						Config:          clientConfig,
						ReqProtocolType: constants.REQUEST_TCP,
					}
				case err := <-channels.Err:
					logrus.Errorf("waiting for SSH channel to encounter an error in a loop %s", err.Data)
					return
				}
			}
		}(config, netChan)
	}
}
