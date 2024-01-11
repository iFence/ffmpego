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

var version = "2024.1-0.1"

//输入文件
var inputFile string
//输出文件
var outputFile string

func init(){
	rootCmd.Version = "2024.1-0.1"
	rootCmd.SetVersionTemplate("v2024.1-0.1 created by yulei<vyulei97@gmail.com>")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
