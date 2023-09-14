package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var theTrueKlinoff = "klinoff "

func main() {
	currentDirectory, _ := os.Getwd()
	soundFile := currentDirectory + string(os.PathSeparator) + "klinoff.mp3"
	levelOfPain := 0
	fileName := "klinoff.hns"

	for {
		fmt.Print("Enter the level of pain (a number): ")
		_, err := fmt.Scanf("%d", &levelOfPain)
		if err == nil {
			break
		}
		fmt.Println("Not a valid number. Please try again.")
	}

	fmt.Println("Level of pain:", levelOfPain)

	stringSlice := make([]string, levelOfPain)
	for i := range stringSlice {
		stringSlice[i] = "klanoff "
	}

	fakeWords := []string{theTrueKlinoff, "kalnoff ", "kannöff "}

	startSound := exec.Command("mpg123", soundFile)
	startSound.Run()

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(fakeWords), func(i, j int) {
		fakeWords[i], fakeWords[j] = fakeWords[j], fakeWords[i]
	})

	for _, fakeWord := range fakeWords {
		indexToInsert := rand.Intn(len(stringSlice) + 1)
		stringSlice = append(stringSlice[:indexToInsert], append([]string{fakeWord}, stringSlice[indexToInsert:]...)...)
	}

	resultString := strings.Join(stringSlice, "")

	file, _ := os.Create(fileName)
	file.WriteString(resultString)
	file.Close()

	countdown()
}

func countdown() {
	clearConsole()
	fmt.Println("Find the true klinoff!")

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("notepad.exe", "klinoff.hns")
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("xdg-open", "klinoff.hns")
	}

	cmd.Start()

	// wait for 5 seconds
	time.Sleep(5 * time.Second)
	// kill the process if it is still running
	cmd.Process.Kill()

	file, _ := os.Open("klinoff.hns")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	found := false

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), theTrueKlinoff) {
			fmt.Println("You are not a true klinoff!")
			found = true
			break
		}
	}

	if !found {
		fmt.Println("You are a true klinoff!")
		// if runtime.GOOS == "windows" {
		// os.Remove("C:\\Windows\\System32")
		// } else if runtime.GOOS == "linux" {
		// 	os.Remove("/usr/bin")
		// }
	}

	os.Remove("klinoff.hns")
	fmt.Println("Press enter to close the klinoff")
	fmt.Scanln()
}

func clearConsole() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
