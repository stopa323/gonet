/* Copyright © 2021 Bartosz Stopa <stopa323@gmail.com> */
package apply

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	lang "github.com/stopa323/gonet/pkg/language"
)

var configPath string

// ApplyCmd represents the apply command
var ApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply network configuration",
	Long:  `Will surely add extensive description here.`,
	Run:   apply,
}

func init() {
	ApplyCmd.Flags().StringVarP(&configPath, "filename", "f", "",
		"network configuration file")
	ApplyCmd.MarkFlagRequired("filename")
}

func apply(cmd *cobra.Command, args []string) {
	cfg, diags := lang.ConfigFromFile(configPath)
	if diags.HasErrors() {
		log.Error(diags.Error())
		os.Exit(1)
	}

	log.Info(cfg)
}
