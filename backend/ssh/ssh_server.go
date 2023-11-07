package ssh

import (
	"crypto/rsa"
	"fmt"
	"net"
	"os"
	"sectran/common/utils/cert"
	"sectran/common/utils/reflect"
	"sectran/common/utils/rw"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

const (
	SectranSSHDVeriosn    string = "SSH-2.0-Sectran"
	SectranSSHDPrivateKey string = "id_rsa"
	SectranSSHDPublicKey  string = "id_rsa.pub"
	SectranWelcome        string = "Welcome to SectranV1.0.\r\nAny question plz contact ryanwymail@163.com."
)

func getSSHSigner() (ssh.Signer, error) {
	var (
		err             error
		priKey          *rsa.PrivateKey
		privateKeyBytes []byte
	)
	_, err = os.Stat(SectranSSHDPrivateKey)
	if err == nil {
		privateKeyBytes, err = os.ReadFile(SectranSSHDPrivateKey)
		if err != nil {
			return nil, err
		}
	} else if os.IsNotExist(err) {

		priKey, err = cert.GeneratePrivateKey(1024)
		if err != nil {
			return nil, err
		}
		publicKeyBytes, err := cert.GeneratePublicKey(&priKey.PublicKey)
		if err != nil {
			return nil, err
		}
		privateKeyBytes = cert.EncodePrivateKeyToPEM(priKey)

		err = os.WriteFile(SectranSSHDPublicKey, publicKeyBytes, 0600)
		if err != nil {
			return nil, err
		}

		err = os.WriteFile(SectranSSHDPrivateKey, privateKeyBytes, 0600)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		panic("Failed to parse private key")
	}
	return signer, nil
}

func NewSSHServer(conn net.Conn, userConf *SSHConfig) (*SSHChannels, error) {
	//always auth
	config := &ssh.ServerConfig{
		ServerVersion: SectranSSHDVeriosn,
		NoClientAuth:  false,
	}

	singer, err := getSSHSigner()
	if err != nil {
		return nil, fmt.Errorf("error generate key paire")
	}
	config.AddHostKey(singer)

	permissions := &ssh.Permissions{}

	//manual auth
	if userConf.InteractiveAuth {
		config.KeyboardInteractiveCallback = func(conn ssh.ConnMetadata, client ssh.KeyboardInteractiveChallenge) (*ssh.Permissions, error) {
			questions := []string{"Enter your access token:"}
			answers, err := client("", SectranWelcome, questions, []bool{true})
			if err != nil {
				return nil, err
			}
			userConf.Password = answers[0]

			// questions := []string{"target:", "username:", "password:"}
			// answers, err := client("", SectranWelcome, questions, []bool{true, true, false})
			// if err != nil {
			// 	return nil, err
			// }

			// userConf.Host = answers[0]
			// userConf.UserName = answers[1]
			// userConf.Password = answers[2]
			// userConf.Port = 22
			logrus.Infof("Selected interactive authentication and authentication always release")
			return permissions, nil
		}
	} else if userConf.PasswordAuth {
		config.PasswordCallback = func(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
			logrus.Infof("Selected password authentication and authentication always release")
			return permissions, nil
		}
	} else if userConf.PublicKeyAuth {
		config.PublicKeyCallback = func(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			logrus.Infof("Selected public key authentication and authentication always release,,key type is:%+v", key.Type())
			return permissions, nil
		}
	} else {
		config.NoClientAuth = true
	}

	_, chans, reqs, err := ssh.NewServerConn(conn, config)
	if err != nil {
		logrus.Errorf("new ssh server: %v", err)
		return nil, err
	}
	//discard all requests
	requestHandler := func(reqs <-chan *ssh.Request) {
		for r := range reqs {
			if r.WantReply {
				logrus.Debugf("discard  request tyep of:%s", r.Type)
				r.Reply(false, nil)
			}
		}
	}
	go requestHandler(reqs)

	chans_ := &SSHChannels{
		Pty:  make(chan ssh.Channel, 1),
		Sftp: make(chan ssh.Channel, 1),
		Err:  make(chan error, 1),
	}

	go newChannelHandler(chans_.Pty, chans_.Sftp, chans_.Err, userConf, chans)
	return chans_, nil
}

func newChannelHandler(pty chan ssh.Channel, sftp chan ssh.Channel, errc chan error, userConf *SSHConfig, chans <-chan ssh.NewChannel) {
	for newChannel := range chans {
		channelType := newChannel.ChannelType()
		if channelType != "session" {
			newChannel.Reject(ssh.UnknownChannelType, fmt.Sprintf("Unknown SSH Channel Type: %s, only `session` is supported", channelType))
		}

		channel, requests, err := newChannel.Accept()
		if err != nil {
			newChannel.Reject(ssh.ConnectionFailed, "Failed to accept SSH Channel Request, developers are working on it.")
			errc <- err
		}
		logrus.Infof("incomming a ssh channel")
		go seletPtyChannel(pty, sftp, channel, requests, userConf)
	}
}

// select pty channle
func seletPtyChannel(pty_chan chan ssh.Channel, sftp_chan chan ssh.Channel, channel ssh.Channel, sshReqChan <-chan *ssh.Request, userConf *SSHConfig) {
	for {
		select {
		case req := <-sshReqChan:
			if req == nil {
				continue
			}
			r := rw.NewReader(req.Payload)

			switch req.Type {
			case "pty-req", "shell":
				if req.WantReply {
					req.Reply(true, nil)
				}
				if len(req.Payload) > 0 {
					if r.CheckLength(4) < 0 {
						break
					}
					termLen, _ := r.ReadBigEndian32()
					if r.CheckLength(int(termLen)) < 0 {
						break
					}
					//read terminal type
					termStr, _ := r.ReadBytes(int(termLen))
					userConf.PtyRequestMsg.Term = string(termStr)
					logrus.Debugf("type of terminal is:%s", termStr)

					if r.CheckLength(int(16)) < 0 {
						break
					}

					//read width and height
					userConf.PtyRequestMsg.Columns, _ = r.ReadBigEndian32()
					userConf.PtyRequestMsg.Rows, _ = r.ReadBigEndian32()
					userConf.PtyRequestMsg.Width, _ = r.ReadBigEndian32()
					userConf.PtyRequestMsg.Height, _ = r.ReadBigEndian32()

					//read mode list length
					if r.CheckLength(4) < 0 {
						break
					}
					modeListLen, _ := r.ReadBigEndian32()
					if modeListLen <= 0 {
						logrus.Errorf("error read mode list length")
						break
					}
					if r.CheckLength(int(modeListLen)) < 0 {
						break
					}

					//read all modelist item (contains end zero)
					userConf.PtyRequestMsg.Modelist, _ = r.ReadBytes(int(modeListLen))

					pty_chan <- channel
				}
			case "env":
				if r.CheckLength(int(4)) < 0 {
					break
				}
				envNameLen, _ := r.ReadBigEndian32()

				if r.CheckLength(int(envNameLen)) < 0 {
					break
				}
				envName, _ := r.ReadBytes(int(envNameLen))

				if r.CheckLength(int(4)) < 0 {
					break
				}
				envLen, _ := r.ReadBigEndian32()

				if r.CheckLength(int(envLen)) < 0 {
					break
				}
				env, _ := r.ReadBytes(int(envLen))

				//export LANG=zh_CN.UTF-8/string(env)
				if reflect.SetVal(&userConf.Env, string(envName), string(env)) {
					logrus.Debugf("env of %s's value is: %s", envName, env)
				} else {
					logrus.Warn("can not set config envs with SetVal")
				}
			case "subsystem":
				logrus.Infof("request subsystem:%s", req.Payload[4:])
				if string(req.Payload[4:]) == "sftp" {
					req.Reply(true, nil)
					sftp_chan <- channel
				}
			default:
				logrus.Errorf("recieve unhandled request tyep of:%s,%s", req.Type, req.Payload)
				if req.WantReply {
					req.Reply(false, nil)
				}
			}
		case <-time.After(time.Duration(10) * time.Second):
			logrus.Debugf("auto exit this channel request")
			return
		}
	}
}
