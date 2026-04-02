package main

import (
	"fmt"
	"math/big"
	"os"
	"sort"
	"strings"
	"time"
)

func resultGatherer(context context) {
	context.log("10 result-gatherer")

	success := true

	successCount := 0
	failureCount := 0

	outputs := map[*big.Int]string{}

	for completedContext := range context.ResultGathererChannel {
		if completedContext.Err == nil {
			// SUCCESS
			successCount++
			outputs[completedContext.ID] +=
				fmt.Sprintf(
					"%s%s%s %s\n",
					Green,
					completedContext.SpecTriplet.String(),
					Reset,
					completedContext.SpecTriplet.Duration.String(),
				)

			if completedContext.ShowSubstitutions && len(completedContext.Substitutions) > 0 {
				outputs[completedContext.ID] +=
					substitutionsTable(completedContext.Substitutions)
			}
		} else {
			// FAILURE
			success = false
			failureCount++

			location := ""
			response := ""

			if completedContext.File == nil {
				// file open failure
				location += "[" + completedContext.Pathname + "]"
			} else {
				if completedContext.SpecTriplet == nil {
					// request/response parsing failure
					location += completedContext.File.String()
				} else {
					// request/response matching failure
					location +=
						completedContext.SpecTriplet.String() + " " +
							completedContext.SpecTriplet.Duration.String()

					if completedContext.SpecTriplet.ActualResponse != nil {
						response = completedContext.SpecTriplet.ActualResponse.String() + "\n"
					}
				}
			}

			if len(completedContext.Substitutions) > 0 {
				outputs[completedContext.ID] +=
					substitutionsTable(completedContext.Substitutions)
			}

			outputs[completedContext.ID] +=
				fmt.Sprintf(
					"%s%s%s %s\n%s\n",
					Red,
					location,
					Reset,
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
			Red,
			Reset,
			Green,
			successCount,
			Reset,
			Red,
			failureCount,
			Reset,
			duration.String(),
		)

		os.Exit(1)
	}

	fmt.Printf(
		"\n%sSUCCESS:%s %s+%d%s %s\n",
		Green,
		Reset,
		Green,
		successCount,
		Reset,
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

	b.WriteString(Yellow)

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

		b.WriteString(fmt.Sprintf("│ %-*s │ %-*s │\n", nameWidth, k, valueWidth, substitutions[k]))
	}

	b.WriteString("└")
	b.WriteString(strings.Repeat("─", nameWidth+2))
	b.WriteString("┴")
	b.WriteString(strings.Repeat("─", valueWidth+2))
	b.WriteString("┘\n")

	b.WriteString(Reset + "\n")

	return b.String()
}
