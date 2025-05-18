/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

type ProductPage struct {
	User string
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "server",
	Long: `start bookinfo server:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		r.GET("/", func(c *gin.Context) {
			data, _ := os.ReadFile("templates/index.html")
			tmpl, err := template.New("productpage").Parse(string(data))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}
			product := ProductPage{}
			buf := bytes.Buffer{}
			err = tmpl.Execute(&buf, product)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}
			c.Data(http.StatusOK, "text/html, charset=utf-8", buf.Bytes())
		})

		r.GET("/productpage", func(c *gin.Context) {
			data, _ := os.ReadFile("templates/productpage.html")
			tmpl, err := template.New("productpage").Parse(string(data))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}
			product := ProductPage{
				User: "jason",
				// User: "jason",
			}
			buf := bytes.Buffer{}
			err = tmpl.Execute(&buf, product)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}
			c.Data(http.StatusOK, "text/html, charset=utf-8", buf.Bytes())
		})
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
