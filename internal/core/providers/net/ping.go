package net

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func Ping(url string) bool {
	timeout := time.Second * 10
	client := http.Client{
		Timeout: timeout,
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}
