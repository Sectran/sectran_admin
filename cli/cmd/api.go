package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "open api",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		initRoute()
	},
}

func initRoute() {
	r := gin.Default()
	//curl http://localhost:8080/hello 获取到json返回值
	//{“name”:"hello world"}
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
