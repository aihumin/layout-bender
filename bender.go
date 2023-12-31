package main;

import (
    "log"
    "context"
    "fmt"
    "os"
    openai "github.com/sashabaranov/go-openai"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load();
    if err != nil {
        log.Fatal("Error loading .env file");
    }

    client := openai.NewClient(os.Getenv("OPENAI_API_KEY"));
    resp, err := client.CreateChatCompletion(
        context.Background(),
        openai.ChatCompletionRequest{
            Model: openai.GPT3Dot5Turbo,
            Messages: []openai.ChatCompletionMessage{
                {
                    Role:    openai.ChatMessageRoleUser,
                    Content: "Hello!",
                },
            },
        },
    );

    if err != nil {
        log.Fatal("ChatCompletion error: ", err);
        return;
    }

    fmt.Println(resp.Choices[0].Message.Content);
}

