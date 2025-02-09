package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/chkp-roniz/gh-copilot-insights/src/api"
	"github.com/chkp-roniz/gh-copilot-insights/src/usage"
	"github.com/sirupsen/logrus"
)

func main() {
	scope := flag.String("scope", "", "The name of the organization or enterprise for which to retrieve insights")
	output := flag.String("output", "json", "The output format, either 'json', 'summary', or 'table'")
	extended := flag.Bool("extended", false, "Include extended metrics in the output")
	debug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetOutput(os.Stdout)
		logrus.Debug("Debug mode enabled")
		logrus.Debugf("Scope: %s, Output: %s, Extended: %v", *scope, *output, *extended)
	}

	if *scope == "" {
		fmt.Println("Error: --scope is required")
		flag.Usage()
		os.Exit(1)
	}

	// Fetch Copilot usage insights
	usageData, err := api.FetchCopilotUsage(*scope)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"scope": *scope,
		}).Debugf("Error: %v", err)
		fmt.Println("Error fetching Copilot insights. Please try again.")
		os.Exit(1)
	}

	// Print output
	switch *output {
	case "json":
		usage.PrintJSON(usageData)
	case "summary":
		usage.PrintSummary(usageData, *extended)
	case "table":
		usage.PrintTable(usageData, *extended)
	default:
		fmt.Println("Invalid output format. Use 'json', 'summary', or 'table'.")
		os.Exit(1)
	}

	logrus.Debug("Execution completed")
}
