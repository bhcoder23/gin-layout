/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	appserver "github.com/bhcoder23/gin-layout/internal/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startHTTPServer = appserver.StartHTTPServer

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "Start the HTTP server",
	Example: "gin-layout server -c configs/config.yaml",
	RunE: func(cmd *cobra.Command, args []string) error {
		host := viper.GetString("server.host")
		port := viper.GetInt("server.port")
		ctx := cmd.Context()
		if ctx == nil {
			ctx = context.Background()
		}
		return startHTTPServer(ctx, host, port)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
