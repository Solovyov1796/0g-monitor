package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cmd = cobra.Command{
	Use:   "0g-monitor",
	Short: "Daemon to monitor all 0G service status",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

// Execute is the command line entrypoint.
func Execute() {
	if err := cmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("Failed to execute command")
	}
}
