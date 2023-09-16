package cmd

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"sectran/conn"
	"sectran/proto/ssh"
	"sectran/user/config"
	"sectran/utils"
	"strconv"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// server config
var sc *config.SSHConfig = config.NewSSHConfig()

// opensslCmd represents the openssl command
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "use sectran to start a sshd service or a sshd service with proxy",
	Long: `use secTran to start a web or tcp ssh proxy,Let all traffic pass through our agent, and we will make a detailed audit information of user operations.
sectran ssh
sectran ssh -p 10022 -t
`,
	Run: SSHCommandFunc,
}

func init() {
	sshCmd.PersistentFlags().StringVar(&sc.UserName, "username", "", "username of sectran sshd service")
	sshCmd.PersistentFlags().StringVar(&sc.Password, "password", "", "password of sectran sshd  service")
	sshCmd.PersistentFlags().BoolVar(&sc.PasswordAuth, "password-auth", true, "sectran sshd service authentication method")
	sshCmd.PersistentFlags().BoolVar(&sc.PublicKeyAuth, "publickey-auth", false, "sectran sshd service authentication method")
	sshCmd.PersistentFlags().BoolVar(&sc.InteractiveAuth, "interactive-auth", false, "user manually login in using the keyboard")
	sshCmd.PersistentFlags().StringVar(&sc.PrivateKey, "privatekey", "", " path in proxy mode,if you want to use public key auth in proxy,you must offer private key")
	rootCmd.AddCommand(sshCmd)
}

func SSHCommandFunc(command *cobra.Command, args []string) {
	var (
		err     error
		stopper chan os.Signal
	)

	stopper = make(chan os.Signal, 1)
	signal.Notify(stopper, os.Interrupt, syscall.SIGTERM)

	// save global config
	GConfig, _ = getGlobalConf(command)
	sc.Port = int32(GConfig.port)

	out := fmt.Sprintf("Serctran SSHD Service start in port: %d", sc.Port)
	logrus.Debug(out)

	err = sshdStartup(GConfig)
	if err != nil {
		logrus.Errorf("failed to start sectran sshd :%s", err)
	}
}

func sshdStartup(conf GlobalFlags) error {
	var (
		l   net.Listener
		err error
	)

	if !config.CheckSSHConfig(sc) {
		return fmt.Errorf("invalid user SSH configuration")
	}
	sc.Host = fmt.Sprintf(":%d", conf.port)

	l, err = net.Listen("tcp", sc.Host)
	if err != nil {
		logrus.Fatalf("error listen port %d.due to %s", conf.port, err)
	}
	defer l.Close()

	signalChan := utils.ReigisterSignal()
	for {
		select {
		case conn := <-conn.AcceptConnection(l, conf.protocol == "websocket"):
			go handlePeerConnection(conn)
		case s := <-signalChan:
			logrus.Infof("recieve signal:%d,sectran will stop in one while", s)
			//todo:other work
			os.Exit(0)
		}
	}

}

func handlePeerConnection(conn net.Conn) {
	defer conn.Close()

	//copy config from server config
	userConf := *sc

	rwcs, err := ssh.NewSSHServer(conn, &userConf)
	if err != nil {
		logrus.Errorf("error to build SSH server")
		return
	}

	// Split the address into its components (host and port)
	hostStr, portStr, err := net.SplitHostPort(GConfig.target)
	if err != nil {
		logrus.Errorf("error to parse target server address")
		return
	}

	// Convert the string to uint32
	port, err := strconv.ParseUint(portStr, 10, 32)
	if err != nil {
		logrus.Errorf("error to parse target server address:%s", err)
		return
	}
	userConf.Host = hostStr
	userConf.Port = int32(port)

	rwcc, err := ssh.NewSSHClient(&userConf)
	if err != nil {
		logrus.Errorf("error to get SSH client instance:%s", err)
		return
	}
	ctx, cf := context.WithCancel(context.Background())

	reversedFunc := func(r io.Reader, w io.Writer, cf context.CancelFunc) {
		defer cf()

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
				}
			}
		}
	}

	go reversedFunc(rwcs, rwcc, cf)
	go reversedFunc(rwcc, rwcs, cf)

	<-ctx.Done()
	logrus.Info("a connection is stoped")
}
