package ssh

import (
	"context"
	"io"
	"net"
	"sectran/common/constants"
	"strconv"

	"github.com/sirupsen/logrus"
)

type SSHConnRequest struct {
	ReqProtocolType constants.REQ_PROTOCOL_TYPE //user terminal protocol
	Config          *SSHConfig                  //terget server config
	Conn            io.ReadWriteCloser          //net connection
}

type SSHConnResponse struct {
	err error
}

type SSHModuleMessage struct {
	RequestChan  chan *SSHConnRequest
	ResponseChan chan *SSHConnResponse
	Cancle       context.CancelFunc
}

// global ssh service config
var SSHModuleConfig SSHConfig

// chan SSHConnRequest: request a proxy
// chan SSHConnResponse: response for request channle
// CancelFunc: if want to stop ssh service gracefully
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

	//start tcp service
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
		err  error
		req  *SSHConnRequest
		rwcc io.ReadWriteCloser
		rwcs io.ReadWriteCloser
	)

	for {
		select {
		case req = <-message.RequestChan:
			rwcs = req.Conn
			rwcc, err = NewSSHClient(req.Config)
			if err != nil {
				message.ResponseChan <- &SSHConnResponse{
					err: err,
				}
				break
			}

			go handleReveredClientConnection(rwcc, rwcs)
			go handleReveredServerConnection(rwcs, rwcc)

		case <-ctx.Done():
			logrus.Infof("a connection is done")
			//clean
		}
	}

}

func handleReveredClientConnection(r io.ReadWriteCloser, w io.ReadWriteCloser) {
	// Just have a method to trigger it
	defer r.Close()
	defer w.Close()

	var stopper chan int = make(chan int, 1)
	reversedFunc(r, w, stopper)

	<-stopper
}

func handleReveredServerConnection(r io.ReadWriteCloser, w io.ReadWriteCloser) {
	var stopper chan int = make(chan int, 1)
	reversedFunc(r, w, stopper)

	<-stopper
}

func reversedFunc(r io.Reader, w io.Writer, stopper chan int) {
	close := func(stopper chan int) {
		stopper <- 1
	}
	defer close(stopper)

	var buffer []byte = make([]byte, 4096)
	for {
		n, err := r.Read(buffer)
		if err != nil {
			logrus.Errorf("read error :%s", err)
			return
		}

		if n > 0 {
			_, err = w.Write(buffer[:n])
			if err != nil {
				logrus.Errorf("write error :%s", err)
				return
			}
		}
	}
}

func startSSHTcpService(config *SSHConfig, netChan chan *SSHConnRequest, addr string) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatalf("error listen in addr %s.due to %s", addr, err)
	}
	defer l.Close()
	logrus.Infof("start ssh proxy tcp service with %s", addr)

	for {
		c, err := l.Accept()
		if err != nil {
			logrus.Errorf("error accept client due to %s", err)
		}

		clientConfig := *config
		rwc, err := NewSSHServer(c, &clientConfig)
		if err != nil {
			logrus.Errorf("error New SSH Server due to %s", err)
		}
		logrus.Infof("destnation number is:%s", clientConfig.Password)

		//todo:change client config to what you want here
		clientConfig.Host = "192.168.31.100"
		clientConfig.Port = 22
		clientConfig.PasswordAuth = true
		clientConfig.UserName = "root"
		clientConfig.Password = "Ryan@1218pass"

		netChan <- &SSHConnRequest{
			Conn:            rwc,
			Config:          &clientConfig,
			ReqProtocolType: constants.REQUEST_TCP,
		}
	}
}
