package ssh

import (
	"fmt"
	"io"
	"net"
	"os"
	"sectran/user/config"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

type PtyReqMsg struct {
	Term     string
	Columns  uint32
	Rows     uint32
	Width    uint32
	Height   uint32
	Modelist string
}

const (
	SectranSSHCVeriosn string = "SSH-2.0-Sectran"
)

func NewSSHClient(userConf *config.SSHConfig) (io.ReadWriteCloser, error) {
	if !config.CheckSSHConfig(userConf) {
		return nil, fmt.Errorf("invalid user SSH configuration")
	}

	var (
		auth    []ssh.AuthMethod
		err     error
		signer  ssh.Signer
		pri     []byte
		config  *ssh.ClientConfig
		client  *ssh.Client
		request <-chan *ssh.Request
		channel ssh.Channel
		// buffer    *bytes.Buffer
		// writer    *bufio.Writer
		modeList  []byte
		ptyReqMes PtyReqMsg
	)

	//InteractiveAuth and PasswordAuth is the same for client side
	if userConf.PasswordAuth || userConf.InteractiveAuth {
		auth = append(auth, ssh.Password(userConf.Password))
	} else if userConf.PublicKeyAuth {
		pri, err = os.ReadFile(userConf.PrivateKey)
		if err != nil {
			goto end
		}
		signer, err = ssh.ParsePrivateKey(pri)
		if err != nil {
			logrus.Errorf("could not parse private key:%s", err)
			goto end
		}
		auth = append(auth, ssh.PublicKeys(signer))
	}

	config = &ssh.ClientConfig{
		User:          userConf.UserName,
		ClientVersion: SectranSSHCVeriosn,
		// It seems that we don't need to verify the key from the server
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		// BannerCallback: func(message string) error {
		// 	return nil
		// },
		// maybe network is so bad
		Timeout: 10 * time.Second,
		Auth:    auth,
	}

	client, err = ssh.Dial("tcp", net.JoinHostPort(userConf.Host, strconv.Itoa(int(userConf.Port))), config)
	if err != nil {
		goto end
	}

	channel, request, err = client.Conn.OpenChannel("session", nil)
	if err != nil {
		goto end
	}
	go ssh.DiscardRequests(request)

	// buffer = &bytes.Buffer{}
	// writer = bufio.NewWriter(buffer)

	// // terminal length
	// binary.Write(writer, binary.BigEndian, uint32(len(userConf.PtyRequestMsg.Term)))
	// // terminal string
	// writer.WriteString(userConf.PtyRequestMsg.Term)

	// //rows colums width height
	// binary.Write(writer, binary.BigEndian, userConf.PtyRequestMsg.Rows)
	// binary.Write(writer, binary.BigEndian, userConf.PtyRequestMsg.Columns)
	// binary.Write(writer, binary.BigEndian, userConf.PtyRequestMsg.Width)
	// binary.Write(writer, binary.BigEndian, userConf.PtyRequestMsg.Height)

	// //mode list length
	// binary.Write(writer, binary.BigEndian, uint32(len(userConf.ModeList)))

	// //mode list begain
	// for _, v := range userConf.ModeList {
	// 	writer.WriteByte(v.Key)
	// 	binary.Write(writer, binary.BigEndian, v.Val)
	// }
	// //mode list end
	// writer.WriteByte(0)

	// writer.Flush()

	for _, v := range userConf.ModeList {
		modeList = append(modeList, ssh.Marshal(&v)...)
	}
	modeList = append(modeList, 0)
	ptyReqMes.Term = userConf.PtyRequestMsg.Term
	ptyReqMes.Columns = userConf.PtyRequestMsg.Columns
	ptyReqMes.Rows = userConf.PtyRequestMsg.Rows
	ptyReqMes.Width = userConf.PtyRequestMsg.Width
	ptyReqMes.Height = userConf.PtyRequestMsg.Height
	ptyReqMes.Modelist = string(modeList)

	_, err = channel.SendRequest("pty-req", true, ssh.Marshal(&ptyReqMes))
	if err != nil {
		logrus.Errorf("internal error, error to send message to pty request channel")
		goto end
	}

	_, err = channel.SendRequest("shell", true, nil)
	if err != nil {
		logrus.Errorf("internal error, error to send message to shell channel")
		goto end
	}

	return channel, nil
end:
	return nil, err
}
