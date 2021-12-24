package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func removeResults() error {
	err := exec.Command("rm", "-rf", "results").Run()
	fmt.Println("[X] REMOVE PREVIOUS")
	if err != nil {
		return err
	}
	return err
}

func genID() string {
	time := time.Now()
	id := time.Unix()
	return strconv.FormatInt(id, 10)
}

func runGFPGAN() error {
	var err error
	id := genID()
	dir := fmt.Sprintf("results/%s", id)
	cmd := exec.Command("./GFPGAN/venv/bin/python3", "inference_gfpgan.py", "--upscale", "4", "--test_path", "inputs/1", "--save_root", dir, "--model_path", "experiments/pretrained_models/GFPGANCleanv1-NoCE-C2.pth", "--bg_upsampler", "realesrgan")

	_, err = cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	_, err = cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	cmd.Wait()
	fmt.Println("[X] Done GFPGAN")
	return err
}

func execute() {
	var err error
	err = removeResults()
	if err != nil {
		panic(err)
	}

	err = runGFPGAN()
	if err != nil {
		panic(err)
	}

}

func copyOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

/* 
? TODO
[ ] Add gRPC Server
		[ ] Send Data To app Service
		[ ] Recive Data from app
[ ] Handle Image Processing
[ ] Handle Remove/Delete Image 
[ ] Handle Event Driven
*/

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute()
	}
}
