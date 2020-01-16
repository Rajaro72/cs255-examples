package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

const SuccessEmoji = "✔️"
const NeutralEmoji = "🤷"
const FailureEmoji = "❌"

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	sort.Slice(files, func(i, j int) bool {
		// Sort problem files by problem number
		if strings.HasPrefix(files[i].Name(), "Problem") && strings.HasPrefix(files[j].Name(), "Problem") {
			return problemNumber(files[i]) < problemNumber(files[j])
		}

		// Place non problem files at the beginning
		if strings.HasPrefix(files[i].Name(), "Problem") {
			return false
		}
		if strings.HasPrefix(files[j].Name(), "Problem") {
			return true
		}

		// Sort the non problem files by name
		return files[i].Name() < files[j].Name()
	})

	allOkay := true
	for _, f := range files {
		if !fileIsOk(f) {
			allOkay = false
		}
	}

	if !allOkay {
		os.Exit(1)
	}
}

func problemNumber(file os.FileInfo) int {
	n, err := strconv.Atoi(strings.TrimRight(strings.TrimLeft(file.Name(), "Problem"), ".txt"))
	if err != nil {
		fmt.Println(FailureEmoji, "Failed to determine problem number for file", file.Name())
		os.Exit(4)
	}
	return n
}

// Checks any file to see if it's a problem file (true if not a problem file), and if it is whether it's okay.
func fileIsOk(f os.FileInfo) bool {
	if f.IsDir() {
		fmt.Println(NeutralEmoji, f.Name(), "is a directory, skipping")
		return true
	}

	if !strings.HasSuffix(f.Name(), ".txt") || !strings.HasPrefix(f.Name(), "Problem") {
		fmt.Println(NeutralEmoji, f.Name(), "is not a problem file, skipping")
		return true
	}

	issueMessages := checkProblemFile(f)
	if len(issueMessages) > 0 {
		for _, issueMessage := range issueMessages {
			fmt.Println(issueMessage)
		}
		return false
	} else {
		fmt.Println(SuccessEmoji, f.Name(), "looks good")
		return true
	}
}

// Checks a problem file for issues. Returns a slice of error messages.
func checkProblemFile(f os.FileInfo) []string {
	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	var sections []string
	if runtime.GOOS == "windows" {
		sections = strings.Split(strings.TrimSpace(string(data)), "\r\n===\r\n")
	} else {
		sections = strings.Split(strings.TrimSpace(string(data)), "\n===\n")
	}
	if len(sections) < 2 {
		return []string{FailureEmoji + " " + f.Name() + " did not have at least two sections"}
	}

	// Task 2/3: 75 credits in a week, 4 credits per tutor => at least 18.75 tutors
	tutors := strings.Split(sections[0], "\n")
	if len(tutors) < 19 {
		return []string{FailureEmoji + " " + f.Name() + " has " + strconv.Itoa(len(tutors)) + " tutors, should have at least 19"}
	}

	modules := strings.Split(sections[1], "\n")
	if len(modules) != 25 {
		return []string{FailureEmoji + " " + f.Name() + " has " + strconv.Itoa(len(modules)) + " modules, should have 25"}
	}

	return []string{}
}
