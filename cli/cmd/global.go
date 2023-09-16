package cmd

import (
	"github.com/spf13/cobra"
)

var GConfig GlobalFlags

// GlobalFlags are flags that defined globally
// and are inherited to all sub-commands.
type GlobalFlags struct {
	protocol   string // websocket service
	port       int64  // service port
	logLevel   int64  // log level
	loggerFile string // save file
	target     string
}

func getGlobalConf(command *cobra.Command) (conf GlobalFlags, err error) {
	conf.protocol, err = command.Flags().GetString("protocol")
	if err != nil {
		return
	}

	conf.port, err = command.Flags().GetInt64("port")
	if err != nil {
		return
	}

	conf.logLevel, err = command.Flags().GetInt64("log-level")
	if err != nil {
		return
	}

	conf.loggerFile, err = command.Flags().GetString("log-file")
	if err != nil {
		return
	}

	conf.target, err = command.Flags().GetString("target")
	if err != nil {
		return
	}
	return
}
