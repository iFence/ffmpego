package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ffmpego",
	Short: "ffmpego是一个命令行录屏和视频编辑工具",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("输入 [ffmpego -h] 根据提示进行操作")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
