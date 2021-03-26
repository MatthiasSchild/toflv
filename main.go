package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

const (
	targetEncoding = "libx264"
	targetCRF      = "28"
)

func convert(source, target string) bool {
	args := []string{
		"-i", source,
		"-c:v", targetEncoding,
		"-crf", targetCRF,
		target,
	}

	fmt.Println("ffmpeg", strings.Join(args, " "))
	cmd := exec.Command("ffmpeg", args...)
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return false
	}

	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				fmt.Printf("Exit Status: %d\n", status.ExitStatus())
				return false
			}
		} else {
			fmt.Println(err)
			return false
		}
	}
	return true
}

func main() {
	filenames := os.Args[1:]

	for _, filename := range filenames {
		ext := filepath.Ext(filename)
		target := strings.TrimSuffix(filename, ext) + ".flv"

		fmt.Println("Source:", filename)
		fmt.Println("Target:", target)

		success := convert(filename, target)
		if success {
			fmt.Println("# Success")
		} else {
			fmt.Println("! Failed")
		}
		fmt.Println()
	}
}
