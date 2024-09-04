package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"os"
)

type Input struct {
	Payload string `json:"payload"`
}

type Output struct {
	Body string `json:"body"`
	Code int    `json:"code"`
}

func main() {
	Start(func(payload Input) (Output, error) {
		return Output{
			Body: payload.Payload,
			Code: 200,
		}, nil
	})
}

func Start(handler any) {
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		lambda.Start(handler)
	} else {
		CliStart(handler)
	}
}

func CliStart(handler any) {
	h := lambda.NewHandlerWithOptions(handler)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		invoke, err := h.Invoke(context.Background(), scanner.Bytes())
		if err != nil {
			log.Println("error:", err)
		}

		if _, err := fmt.Fprintln(os.Stdout, string(invoke)); err != nil {
			log.Println("error:", err)
		}
	}
}
