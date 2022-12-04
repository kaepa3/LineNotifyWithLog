package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/utahta/go-linenotify"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return
}

func main() {
	loadEnv()
	token := os.Getenv("LINE_TOKEN")
	text := createText(os.Getenv("LOG_PATH"))
	title := os.Getenv("TITLE")
	c := linenotify.NewClient()
	msg := fmt.Sprintf("%s \n%s", title, text)
	c.Notify(context.Background(), token, msg, "", "", nil)
}

func createText(path string) string {
	if Exists(path) {
		return readFile(path)
	}
	return "file is nothing"
}

func readFile(path string) string {
	fp, err := os.Open(path)
	if err != nil {
		return err.Error()
	}
	defer fp.Close()

	buf := make([]byte, 500)
	n, err := fp.Read(buf)
	if n == 0 {
		return "empty"
	}
	if err != nil {
		return err.Error()
	}
	return string(buf)
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
