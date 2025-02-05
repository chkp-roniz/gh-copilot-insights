package api

type CopilotBilling struct {
	Total int `json:"total_seats"`
	// SeatBreakdown         SeatBreakdown `json:"seat_breakdown"`
	// SeatManagementSetting string        `json:"seat_management_setting"`
	// IDEChat               string        `json:"ide_chat"`
	// PlatformChat          string        `json:"platform_chat"`
	// CLI                   string        `json:"cli"`
	// PublicCodeSuggestions string        `json:"public_code_suggestions"`
	// PlanType              string        `json:"plan_type"`
}

// type SeatBreakdown struct {
// 	Total               int `json:"total"`
// 	AddedThisCycle      int `json:"added_this_cycle"`
// 	PendingInvitation   int `json:"pending_invitation"`
// 	PendingCancellation int `json:"pending_cancellation"`
// 	ActiveThisCycle     int `json:"active_this_cycle"`
// 	InactiveThisCycle   int `json:"inactive_this_cycle"`
// }

type CopilotMetrics struct {
	Date                      string                `json:"date"`
	CopilotIDEChat            IDEChatMetrics        `json:"copilot_ide_chat"`
	TotalActiveUsers          int                   `json:"total_active_users"`
	CopilotDotcomChat         DotcomChatMetrics     `json:"copilot_dotcom_chat"`
	TotalEngagedUsers         int                   `json:"total_engaged_users"`
	CopilotDotcomPullRequests PullRequestMetrics    `json:"copilot_dotcom_pull_requests"`
	CopilotIDECodeCompletions CodeCompletionMetrics `json:"copilot_ide_code_completions"`
}

type IDEChatMetrics struct {
	Editors           []EditorMetrics `json:"editors"`
	TotalEngagedUsers int             `json:"total_engaged_users"`
}

type EditorMetrics struct {
	Name              string         `json:"name"`
	Models            []ModelMetrics `json:"models"`
	TotalEngagedUsers int            `json:"total_engaged_users"`
}

type ModelMetrics struct {
	Name                     string            `json:"name"`
	TotalChats               int               `json:"total_chats"`
	IsCustomModel            bool              `json:"is_custom_model"`
	TotalEngagedUsers        int               `json:"total_engaged_users"`
	TotalChatCopyEvents      int               `json:"total_chat_copy_events"`
	TotalChatInsertionEvents int               `json:"total_chat_insertion_events"`
	Languages                []LanguageMetrics `json:"languages"`
}

type LanguageMetrics struct {
	Name                    string `json:"name"`
	TotalEngagedUsers       int    `json:"total_engaged_users"`
	TotalCodeAcceptances    int    `json:"total_code_acceptances"`
	TotalCodeSuggestions    int    `json:"total_code_suggestions"`
	TotalCodeLinesAccepted  int    `json:"total_code_lines_accepted"`
	TotalCodeLinesSuggested int    `json:"total_code_lines_suggested"`
}

type DotcomChatMetrics struct {
	TotalEngagedUsers int `json:"total_engaged_users"`
}

type PullRequestMetrics struct {
	Repositories      []RepositoryMetrics `json:"repositories"`
	TotalEngagedUsers int                 `json:"total_engaged_users"`
}

type RepositoryMetrics struct {
	Name              string         `json:"name"`
	Models            []ModelMetrics `json:"models"`
	TotalEngagedUsers int            `json:"total_engaged_users"`
}

type CodeCompletionMetrics struct {
	Editors           []EditorMetrics   `json:"editors"`
	Languages         []LanguageMetrics `json:"languages"`
	TotalEngagedUsers int               `json:"total_engaged_users"`
}

type CopilotUsage struct {
	Day                   string `json:"day"`
	TotalSuggestionsCount int    `json:"total_suggestions_count"`
	TotalAcceptancesCount int    `json:"total_acceptances_count"`
	TotalLinesSuggested   int    `json:"total_lines_suggested"`
	TotalLinesAccepted    int    `json:"total_lines_accepted"`
	TotalActiveUsers      int    `json:"total_active_users"`
	TotalChatAcceptances  int    `json:"total_chat_acceptances"`
	TotalChatTurns        int    `json:"total_chat_turns"`
	TotalActiveChatUsers  int    `json:"total_active_chat_users"`
}

type Insight struct {
	ScopeName            string                      `json:"scope_name"`
	ScopeType            string                      `json:"scope_type"`
	AdoptionUtilization  AdoptionUtilizationMetrics  `json:"adoption_utilization"`
	ProductivityImpact   ProductivityImpactMetrics   `json:"productivity_impact"`
	ROICostEfficiency    ROICostEfficiencyMetrics    `json:"roi_cost_efficiency"`
	WorkflowAcceleration WorkflowAccelerationMetrics `json:"workflow_acceleration"`
	StrategicGrowth      StrategicGrowthMetrics      `json:"strategic_growth"`
}

type AdoptionUtilizationMetrics struct {
	SeatUtilizationRate   Metric            `json:"seat_utilization_rate"`
	ActiveVsEngagedUsers  Metric            `json:"active_vs_engaged_users"`
	FeatureEngagementRate map[string]Metric `json:"feature_engagement_rate"`
	IDEAdoption           Metric            `json:"ide_adoption"`
	DotcomAdoption        Metric            `json:"dotcom_adoption"`
}

type ProductivityImpactMetrics struct {
	CodeAcceptanceRate     Metric `json:"code_acceptance_rate"`
	CodeAdoptionEfficiency Metric `json:"code_adoption_efficiency"`
	AIChatEngagement       Metric `json:"ai_chat_engagement"`
}

type ROICostEfficiencyMetrics struct {
	CostPerEngagedUser    Metric `json:"cost_per_engaged_user"`
	CustomModelEfficiency Metric `json:"custom_model_efficiency"`
}

type WorkflowAccelerationMetrics struct {
	PRAutomationImpact Metric `json:"pr_automation_impact"`
	AIDrivenCodeSpeed  Metric `json:"ai_driven_code_speed"`
}

type StrategicGrowthMetrics struct {
	ExpansionPotential    Metric            `json:"expansion_potential"`
	EditorPreferenceIndex map[string]Metric `json:"editor_preference_index"`
}

type Metric struct {
	Value       float64 `json:"value"`
	DisplayName string  `json:"display_name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}
