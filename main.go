package main

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

const FixedPrompt = `
I would like you to be an NLP tokenizer capable of performing multilingual sentiment analysis. 
I will send you a piece of text, and your task is to identify countries, tag words, tag parts of speech, count word frequency, perform sentiment analysis, and translate these words. I can only accept results in the specified JSON format. Other formats will not be accepted. Please remember to sort the "word" field in descending order according to word frequency. For fields starting with "translate," please remember to translate the language of the word after the underscore, i.e., the country code. For example, "translate_cn" should be translated into Chinese.
{
    "text":"Original text",
    "translate_cn":"Translated text",
    "language":"Chinese",
    "words":[
        {
            "word":"Word1",
            "count":1,
            "translation_cn":"Translation1"
        },
        {
            "word":"Word2",
            "count":2,
            "translation_cn":"Translation2"
        }
    ],
    "sentiment":"Sentiment analysis result",
    "aspects":[
        {
            "aspect":"Liked aspect1",
            "polarity":"Polarity1",
            "translation_cn":"aspect Translation1"
        },
        {
            "aspect":"Liked aspect2",
            "polarity":"Polarity2",
            "translation_cn":"aspect Translation2"
        }
    ]
}
Returning text other than json is not allowed.
My first problem is "%s".
`

func main() {

	app := fiber.New()

	client := resty.New()
	client.SetDebug(true)
	client.SetError(&ErrorResponse{})
	client.SetHeader("Content-Type", "application/json")

	token := os.Getenv("OPENAI_API_KEY")
	if token == "" {
		panic("Missing `OPENAI_API_KEY` environment variable.")
	}

	client.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))

	app.Post("/api/tokenizer", func(c *fiber.Ctx) error {

		inputText := c.FormValue("input_text")
		if inputText == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Missing input_text parameter",
			})
		}

		body := ChatCompletionRequest{
			Model: "gpt-3.5-turbo",
			Messages: []ChatCompletionMessage{
				{
					Role:    "assistant",
					Content: fmt.Sprintf(FixedPrompt, inputText),
				},
			},
		}

		var result ChatCompletionResponse
		resp, err := client.R().
			SetBody(body).SetResult(&result).
			Post("https://api.openai.com/v1/chat/completions")

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if resp.IsError() {
			if errorResponse, ok := resp.Error().(*ErrorResponse); ok {
				return c.JSON(errorResponse)
			}
			return c.JSON(resp)
		}

		if result.Choices != nil && len(result.Choices) > 0 && result.Choices[0].Message.Content != "" {
			if result.Choices[0].FinishReason == "length" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error":  "The returned content is incomplete.",
					"result": result,
				})
			}
			var response Response
			err = jsoniter.Unmarshal([]byte(result.Choices[0].Message.Content), &response)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			return c.JSON(response)
		}

		return c.JSON(resp)
	})

	_ = app.Listen(":3000")
}
