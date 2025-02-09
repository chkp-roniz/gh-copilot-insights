package api

import (
	"fmt"

	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/auth"

	// "github.com/cli/go-gh/pkg/api"
	"github.com/sirupsen/logrus"
)

func determineEndpoint(scope string) (string, error) {
	// stdOut, stdErr, err := gh.Exec("auth", "token")
	// if err != nil {
	// 	logrus.Debugf("Error retriving access token from GitHub Copilot: %v", err)
	// 	logrus.Debugf("%s", stdErr.String())
	// 	return "", err
	// }
	// opts := &api.ClientOptions{
	// 	AuthToken: stdOut.String(),
	// }
	client, err := gh.RESTClient(nil)
	token, source := auth.TokenForHost("github.com")
	logrus.Debugf("Token: %s, Source: %s", token, source)
	if err != nil {
		logrus.Debugf("Error creating REST client: %v", err)
		return "", err
	}

	// Check if the scope is an organization
	orgEndpoint := fmt.Sprintf("orgs/%s", scope)
	var orgResponse map[string]interface{}
	err1 := client.Get(orgEndpoint, &orgResponse)
	if err1 == nil && orgResponse != nil {
		return "orgs", nil
	}

	// Check if the scope is an enterprise
	enterpriseEndpoint := fmt.Sprintf("enterprises/%s/properties/schema", scope)
	var enterpriseResponse []map[string]interface{}
	err2 := client.Get(enterpriseEndpoint, &enterpriseResponse)
	if err2 == nil && enterpriseResponse != nil {
		return "enterprises", nil
	}

	logrus.Debugf("Invalid scope: %s", scope)
	logrus.Debugf("Error orgs: %v", err1)
	logrus.Debugf("Error enterprises: %v", err1)
	return "", fmt.Errorf("invalid scope: %s", scope)
}

///
/// 1. Adoption & Utilization
// 🚀 Seat Utilization Rate = Active Users This Cycle / Total Paid Seats

// Insight: Measures how well the organization is using its purchased seats.
// Action: If low, investigate onboarding gaps or unused licenses.
// 📈 Active vs. Engaged Users = Total Engaged Users / Total Active Users

// Insight: Not all active users engage deeply. Tracks meaningful usage.
// Action: Identify under-engaged users and provide training.
// 📊 Feature Engagement Rate = Users of Feature X / Total Engaged Users

// Insight: Highlights which features (chat, IDE completions, PR summaries) are driving value.
// Action: Promote underused high-impact features.
// 2. Productivity Impact
// ✅ Code Acceptance Rate = Total Acceptances / Total Suggestions

// Insight: Tracks AI relevance and developer trust in suggestions.
// Action: If low, assess prompt quality or custom model need.
// 📏 Code Adoption Efficiency = Total Code Lines Accepted / Total Code Lines Suggested

// Insight: Measures AI's direct contribution to production code.
// Action: If low, check for friction in adoption (e.g., formatting, relevance).
// 🤖 AI Chat Engagement = Chat Users / Total Engaged Users

// Insight: Determines if chat is enhancing workflows.
// Action: If low, educate teams on use cases.
// 3. ROI & Cost Efficiency
// 💰 Cost per Engaged User = Total Paid Cost / Total Engaged Users

// Insight: Evaluates per-user ROI.
// Action: If high, optimize licenses or engagement strategies.
// 🔄 Custom Model Efficiency = Custom Model Users / Default Model Users

// Insight: Tracks the impact of fine-tuned models vs. default AI.
// Action: If high, consider scaling custom models.
// 4. Workflow Acceleration
// ⚡ PR Automation Impact = PR Summaries Created / PRs Opened

// Insight: Measures AI-driven automation in code review.
// Action: If low, encourage teams to use PR summaries.
// ⌛ AI-Driven Code Speed = Time to Merge (w/ AI) / Time to Merge (w/o AI)

// Insight: Tracks how much Copilot accelerates dev cycles.
// Action: If impact is low, analyze blockers (e.g., review friction).
// 5. Strategic Growth Metrics
// 📣 Expansion Potential = New Users Added / Total Users

// Insight: Gauges organic adoption growth.
// Action: If low, reassess AI adoption strategy.
// 📌 Editor Preference Index = Users per Editor / Total Users

// Insight: Identifies IDE preference trends (VSCode vs. JetBrains, Neovim).
// Action: Optimize support/training per IDE.

