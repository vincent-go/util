package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/joho/godotenv"
)

type Config struct {
	TargetFile string `json:"target_file"`
}

func readfile() {
	// Read config file
	configData, err := os.ReadFile("config.json")
	if err != nil {
		log.Println("Error reading config file:", err)
		return
	}

	// Unmarshal config data
	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		log.Println("Error unmarshaling config:", err)
		return
	}

	// Read target file
	data, err := os.ReadFile(config.TargetFile)
	if err != nil {
		log.Println("Error reading target file:", err)
		return
	}

	log.Println(string(data))
}

func firstToLower(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && size <= 1 {
		return s
	}
	lc := unicode.ToLower(r)
	if r == lc {
		return s
	}
	return string(lc) + s[size:]
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fileName := os.Getenv("First file")
	// secretKey := os.Getenv("SECRET_KEY")

	fileName = strings.TrimRight(fileName, "\r\n")

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(dir)

	fp := filepath.Join(dir, fileName) // Construct full path

	pattern := `hello` // Replace with your regex pattern

	fp = strings.Replace(fp, "\\", "/", -1)

	file, err := os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			log.Println("Match found:", line)
			// Do something with the matched line
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println("Job is done, press 'Enter' key to close the window")
	fmt.Scanf("h")
}
