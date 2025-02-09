package usage

import (
	"encoding/json"
	"fmt"
	"math"
	"os"

	"github.com/chkp-roniz/gh-copilot-insights/src/api"
	"github.com/olekukonko/tablewriter"
)

func PrintJSON(insights []api.Insight) {
	data, err := json.MarshalIndent(insights, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}

func toPercentage(value float64) string {
	return fmt.Sprintf("%.0f%%", math.Min(value, 1)*100)
}

func printMetric(category, displayName, description string, value float64) {
	fmt.Printf("## %s\n\n", category)
	fmt.Printf("**%s**: %s\n", displayName, toPercentage(value))
	fmt.Printf("%s\n\n", description)
}

func appendMetric(table *tablewriter.Table, category, displayName, description string, value float64) {
	table.Append([]string{category, displayName, toPercentage(value), description})
}

func PrintSummary(insights []api.Insight, extended bool) {
	for _, insight := range insights {
		if len(insights) > 0 {
			fmt.Printf("# GitHub Copilot Insights for %s (%s)\n\n", insight.ScopeName, insight.ScopeType)
		}
		printMetric(insight.AdoptionUtilization.SeatUtilizationRate.Category, insight.AdoptionUtilization.SeatUtilizationRate.DisplayName, insight.AdoptionUtilization.SeatUtilizationRate.Description, insight.AdoptionUtilization.SeatUtilizationRate.Value)
		printMetric(insight.AdoptionUtilization.ActiveVsEngagedUsers.Category, insight.AdoptionUtilization.ActiveVsEngagedUsers.DisplayName, insight.AdoptionUtilization.ActiveVsEngagedUsers.Description, insight.AdoptionUtilization.ActiveVsEngagedUsers.Value)
		if extended {
			for feature, metric := range insight.AdoptionUtilization.FeatureEngagementRate {
				printMetric(metric.Category, fmt.Sprintf("%s (%s)", metric.DisplayName, feature), metric.Description, metric.Value)
			}
		}
		printMetric(insight.AdoptionUtilization.IDEAdoption.Category, insight.AdoptionUtilization.IDEAdoption.DisplayName, insight.AdoptionUtilization.IDEAdoption.Description, insight.AdoptionUtilization.IDEAdoption.Value)
		printMetric(insight.AdoptionUtilization.DotcomAdoption.Category, insight.AdoptionUtilization.DotcomAdoption.DisplayName, insight.AdoptionUtilization.DotcomAdoption.Description, insight.AdoptionUtilization.DotcomAdoption.Value)

		printMetric(insight.ProductivityImpact.CodeAcceptanceRate.Category, insight.ProductivityImpact.CodeAcceptanceRate.DisplayName, insight.ProductivityImpact.CodeAcceptanceRate.Description, insight.ProductivityImpact.CodeAcceptanceRate.Value)
		printMetric(insight.ProductivityImpact.CodeAdoptionEfficiency.Category, insight.ProductivityImpact.CodeAdoptionEfficiency.DisplayName, insight.ProductivityImpact.CodeAdoptionEfficiency.Description, insight.ProductivityImpact.CodeAdoptionEfficiency.Value)
		printMetric(insight.ProductivityImpact.AIChatEngagement.Category, insight.ProductivityImpact.AIChatEngagement.DisplayName, insight.ProductivityImpact.AIChatEngagement.Description, insight.ProductivityImpact.AIChatEngagement.Value)

		printMetric(insight.ROICostEfficiency.CostPerEngagedUser.Category, insight.ROICostEfficiency.CostPerEngagedUser.DisplayName, insight.ROICostEfficiency.CostPerEngagedUser.Description, insight.ROICostEfficiency.CostPerEngagedUser.Value)
		printMetric(insight.ROICostEfficiency.CustomModelEfficiency.Category, insight.ROICostEfficiency.CustomModelEfficiency.DisplayName, insight.ROICostEfficiency.CustomModelEfficiency.Description, insight.ROICostEfficiency.CustomModelEfficiency.Value)

		// printMetric(insight.WorkflowAcceleration.PRAutomationImpact.Category, insight.WorkflowAcceleration.PRAutomationImpact.DisplayName, insight.WorkflowAcceleration.PRAutomationImpact.Description, insight.WorkflowAcceleration.PRAutomationImpact.Value)
		// printMetric(insight.WorkflowAcceleration.AIDrivenCodeSpeed.Category, insight.WorkflowAcceleration.AIDrivenCodeSpeed.DisplayName, insight.WorkflowAcceleration.AIDrivenCodeSpeed.Description, insight.WorkflowAcceleration.AIDrivenCodeSpeed.Value)

		if extended {
			fmt.Printf("## 📣 %s\n\n", insight.StrategicGrowth.ExpansionPotential.Category)
			// printMetric(insight.StrategicGrowth.ExpansionPotential.Category, insight.StrategicGrowth.ExpansionPotential.DisplayName, insight.StrategicGrowth.ExpansionPotential.Description, insight.StrategicGrowth.ExpansionPotential.Value)
			for editor, metric := range insight.StrategicGrowth.EditorPreferenceIndex {
				printMetric(metric.Category, fmt.Sprintf("%s (%s)", metric.DisplayName, editor), metric.Description, metric.Value)
			}
		}
	}
}

func PrintTable(insights []api.Insight, extended bool) {
	table := tablewriter.NewWriter(os.Stdout)
	headers := []string{"Category", "Metric", "Value", "Description"}
	if extended {
		headers = append(headers, "Extended Info")
	}
	table.SetHeader(headers)

	for _, insight := range insights {
		if len(insights) > 0 {
			fmt.Printf("# GitHub Copilot Insights for %s (%s)\n\n", insight.ScopeName, insight.ScopeType)
		}
		appendMetric(table, "🚀 "+insight.AdoptionUtilization.SeatUtilizationRate.Category, insight.AdoptionUtilization.SeatUtilizationRate.DisplayName, insight.AdoptionUtilization.SeatUtilizationRate.Description, insight.AdoptionUtilization.SeatUtilizationRate.Value)
		appendMetric(table, "🚀 "+insight.AdoptionUtilization.ActiveVsEngagedUsers.Category, insight.AdoptionUtilization.ActiveVsEngagedUsers.DisplayName, insight.AdoptionUtilization.ActiveVsEngagedUsers.Description, insight.AdoptionUtilization.ActiveVsEngagedUsers.Value)
		appendMetric(table, "🚀 "+insight.AdoptionUtilization.IDEAdoption.Category, insight.AdoptionUtilization.IDEAdoption.DisplayName, insight.AdoptionUtilization.IDEAdoption.Description, insight.AdoptionUtilization.IDEAdoption.Value)
		appendMetric(table, "🚀 "+insight.AdoptionUtilization.DotcomAdoption.Category, insight.AdoptionUtilization.DotcomAdoption.DisplayName, insight.AdoptionUtilization.DotcomAdoption.Description, insight.AdoptionUtilization.DotcomAdoption.Value)
		if extended {
			for feature, metric := range insight.AdoptionUtilization.FeatureEngagementRate {
				appendMetric(table, "🚀 "+metric.Category, fmt.Sprintf("%s (%s)", metric.DisplayName, feature), metric.Description, metric.Value)
			}
		}
		appendMetric(table, "🤖 "+insight.ProductivityImpact.CodeAcceptanceRate.Category, insight.ProductivityImpact.CodeAcceptanceRate.DisplayName, insight.ProductivityImpact.CodeAcceptanceRate.Description, insight.ProductivityImpact.CodeAcceptanceRate.Value)
		appendMetric(table, "🤖 "+insight.ProductivityImpact.CodeAdoptionEfficiency.Category, insight.ProductivityImpact.CodeAdoptionEfficiency.DisplayName, insight.ProductivityImpact.CodeAdoptionEfficiency.Description, insight.ProductivityImpact.CodeAdoptionEfficiency.Value)
		appendMetric(table, "🤖 "+insight.ProductivityImpact.AIChatEngagement.Category, insight.ProductivityImpact.AIChatEngagement.DisplayName, insight.ProductivityImpact.AIChatEngagement.Description, insight.ProductivityImpact.AIChatEngagement.Value)
		appendMetric(table, "💰 "+insight.ROICostEfficiency.CostPerEngagedUser.Category, insight.ROICostEfficiency.CostPerEngagedUser.DisplayName, insight.ROICostEfficiency.CostPerEngagedUser.Description, insight.ROICostEfficiency.CostPerEngagedUser.Value)
		appendMetric(table, "💰 "+insight.ROICostEfficiency.CustomModelEfficiency.Category, insight.ROICostEfficiency.CustomModelEfficiency.DisplayName, insight.ROICostEfficiency.CustomModelEfficiency.Description, insight.ROICostEfficiency.CustomModelEfficiency.Value)
		appendMetric(table, "⚡ "+insight.WorkflowAcceleration.PRAutomationImpact.Category, insight.WorkflowAcceleration.PRAutomationImpact.DisplayName, insight.WorkflowAcceleration.PRAutomationImpact.Description, insight.WorkflowAcceleration.PRAutomationImpact.Value)
		// appendMetric(table, "⚡ "+insight.WorkflowAcceleration.AIDrivenCodeSpeed.Category, insight.WorkflowAcceleration.AIDrivenCodeSpeed.DisplayName, insight.WorkflowAcceleration.AIDrivenCodeSpeed.Description, insight.WorkflowAcceleration.AIDrivenCodeSpeed.Value)
		// appendMetric(table, "📣 "+insight.StrategicGrowth.ExpansionPotential.Category, insight.StrategicGrowth.ExpansionPotential.DisplayName, insight.StrategicGrowth.ExpansionPotential.Description, insight.StrategicGrowth.ExpansionPotential.Value)
		if extended {
			for editor, metric := range insight.StrategicGrowth.EditorPreferenceIndex {
				appendMetric(table, "📣 "+metric.Category, fmt.Sprintf("%s (%s)", metric.DisplayName, editor), metric.Description, metric.Value)
			}
		}
	}

	table.Render()
}
