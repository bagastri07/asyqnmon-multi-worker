package cmd

import (
	"os"

	"github.com/bagastri07/asyqnmon-multi-worker/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "anysqmon-multi-worker",
	Short: "A simple queue example",
}

func Execute() {
	config.GetConf()

	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
