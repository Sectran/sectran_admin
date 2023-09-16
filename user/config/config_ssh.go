package config

// RFC 4254 Section 6.2.
type PtyReqMsg struct {
	Term     string
	Columns  uint32
	Rows     uint32
	Width    uint32
	Height   uint32
	Modelist string
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
	PtyRequestMsg   PtyReqMsg `json:"PtyRequestMsg"`   //pty channle request meesage
	ModeList        []Mode    `json:"ModeList"`        //ssh mode list info
	Env             Env       `json:"Env"`             //ssh client envs
}

func NewSSHConfig() *SSHConfig {
	config := &SSHConfig{}
	return config
}

func CheckSSHConfig(config *SSHConfig) bool {
	var result bool = false

	if config.Port <= 0 {
		goto end
	}

	if config.InteractiveAuth {
		result = true
		config.PasswordAuth = false
		config.PublicKeyAuth = false
		goto end
	}

	if len(config.UserName) <= 0 {
		goto end
	}

	if config.PublicKeyAuth {
		result = true
		config.InteractiveAuth = false
		config.PasswordAuth = false
		goto end
	}

	if config.PasswordAuth {
		if len(config.Password) <= 0 {
			goto end
		}
		config.PublicKeyAuth = false
		config.InteractiveAuth = false
	}
	result = true
end:
	return result
}
