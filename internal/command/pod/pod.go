package pod

import (
	"log/slog"

	"github.com/spf13/cobra"
)

var CmdPod = &cobra.Command{
	Use:   "pod",
	Short: "pod",
	Long:  `Pod helper`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var cmdPodClean = &cobra.Command{
	Use:   "clean",
	Short: "clean pod",
	Long:  `clean pod`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("clean pod")
	},
}

func init() {
	CmdPod.AddCommand(cmdPodClean)
}
