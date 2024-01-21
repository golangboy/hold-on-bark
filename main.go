package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	BarkKey := os.Getenv("BARK_KEY")
	Interval := os.Getenv("INTERVAL")
	if BarkKey == "" {
		fmt.Println("BARK_KEY is empty")
		return
	}
	if Interval == "" {
		Interval = "60"
	}
	intevalInt, err := strconv.Atoi(Interval)
	if err != nil {
		fmt.Println("INTERVAL is not a number")
		return
	}
	file, err := os.ReadFile("./holdon.txt")
	if err != nil {
		return
	}
	txt := string(file)
	sentences := strings.Split(txt, "\n")
	for {
		for sentence := range sentences {
			api := fmt.Sprintf("https://api.day.app/%s/title/%s", BarkKey, sentences[sentence])
			http.Get(api)
			time.Sleep(time.Duration(intevalInt) * time.Minute)
		}
	}
}
