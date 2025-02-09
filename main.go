package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/chkp-roniz/gh-copilot-insights/src/api"
	"github.com/chkp-roniz/gh-copilot-insights/src/usage"
	logger "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

func main() {
	scope := flag.String("scope", "", "The name of the organization or enterprise for which to retrieve insights")
	output := flag.String("output", "json", "The output format, either 'json', 'summary', or 'table'")
	extended := flag.Bool("extended", false, "Include extended metrics in the output")
	debug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	if *debug {
		logger.SetLevel(logger.DebugLevel)
		logger.SetFormatter(&easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "%time% [%lvl%]: %msg%\n",
		})
		logger.SetOutput(os.Stdout)
		logger.Debug("Debug mode enabled")
		logger.Debugf("Scope: %s, Output: %s, Extended: %v", *scope, *output, *extended)
	}

	if *scope == "" {
		fmt.Println("Error: --scope is required")
		flag.Usage()
		os.Exit(1)
	}

	// Fetch Copilot usage insights
	usageData, err := api.FetchCopilotUsage(*scope)
	if err != nil {
		logger.WithFields(logger.Fields{
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

	logger.Debug("Execution completed")
}
