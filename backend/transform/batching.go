package transform

import (
	"strings"

	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/pb"
)

type BatchedPrompt struct {
	Prompt    string
	Delegates []*pb.DelegateInput
}

func BatchAndGeneratePrompts(delegates []*pb.DelegateInput, batchSize int) []BatchedPrompt {
	var result []BatchedPrompt

	for i := 0; i < len(delegates); i += batchSize {
		end := i + batchSize
		if end > len(delegates) {
			end = len(delegates)
		}

		batch := delegates[i:end]
		var sb strings.Builder

		for _, d := range batch {
			sb.WriteString(ConvertDelegateToPrompt(d))
			sb.WriteString("\n\n")
		}

		fullPrompt := GetFullPrompt(sb.String())
		result = append(result, BatchedPrompt{
			Prompt:    fullPrompt,
			Delegates: batch,
		})
	}

	return result
}
