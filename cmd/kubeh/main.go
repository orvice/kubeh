package main

import (
	"log"

	"github.com/orvice/kubeh/internal/command/pod"
	"github.com/orvice/kubeh/internal/config"
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
	rootCmd.PersistentFlags().StringVar(&config.ConfigPath, "configPath", "kubeconfig path", "kube config path")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
