package main

import (
	"log"

	"github.com/orvice/kubeh/internal/command/pod"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "kubeh",
	Short:   "kube helper",
	Long:    `a cli helper for kubernetes`,
	Version: "0.0.1",
}

func init() {
	rootCmd.AddCommand(pod.CmdPod)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
