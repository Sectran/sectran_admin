package ssh

import (
	"io"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

const (
	SectranSSHCVeriosn string = "SSH-2.0-Sectran"
)

func NewSSHClient(userConf *SSHConfig) (io.ReadWriteCloser, error) {
	if err := CheckSSHConfig(userConf); err != nil {
		return nil, err
	}

	var (
		auth      []ssh.AuthMethod
		err       error
		signer    ssh.Signer
		pri       []byte
		config    *ssh.ClientConfig
		client    *ssh.Client
		request   <-chan *ssh.Request
		channel   ssh.Channel
		ptyReqMes PtyReqMsg
	)

	//InteractiveAuth and PasswordAuth is the same for client side
	if userConf.PasswordAuth || userConf.InteractiveAuth || userConf.NoAuth {
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
		BannerCallback: func(message string) error {
			return nil
		},
		// maybe network is so bad？
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
