package cmd

import (
	"os"
	"sectran/cli/cobrautl"
	"sectran/cli/version"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	cliName        = "sectran"
	cliDescription = "proxy your ssh and telnet with audit."
)

var (
	GitVersion = "v0.0.0_unknow"
)

const (
	defaultListenPort int64  = 19527
	defaultLogLevel   int64  = int64(logrus.DebugLevel)
	defaultProtocol   string = "tcp"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:        cliName,
	Short:      cliDescription,
	SuggestFor: []string{"sectran_char"},

	Long: `"sectran_char" is a module of "sectran" that provides web and TCP proxies for 
SSH/Telnet, along with features such as session auditing, session blocking, session sharing, and more.`,
}

func init() {
	cobra.EnablePrefixMatching = true
	var globalFlags = GlobalFlags{}

	rootCmd.PersistentFlags().StringVarP(&globalFlags.protocol, "protocol", "p", defaultProtocol, "which protocol you want to connect to Sectran[tcp or websocket]")
	rootCmd.PersistentFlags().Int64VarP(&globalFlags.port, "port", "P", defaultListenPort, "listening port")
	rootCmd.PersistentFlags().Int64VarP(&globalFlags.logLevel, "log-level", "", defaultLogLevel, "sectran log level,default is DEBUG,0 PanicLevel 1 FatalLevel 2 ErrorLevel 3 WarnLevel 4 InfoLevel 5 DebugLevel 6 TraceLevel")
	rootCmd.PersistentFlags().StringVarP(&globalFlags.loggerFile, "log-file", "l", "", "save the logs to file")
	rootCmd.PersistentFlags().StringVarP(&globalFlags.target, "target", "t", "", "the target server you want to proxy to[ip:port]")

	// 0 PanicLevel 1 FatalLevel 2 ErrorLevel 3 WarnLevel
	// 4 InfoLevel 5 DebugLevel 6 TraceLevel
	logrus.SetLevel(logrus.Level(globalFlags.logLevel))
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})

	if len(globalFlags.loggerFile) > 0 {
		f, _ := os.OpenFile(globalFlags.loggerFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
		logrus.SetOutput(f)
	}

	logrus.Infof("Commit: %s Branch: %s BuildTime:%s", version.Commit, version.Branch, version.BuildTime)
}

func usageFunc(command *cobra.Command) error {
	return cobrautl.UsageFunc(command, GitVersion)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.SetUsageFunc(usageFunc)
	rootCmd.SetHelpTemplate(`{{.UsageString}}`)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Version = GitVersion
	rootCmd.SetVersionTemplate(`{{with .Name}}{{printf "%s " .}}{{end}}{{printf "version:\t%s" .Version}}
	`)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