func getInsights(scopeName, scopeType string, usage []CopilotUsage, metrics []CopilotMetrics, billing CopilotBilling) Insight {
	var totalActiveUsers, totalActiveUsersMetrics, totalEngagedUsers, totalSuggestions, totalAcceptances, totalLinesSuggested, totalLinesAccepted int
	var totalIDEUsers, totalDotcomUsers int
	featureEngagementRate := make(map[string]float64)
	editorPreferenceIndex := make(map[string]float64)

	for _, u := range usage {
		totalActiveUsers += u.TotalActiveUsers
		totalSuggestions += u.TotalSuggestionsCount
		totalAcceptances += u.TotalAcceptancesCount
		totalLinesSuggested += u.TotalLinesSuggested
		totalLinesAccepted += u.TotalLinesAccepted
	}

	for _, m := range metrics {
		totalEngagedUsers += m.TotalEngagedUsers
		totalActiveUsersMetrics += m.TotalActiveUsers
		totalIDEUsers += m.CopilotIDEChat.TotalEngagedUsers + m.CopilotIDECodeCompletions.TotalEngagedUsers
		totalDotcomUsers += m.CopilotDotcomChat.TotalEngagedUsers + m.CopilotDotcomPullRequests.TotalEngagedUsers
		for _, editor := range m.CopilotIDECodeCompletions.Editors {
			editorPreferenceIndex[editor.Name] += float64(editor.TotalEngagedUsers)
		}
		for _, feature := range []struct {
			name  string
			value int
		}{
			{"ide_chat", m.CopilotIDEChat.TotalEngagedUsers},
			{"dotcom_chat", m.CopilotDotcomChat.TotalEngagedUsers},
			{"pull_requests", m.CopilotDotcomPullRequests.TotalEngagedUsers},
		} {
			featureEngagementRate[feature.name] += float64(feature.value)
		}
	}

	avgEngagedUsers := float64(totalEngagedUsers) / float64(len(metrics))

	seatUtilizationRate := float64(avgEngagedUsers) / float64(billing.Total)
	activeVsEngagedUsers := float64(totalEngagedUsers) / float64(totalActiveUsersMetrics)
	codeAcceptanceRate := float64(totalAcceptances) / float64(totalSuggestions)
	codeAdoptionEfficiency := float64(totalLinesAccepted) / float64(totalLinesSuggested)
	costPerEngagedUser := float64(billing.Total) / float64(totalEngagedUsers)
	ideAdoption := float64(totalIDEUsers) / float64(totalEngagedUsers)
	dotcomAdoption := float64(totalDotcomUsers) / float64(totalEngagedUsers)

	for key := range featureEngagementRate {
		featureEngagementRate[key] /= float64(totalEngagedUsers)
	}

	for key := range editorPreferenceIndex {
		editorPreferenceIndex[key] /= float64(totalEngagedUsers)
	}

	return Insight{
		ScopeName: scopeName,
		ScopeType: scopeType,
		AdoptionUtilization: AdoptionUtilizationMetrics{
			SeatUtilizationRate: Metric{
				Value:       seatUtilizationRate,
				DisplayName: "Seat Utilization Rate",
				Description: "Measures how well the organization is using its purchased seats. Calculated as Engaged Users This Cycle / Total Paid Seats.",
				Category:    "Adoption & Utilization",
			},
			ActiveVsEngagedUsers: Metric{
				Value:       activeVsEngagedUsers,
				DisplayName: "Active vs. Engaged Users",
				Description: "Tracks meaningful usage by comparing total engaged users to total active users. Calculated as Total Engaged Users / Total Active Users.",
				Category:    "Adoption & Utilization",
			},
			IDEAdoption: Metric{
				Value:       ideAdoption,
				DisplayName: "IDE Adoption",
				Description: "Measures the adoption rate of IDE. Calculated as Total IDE Users / Total Engaged Users.",
				Category:    "Adoption & Utilization",
			},
			DotcomAdoption: Metric{
				Value:       dotcomAdoption,
				DisplayName: "Dotcom Adoption",
				Description: "Measures the adoption rate of using the contextual GitHub Copilot for GitHub hosted repositories. Calculated as Total Dotcom Users / Total Engaged Users.",
				Category:    "Adoption & Utilization",
			},
			FeatureEngagementRate: map[string]Metric{
				"ide_chat": {
					Value:       featureEngagementRate["ide_chat"],
					DisplayName: "Feature Engagement Rate",
					Description: "Highlights which features (chat, IDE completions, PR summaries) are driving value. Calculated as Users of Feature X / Total Engaged Users.",
					Category:    "Adoption & Utilization",
				},
				"dotcom_chat": {
					Value:       featureEngagementRate["dotcom_chat"],
					DisplayName: "Feature Engagement Rate",
					Description: "Highlights which features (chat, IDE completions, PR summaries) are driving value. Calculated as Users of Feature X / Total Engaged Users.",
					Category:    "Adoption & Utilization",
				},
				"pull_requests": {
					Value:       featureEngagementRate["pull_requests"],
					DisplayName: "Feature Engagement Rate",
					Description: "Highlights which features (chat, IDE completions, PR summaries) are driving value. Calculated as Users of Feature X / Total Engaged Users.",
					Category:    "Adoption & Utilization",
				},
			},
		},
		ProductivityImpact: ProductivityImpactMetrics{
			CodeAcceptanceRate: Metric{
				Value:       codeAcceptanceRate,
				DisplayName: "Code Acceptance Rate",
				Description: "Tracks AI relevance and developer trust in suggestions. Calculated as Total Acceptances / Total Suggestions.",
				Category:    "Productivity Impact",
			},
			CodeAdoptionEfficiency: Metric{
				Value:       codeAdoptionEfficiency,
				DisplayName: "Code Adoption Efficiency",
				Description: "Measures AI's direct contribution to production code. Calculated as Total Code Lines Accepted / Total Code Lines Suggested.",
				Category:    "Productivity Impact",
			},
			AIChatEngagement: Metric{
				Value:       featureEngagementRate["ide_chat"],
				DisplayName: "AI Chat Engagement",
				Description: "Determines if chat is enhancing workflows. Calculated as Chat Users / Total Engaged Users.",
				Category:    "Productivity Impact",
			},
		},
		ROICostEfficiency: ROICostEfficiencyMetrics{
			CostPerEngagedUser: Metric{
				Value:       costPerEngagedUser,
				DisplayName: "Cost per Engaged User",
				Description: "Evaluates per-user ROI. Calculated as Total Paid Cost / Total Engaged Users.",
				Category:    "ROI & Cost Efficiency",
			},
			CustomModelEfficiency: Metric{
				Value:       featureEngagementRate["custom_model"],
				DisplayName: "Custom Model Efficiency",
				Description: "Tracks the impact of fine-tuned models vs. default AI. Calculated as Custom Model Users / Default Model Users.",
				Category:    "ROI & Cost Efficiency",
			},
		},
		WorkflowAcceleration: WorkflowAccelerationMetrics{
			PRAutomationImpact: Metric{
				Value:       featureEngagementRate["pull_requests"],
				DisplayName: "PR Automation Impact",
				Description: "Measures AI-driven automation in code review. Calculated as PR Summaries Created / PRs Opened.",
				Category:    "Workflow Acceleration",
			},
			AIDrivenCodeSpeed: Metric{
				Value:       0, // Placeholder, requires additional data
				DisplayName: "AI-Driven Code Speed",
				Description: "Tracks how much Copilot accelerates dev cycles. Calculated as Time to Merge (w/ AI) / Time to Merge (w/o AI).",
				Category:    "Workflow Acceleration",
			},
		},
		StrategicGrowth: StrategicGrowthMetrics{
			ExpansionPotential: Metric{
				Value:       featureEngagementRate["new_users"],
				DisplayName: "Expansion Potential",
				Description: "Gauges organic adoption growth. Calculated as New Users Added / Total Users.",
				Category:    "Strategic Growth",
			},
			EditorPreferenceIndex: map[string]Metric{
				"vscode": {
					Value:       editorPreferenceIndex["vscode"],
					DisplayName: "Editor Preference Index",
					Description: "Identifies IDE preference trends (VSCode vs. JetBrains, Neovim). Calculated as Users per Editor / Total Users.",
					Category:    "Strategic Growth",
				},
				"jetbrains": {
					Value:       editorPreferenceIndex["jetbrains"],
					DisplayName: "Editor Preference Index",
					Description: "Identifies IDE preference trends (VSCode vs. JetBrains, Neovim). Calculated as Users per Editor / Total Users.",
					Category:    "Strategic Growth",
				},
			},
		},
	}
}

