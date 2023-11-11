package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var recordingCmd = &cobra.Command{
	Use:   "rec",
	Short: "录屏",
	Run: func(cmd *cobra.Command, args []string) {
		Recording()
		fmt.Println("开发中，敬请期待！")
	},
}

func init() {
	// recordingCmd.PersistentFlags().StringVarP(&)
	rootCmd.AddCommand(recordingCmd)
}

// 录屏
func Recording() {
	args := []string{""}
	exec.Command("ffmpeg", args...)
}
