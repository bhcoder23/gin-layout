/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
// Package cmd defines the command-line interface.
package cmd

import (
	"os"

	"github.com/bhcoder23/gin-layout/internal/config"
	"github.com/spf13/cobra"
)

var cfgFile string
var env string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gin-layout",
	Short: "A brief description of your application",
	Long:  ``,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return config.Load(cfgFile, env)
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	rootCmd.PersistentFlags().StringVar(&env, "env", "", "runtime environment suffix, e.g. dev loads configs/config.dev.yaml")
}
