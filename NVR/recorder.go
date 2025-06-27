package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"time"
)
 
var recordQueue = make(chan *Camera, 10)

func StartRecorderManager() {
	go func() {
		for cam := range recordQueue {
			go recordCamera(cam)
		}
	}()

	var cams []Camera
	DB.Find(&cams)
	for _, cam := range cams {
		recordQueue <- &cam
	}
}

func recordCamera(cam *Camera) {
	for {
		now := time.Now()
		dir := filepath.Join("storage", cam.Name, now.Format("2006/01/02"))
		os.MkdirAll(dir, 0755)
		file := filepath.Join(dir, now.Format("15")+".mp4")

		Log.Infof("Recording: %s -> %s", cam.Name, file)

		cmd := exec.Command("ffmpeg", "-rtsp_transport", "tcp",
			"-i", cam.RTSPURL,
			"-t", "3600",
			"-c", "copy", file)
		err := cmd.Run()
		if err != nil {
			Log.Errorf("FFmpeg error for %s: %v", cam.Name, err)
		}

		cleanupOldFiles(cam)
	}
}

func cleanupOldFiles(cam *Camera) {
	// Пройтись по папкам и удалить файлы старше ArchiveDays
}
