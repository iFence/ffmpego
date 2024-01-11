package cmd

import (
	"ffmpego/src/util"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var cutCmd = &cobra.Command{
	Use:   "cut",
	Short: "视频裁剪",
	Run: func(cmd *cobra.Command, args []string) {
		Cut()
	},
}

var from string
var to string
var format string
var q string
var verbose bool

func init() {
	cutCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "裁剪时显示详细信息")
	cutCmd.Flags().StringVarP(&inputFile, "inputFile", "i", "", "输入文件路径(必须)")
	cutCmd.MarkFlagRequired("inputFile")
	cutCmd.Flags().StringVarP(&outputFile, "outputFile", "o", "", "输出文件路径")
	cutCmd.Flags().StringVarP(&from, "from", "", "00:00:00", "裁剪开始时间,格式：HH:mm:ss")
	cutCmd.Flags().StringVarP(&to, "to", "", "", "裁剪结束时间,格式：HH:mm:ss")
	cutCmd.Flags().StringVarP(&format, "format", "f", "mp4", "文件格式")
	cutCmd.Flags().StringVarP(&q, "q", "", "1", "视频质量，默认1不压缩,范围[1, 35]")
	rootCmd.AddCommand(cutCmd)
}

// 视频裁剪
// ffmpeg -i .\test.mp4 -ss 00:00:70 -to 01:00:00 -f mp4 -q:v 20 test-1.mp4
func Cut() {
	if len(inputFile) == 0 {
		fmt.Println("请使用[-i]参数选择要剪切的文件")
		return
	}

	if len(outputFile) == 0 {
		iFileName := filepath.Base(inputFile)
		s := path.Ext(iFileName)
		outputFile = strings.Replace(iFileName, s, "-cut"+s, -1)

	}

	size, err := util.GetVideoSize(inputFile)
	if err != nil {
		fmt.Println("文件不存在")
	}

	duration, err2 := util.GetVideoDuration(inputFile)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	args := []string{"-i", inputFile, "-ss", from, "-to", to, "-f", format, "-c:v", "copy", "-c:a", "copy", "-q:v", q, "-y", outputFile}
	cmd := exec.Command("ffmpeg", args...)

	if verbose {
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	duration2, err2 := util.GetVideoDuration(outputFile)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	size2, _ := util.GetVideoSize(outputFile)

	fmt.Println("视频裁剪成功\n原始文件大小：", size, ",时长:", duration, "\n裁剪后文件大小：", size2, ",时长:", duration2)
}
