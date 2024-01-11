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

var compressionCmd = &cobra.Command{
	Use:   "newbee",
	Short: "视频画质压缩",
	Run: func(cmd *cobra.Command, args []string) {
		Compression()
	},
}

var comVerbose bool

func init() {
	compressionCmd.Flags().BoolVarP(&comVerbose, "verbose", "v", false, "压缩时打印日志信息,默认false")
	compressionCmd.Flags().StringVarP(&inputFile, "inputFile", "i", "", "输入文件路径(必须)")
	compressionCmd.MarkFlagRequired("inputFile")
	compressionCmd.Flags().StringVarP(&outputFile, "outputFile", "o", "", "输出文件路径")
	rootCmd.AddCommand(compressionCmd)
}

// 视频压缩
// ffmpeg -i D:\src.mov -preset veryslow -crf 18 -c:a copy -c:v libx264 D:\dest1.mp4
func Compression() {

	size, err := util.GetVideoSize(inputFile)

	if err != nil {
		fmt.Println("视频不存在")
		return
	}

	if len(outputFile) == 0 {
		iFileName := filepath.Base(inputFile)
		s := path.Ext(iFileName)
		outputFile = strings.Replace(iFileName, s, "-newbee"+s, -1)
	}

	// preset选项： ultrafast, superfast, veryfast, faster, fast, medium, slow, slower, veryslow and placebo.
	// 编码加快，意味着信息丢失越严重，输出图像质量越差
	args := []string{"-i", inputFile, "-preset", "medium", "-crf", "26", "-c:a", "copy", "-c:v", "libx264", "-y", outputFile}
	cmd := exec.Command("ffmpeg", args...)

	if verbose {
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	size2, _ := util.GetVideoSize(outputFile)

	fmt.Printf("视频压缩完成，原视频大小：%s, 压缩后视频大小：%s \n", size, size2)
	fmt.Printf("视频压缩即使肉眼看不出区别，视频信息也会有一定丢失，用于专业领域请慎重")
}
