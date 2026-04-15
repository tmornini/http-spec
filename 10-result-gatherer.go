package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func resultGatherer(context context) {
	context.enterStage("10 result-gatherer")

	success := true

	successCount := 0
	failureCount := 0

	outputs := map[string]string{}

	for completedContext := range context.ResultGathererChannel {
		if completedContext.Err == nil {
			// SUCCESS
			successCount++
			outputs[completedContext.ID] +=
				fmt.Sprintf(
					"%s%s%s %s\n",
					colorGreen,
					completedContext.SpecTriplet.String(),
					colorReset,
					completedContext.SpecTriplet.Duration.String(),
				)

			if completedContext.ShowSubstitutions &&
				len(completedContext.Substitutions) > 0 {
				outputs[completedContext.ID] +=
					substitutionsTable(completedContext.Substitutions)
			}
		} else {
			// FAILURE
			success = false
			failureCount++

			location := ""
			response := ""

			switch {
			case strings.HasPrefix(completedContext.Stage, "02"):
				location = "[" + completedContext.Pathname + "]"
			case strings.HasPrefix(completedContext.Stage, "03"):
				location = completedContext.File.String()
			default:
				location =
					completedContext.SpecTriplet.String() + " " +
						completedContext.SpecTriplet.Duration.String()

				if completedContext.SpecTriplet.ActualResponse != nil {
					response = completedContext.SpecTriplet.ActualResponse.String() + "\n"
				}
			}

			if len(completedContext.Substitutions) > 0 {
				outputs[completedContext.ID] +=
					substitutionsTable(completedContext.Substitutions)
			}

			outputs[completedContext.ID] +=
				fmt.Sprintf(
					"%s%s%s %s\n%s\n",
					colorRed,
					location,
					colorReset,
					completedContext.Err.Error(),
					response,
				)
		}
	}

	duration := time.Since(context.StartedAt)

	fmt.Println()

	for _, result := range outputs {
		fmt.Print(result)
	}

	if !success {
		fmt.Printf(
			"%sFAILURE:%s %s+%d%s %s-%d%s %s\n",
			colorRed,
			colorReset,
			colorGreen,
			successCount,
			colorReset,
			colorRed,
			failureCount,
			colorReset,
			duration.String(),
		)

		os.Exit(1)
	}

	fmt.Printf(
		"\n%sSUCCESS:%s %s+%d%s %s\n",
		colorGreen,
		colorReset,
		colorGreen,
		successCount,
		colorReset,
		duration.String(),
	)

	os.Exit(0)
}

func substitutionsTable(substitutions map[string]string) string {
	keys := make([]string, 0, len(substitutions))

	for k := range substitutions {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	nameWidth := 0
	valueWidth := 0

	for _, k := range keys {
		if len(k) > nameWidth {
			nameWidth = len(k)
		}

		if len(substitutions[k]) > valueWidth {
			valueWidth = len(substitutions[k])
		}
	}

	var b strings.Builder

	b.WriteString(colorYellow)

	b.WriteString("Substitutions\n")

	b.WriteString("┌")
	b.WriteString(strings.Repeat("─", nameWidth+2))
	b.WriteString("┬")
	b.WriteString(strings.Repeat("─", valueWidth+2))
	b.WriteString("┐\n")

	divider := "├" +
		strings.Repeat("─", nameWidth+2) +
		"┼" +
		strings.Repeat("─", valueWidth+2) +
		"┤\n"

	for i, k := range keys {
		if i > 0 {
			b.WriteString(divider)
		}

		b.WriteString(
			fmt.Sprintf(
				"│ %-*s │ %-*s │\n",
				nameWidth, k,
				valueWidth, substitutions[k],
			),
		)
	}

	b.WriteString("└")
	b.WriteString(strings.Repeat("─", nameWidth+2))
	b.WriteString("┴")
	b.WriteString(strings.Repeat("─", valueWidth+2))
	b.WriteString("┘\n")

	b.WriteString(colorReset + "\n")

	return b.String()
}
