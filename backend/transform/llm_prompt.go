package transform

import (
	"fmt"
	"strings"

	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/constants"
	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/pb"
)

func ConvertDelegateToPrompt(delegate *pb.DelegateInput) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Delegate Name: %s\n", strings.TrimSpace(delegate.GetDelegateName())))

	for _, param := range delegate.GetParameters() {
		sb.WriteString(fmt.Sprintf("%s: %.2f / %.2f\n", param.GetParameterName(), param.GetReceived(), param.GetHighest()))
	}
	return sb.String()
}

func GetFullPrompt(delegatesData string) string {
	prefix := constants.PromptPrefix
	suffix := constants.PromptSuffix

	return fmt.Sprintf("%s\n\n%s\n\n%s", prefix, delegatesData, suffix)
}

func ParseLLMFeedback(response string) []*pb.DelegateFeedback {
	var results []*pb.DelegateFeedback

	blocks := strings.Split(response, "\n\n")

	for _, block := range blocks {
		lines := strings.Split(block, "\n")
		var name, feedback string

		for _, line := range lines {
			if strings.HasPrefix(line, "Delegate Name:") {
				name = strings.TrimSpace(strings.TrimPrefix(line, "Delegate Name:"))
			} else if strings.HasPrefix(line, "Feedback:") {
				feedback = strings.TrimSpace(strings.TrimPrefix(line, "Feedback:"))
			}
		}

		if name != "" && feedback != "" {
			results = append(results, &pb.DelegateFeedback{
				DelegateName: name,
				FeedbackText: feedback,
			})
		}
	}

	return results
}
