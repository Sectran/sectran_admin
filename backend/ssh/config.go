package ssh

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

const SSH_SERVICE_DEFAULT_PORT int = 15927

// RFC 4254 Section 6.2.
type PtyReqMsg struct {
	Term     string
	Columns  uint32
	Rows     uint32
	Width    uint32
	Height   uint32
	Modelist []byte
}

type Mode struct {
	Key byte
	Val uint32
}

type Env struct {
	LANG        string
	LC_ALL      string
	LC_MESSAGES string
	LC_NUMERIC  string
	LC_COLLATE  string
	LC_MONETARY string
}

type SetEnvRequest struct {
	Name  string
	Value string
}

// ssh connetion params
type SSHConfig struct {
	Host            string    `json:"Host"`            //host
	Port            int32     `json:"port"`            //dst port
	UserName        string    `json:"userName"`        //username for connect
	Password        string    `json:"password"`        //password
	PasswordAuth    bool      `json:"passwordAuth"`    //use password authentificate
	PublicKeyAuth   bool      `json:"PublicKeyAuth"`   //use private key authentificate
	PrivateKey      string    `json:"privateKey"`      //private key
	InteractiveAuth bool      `json:"interactiveAuth"` //interactive authentificate
	NoAuth          bool      `json:"noAuth"`          //no authentificate
	PtyRequestMsg   PtyReqMsg `json:"PtyRequestMsg"`   //pty channle request meesage
	Env             Env       `json:"Env"`             //ssh client envs
}

func (c *SSHConfig) SetModelist(k byte, v uint32) error {
	if k > 129 {
		return fmt.Errorf("invalid mode item key")
	}

	m := Mode{
		Key: k,
		Val: v,
	}
	c.PtyRequestMsg.Modelist = append(c.PtyRequestMsg.Modelist, ssh.Marshal(&m)...)
	return nil
}

func (c *SSHConfig) SetDefaultConfig() {
	c.SetModelist(53, 1)
	c.SetModelist(128, 14400)
	c.SetModelist(129, 14400)

	c.PtyRequestMsg.Term = "xterm256"
	c.PtyRequestMsg.Rows = 180
	c.PtyRequestMsg.Columns = 90
}

func NewSSHConfig() *SSHConfig {
	config := &SSHConfig{}
	return config
}

func CheckSSHConfig(config *SSHConfig) error {
	if config == nil {
		return fmt.Errorf("config can't be nil")
	}

	if config.Port <= 0 {
		return fmt.Errorf("port can't less than zero")
	}

	if config.NoAuth {
		config.NoAuth = true
		config.PasswordAuth = false
		config.PublicKeyAuth = false
		config.InteractiveAuth = false
		return nil
	}

	if config.InteractiveAuth {
		config.PasswordAuth = false
		config.PublicKeyAuth = false
		config.NoAuth = false
		return nil
	}

	if len(config.UserName) <= 0 {
		return fmt.Errorf("you have to provice username if you need to access target ssh server without KbdInteractiveAuth")
	}

	if config.PublicKeyAuth {
		if len(config.PrivateKey) <= 0 {
			return fmt.Errorf("you have to provice PrivateKey if you need to access target ssh server with PublicKeyAuth")
		}
		config.InteractiveAuth = false
		config.PasswordAuth = false
		config.NoAuth = false
	}

	if config.PasswordAuth {
		if len(config.Password) <= 0 {
			return fmt.Errorf("you have to provice password if you need to access target ssh server without KbdInteractiveAuth")
		}
		config.PublicKeyAuth = false
		config.InteractiveAuth = false
		config.NoAuth = false
	}

	return nil
}
