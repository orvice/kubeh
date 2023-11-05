package pod

import (
	"log/slog"

	"github.com/orvice/kubeh/internal/client"
	"github.com/orvice/kubeh/internal/config"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CmdPod = &cobra.Command{
	Use:   "pod",
	Short: "pod",
	Long:  `Pod helper`,
	Run:   podClean,
}

var cmdPodClean = &cobra.Command{
	Use:   "clean",
	Short: "clean pod",
	Long:  `clean pod`,
	Run:   podClean,
}

func init() {
	CmdPod.AddCommand(cmdPodClean)
}

func podClean(cmd *cobra.Command, args []string) {
	cliSet, err := client.NewClient(config.ConfigPath)
	if err != nil {
		slog.Error("init kube config failed", "error", err)
		return
	}
	slog.Info("Pod Clenas", "configPath", config.ConfigPath)
	pods, err := cliSet.CoreV1().Pods("").List(cmd.Context(), metav1.ListOptions{})
	if err != nil {
		slog.Error("list pods failed", "error", err)
		return
	}
	slog.Info("Pods", "total_count", len(pods.Items))
	for _, pod := range pods.Items {
		if pod.Status.Phase == "Running" {
			continue
		}
		slog.Info("Pod start delete", "name", pod.Name, "namespace", pod.Namespace, "status", pod.Status.Phase)
		err := cliSet.CoreV1().Pods(pod.Namespace).Delete(cmd.Context(), pod.Name, metav1.DeleteOptions{})
		if err != nil {
			slog.Error("delete pod failed", "name", pod.Name, "namespace", pod.Namespace, "error", err)
			continue
		}
		slog.Info("Pod delete success", "name", pod.Name, "namespace", pod.Namespace)
	}
}
