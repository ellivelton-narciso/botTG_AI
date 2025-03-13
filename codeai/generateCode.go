package codeai

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	openai "github.com/sashabaranov/go-openai"
)

func GenerateCode(prompt, apiKey string, debug bool) string {
	client := openai.NewClient(apiKey)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: `
						Você é um desenvolvedor experiente especializado em gerar código de alta qualidade. 
						Responda APENAS com o código solicitado, sem introduções ou explicações fora do código. 
						Inclua comentários claros e concisos dentro do código explicando a lógica e funcionalidades principais. 
						O código deve seguir as melhores práticas de legibilidade, eficiência e segurança. 
						Forneça soluções completas e funcionais que possam ser executadas diretamente.
					`,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		log.Printf("Erro ao gerar código: %v", err)
		return "Erro ao gerar código"
	}

	if debug {
		jsonBytes, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			log.Println("Erro ao converter resposta para JSON:", err)
		} else {
			fmt.Println(string(jsonBytes))
		}
	}

	return resp.Choices[0].Message.Content
}
