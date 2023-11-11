package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "视频播放",
	Run: func(cmd *cobra.Command, args []string) {
		Recording()
		fmt.Println("开发中，敬请期待！")
	},
}

func init() {
	// recordingCmd.PersistentFlags().StringVarP(&)
	rootCmd.AddCommand(playCmd)
}

// 视频播放
func Play() {
	args := []string{""}
	exec.Command("ffplay", args...)
}
