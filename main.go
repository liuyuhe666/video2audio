package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "参数错误")
		os.Exit(1)
	}
	input := os.Args[1]
	fmt.Println("input: ", input)
	output := input + ".aac"
	fmt.Println("output: ", output)
	cmd := exec.Command("ffmpeg", "-i", input, "-acodec", "copy", "-vn", output)
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "命令执行失败: %s\n", err)
	}
}
