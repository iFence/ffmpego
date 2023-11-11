package util

import (
	"context"
	"fmt"
	"os"
	"time"

	"gopkg.in/vansante/go-ffprobe.v2"
)

// 获取视频文件大小
func GetVideoSize(filename string) (string, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		fmt.Println("err =", err)
		return "0", err
	}

	size := fi.Size()
	if size < 1024 {
		return fmt.Sprintf("%d %s", size, "B"), nil
	}

	if size < 1024*1024 {
		return fmt.Sprintf("%d %s", size/1024, "KB"), nil
	}

	if size < 1024*1024*1024 {
		return fmt.Sprintf("%d %s", size/1024/1024, "MB"), nil
	}

	return fmt.Sprintf("%d %s", size/1024/1024/1024, "GB"), nil
}

// 获取视频文件时长
func GetVideoDuration(filename string) (string, error) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()
	pd, err := ffprobe.ProbeURL(ctx, filename)

	if err != nil {

		return "00:00:00", err
	}
	duration := int32(pd.Format.DurationSeconds)

	if duration < 60 {
		return fmt.Sprintf("00:00:%02d", duration), nil
	}

	if duration < 60*60 {
		return fmt.Sprintf("00:%02d:%02d", duration/60, duration%60), nil
	}

	hour := duration / (60 * 60)
	var minutes int32
	if duration%60 < 60 {
		minutes = 0
	} else {
		minutes = (duration % 60) / 60
	}
	return fmt.Sprintf("%02d:%02d:%02d", hour, minutes, (duration%60)%60), err
}
