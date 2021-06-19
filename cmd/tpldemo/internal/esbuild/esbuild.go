package esbuild

import (
	"fmt"

	esbuildapi "github.com/evanw/esbuild/pkg/api"
)

func Build(entrypoint, outfile string) esbuildapi.BuildResult {
	return esbuildapi.Build(esbuildapi.BuildOptions{
		EntryPoints:       []string{entrypoint},
		Outfile:           outfile,
		Sourcemap:         esbuildapi.SourceMapLinked,
		Write:             false,
		Bundle:            true,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		Platform:          esbuildapi.PlatformBrowser,
	})
}

func PrintMessages(result esbuildapi.BuildResult) {
	terminalWidth := getTerminalWidth()
	var messages = make([]string, 0, len(result.Errors)+len(result.Warnings))
	messages = append(
		messages,
		esbuildapi.FormatMessages(result.Warnings, esbuildapi.FormatMessagesOptions{
			TerminalWidth: terminalWidth,
			Kind:          esbuildapi.WarningMessage,
			Color:         true,
		})...,
	)
	messages = append(
		messages,
		esbuildapi.FormatMessages(result.Errors, esbuildapi.FormatMessagesOptions{
			TerminalWidth: terminalWidth,
			Kind:          esbuildapi.ErrorMessage,
			Color:         true,
		})...,
	)
	for _, message := range messages {
		fmt.Println(message)
	}
}
