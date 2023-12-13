package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	GET = "GET"
)

func makeRequest(httpMethod, host, pipePath string, includeProtocol bool) error {
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

	if includeProtocol {
		fmt.Printf("%s %s\n", resp.Proto, resp.Status)

		for k, v := range resp.Header {
			fmt.Printf("%s: %s\n", k, strings.Join(v, " "))
		}
	}

	defer resp.Body.Close()

	for {

		buff := make([]byte, 16)
		_, err := resp.Body.Read(buff)

		fmt.Printf("%s", buff)

		if err != nil {
			if err == io.EOF {
				fmt.Fprintln(os.Stderr, "AAAAAAAA")
				break
			}
			fmt.Errorf("error on reading response: %s", err)
			return err
		}

	}

	return nil
}
