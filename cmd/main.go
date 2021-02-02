/* Copyright © 2021 Bartosz Stopa <stopa323@gmail.com> */
package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/stopa323/gonet/cmd/apply"
)

var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gonet",
	Short: "Hello there!",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	cobra.OnInitialize(setLoggingLevel)

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false,
		"enables verbose output")

	rootCmd.AddCommand(apply.ApplyCmd)
}

func setLoggingLevel() {
	if verbose {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	rootCmd.Execute()
}
