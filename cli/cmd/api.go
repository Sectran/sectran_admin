package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sectran/api/db"
	"sectran/api/router"
)

func init() {

}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "open api",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		initialize()
	},
}

func Execute() {
	if err := apiCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func initialize() {
	//连接mysql
	db.MysqlConnect()
	router.InitRouter()
}
