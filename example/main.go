package main

import (
	"fmt"
	"log"

	"github.com/batmac/go-alephalpha"
)

func main() {
	version, err := alephalpha.APIVersion()
	if err != nil {
		log.Print("Error getting API version: ", err)
	} else {
		fmt.Printf("API Version: %s\n", version)
	}

	c := alephalpha.NewClientFromEnv()

	models, err := c.ListAvailableModels()
	if err != nil {
		log.Print("Error getting available models: ", err)
	} else {
		for _, model := range models {
			fmt.Printf("Model %s:\n", model.Name)
			fmt.Printf("  Description: %s\n", model.Description)
			fmt.Printf("  Max Context Size: %d\n", model.MaxContextSize)
			fmt.Printf("  Hostings: %v\n", model.Hostings)
			fmt.Printf("  Image Support: %t\n", model.ImageSupport)
			fmt.Printf("  QA Support: %t\n", model.QASupport)
			fmt.Printf("  Summarization Support: %t\n", model.SummarizationSupport)
			fmt.Printf("  Embedding Types: %v\n", model.EmbeddingTypes)
			if model.MaximumCompletionTokens != nil {
				fmt.Printf("  Maximum Completion Tokens: %d\n", *model.MaximumCompletionTokens)
			} else {
				fmt.Println("  Maximum Completion Tokens: null")
			}
		}
	}

	settings, err := c.QuerySettings()
	if err != nil {
		log.Print("Error getting settings: ", err)
	} else {
		fmt.Printf("Settings:\n")
		fmt.Printf("  ID: %d\n", settings.ID)
		fmt.Printf("  Email: %s\n", settings.Email)
		fmt.Printf("  Role: %s\n", settings.Role)
		fmt.Printf("  Credits Remaining: %v\n", settings.CreditsRemaining)
		fmt.Printf("  Invoice Allowed: %t\n", settings.InvoiceAllowed)
		fmt.Printf("  Out of Credits Threshold: %v\n", settings.OutOfCreditsThreshold)
		fmt.Printf("  Terms of Service Version: %s\n", settings.TermsOfServiceVersion)
	}

	completion, err := c.Completion(alephalpha.CompletionRequest{
		Model:         alephalpha.ModelLuminousBase,
		Prompt:        "c'est quoi un capybara?",
		MaximumTokens: 1000,
	})
	if err != nil {
		log.Print("Error getting completion: ", err)
	} else {
		fmt.Printf("Model Version: %s\n", completion.ModelVersion)
		for i, c := range completion.Completions {
			fmt.Printf("Completion %d:\n", i+1)
			fmt.Printf("\tCompletion: %s\n", c.Completion)
			if c.RawCompletion != nil {
				fmt.Printf("\tRaw Completion: %s\n", *c.RawCompletion)
			}
			if len(c.CompletionTokens) > 0 {
				fmt.Printf("\tCompletion Tokens: %v\n", c.CompletionTokens)
			}
			fmt.Printf("\tFinish Reason: %s\n", c.FinishReason)
		}
		if completion.OptimizedPrompt != nil {
			fmt.Printf("Optimized Prompt: %v\n", completion.OptimizedPrompt)
		}
	}
}
