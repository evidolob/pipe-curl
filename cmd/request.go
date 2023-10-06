package cmd

import (
	"fmt"
	"io"
	"net/http"
)

const (
	GET = "GET"
)

func makeRequest(httpMethod, host, pipePath string) error {
	if httpMethod == "" {
		httpMethod = GET
	}

	client := http.Client{
		Transport: transport(pipePath),
	}

	req, err := http.NewRequest(httpMethod, host, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("Http Error: %s", err)
		return err
	}

	bodyStr, _ := io.ReadAll(resp.Body)
	fmt.Println(string(bodyStr))

	return nil
}
