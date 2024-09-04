package main

import (
	"bufio"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"os"
)

func main() {
	Start(func(payload string) (string, error) {
		fmt.Println("got payload", payload)
		return payload, nil
	})
}

func Start(handler func(string) (string, error)) {
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		lambda.Start(handler)
	} else {
		CliStart(handler)
	}
}

func CliStart(handler func(string) (string, error)) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		payload := scanner.Text()
		result, err := handler(payload)
		if err != nil {
			log.Println("error:", err)
		}

		fmt.Fprintln(os.Stdout, result)
	}
}
