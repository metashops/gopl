package main

import (
	"context"
	"fmt"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {
	key := "key"
	c := gogpt.NewClient(key)
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3Ada,
		MaxTokens: 5,
		Prompt:    "",
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	fmt.Println(resp.Choices[0].Text)
}