func FetchCopilotUsage(scopeName string) ([]Insight, error) {
	// stdOut, stdErr, err := gh.Exec("auth", "token")
	// if err != nil {
	// 	logrus.Debugf("Error retriving access token from GitHub Copilot: %v", err)
	// 	logrus.Debugf("%s", stdErr.String())
	// 	return nil, err
	// }
	// opts := &api.ClientOptions{
	// 	AuthToken: stdOut.String(),
	// }
	client, err := gh.RESTClient(nil)
	if err != nil {
		logrus.Debugf("Error creating REST client: %v", err)
		return nil, err
	}

	scopeType, err := determineEndpoint(scopeName)
	if err != nil {
		logrus.Debugf("Error determining endpoint for scope %s: %v", scopeName, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s/copilot", scopeType, scopeName)

	var usage []CopilotUsage
	err = client.Get(fmt.Sprintf("%s/usage", endpoint), &usage)
	if err != nil {
		logrus.Debugf("Error fetching usage data from endpoint %s: %v", endpoint, err)
		return nil, err
	}

	var metrics []CopilotMetrics
	err = client.Get(fmt.Sprintf("%s/metrics", endpoint), &metrics)
	if err != nil {
		logrus.Debugf("Error fetching metrics data from endpoint %s: %v", endpoint, err)
		return nil, err
	}

	var billing CopilotBilling
	err = client.Get(fmt.Sprintf("%s/billing/seats", endpoint), &billing)
	if err != nil {
		logrus.Debugf("Error fetching billing data from endpoint %s: %v", endpoint, err)
		return nil, err
	}

	insight := getInsights(scopeName, scopeType, usage, metrics, billing)
	return []Insight{insight}, nil
}
