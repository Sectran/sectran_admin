package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"os"
	"sectran/api/db"
)

func init() {

}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "open api",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		initRoute()
	},
}

func Execute() {
	if err := apiCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func initRoute() {
	//连接mysql
	db.MysqlConnect()
	r := gin.Default()
	//curl http://localhost:8080/hello 获取到json返回值

	r.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"name": "hello world",
		})
	})
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
