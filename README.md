# GitHub Copilot Insights

## Objective

The GitHub Copilot Insights plugin provides detailed metrics and insights into the usage and adoption of GitHub Copilot within an organization or enterprise. It helps organizations understand how effectively they are utilizing GitHub Copilot, identify areas for improvement, and measure the impact on productivity and ROI.

## Insights

The plugin provides the following insights:

1. **Adoption & Utilization**
   - 🚀 **Seat Utilization Rate**: Measures how well the organization is using its purchased seats.
   - 📈 **Active vs. Engaged Users**: Tracks meaningful usage by comparing total engaged users to total active users.
   - 📊 **Feature Engagement Rate**: Highlights which features (chat, IDE completions, PR summaries) are driving value.
   - 🔄 **IDE Adoption**: Measures the adoption rate of IDE.
   - 🔄 **Dotcom Adoption**: Measures the adoption rate of using the contextual GitHub Copilot for GitHub hosted repositories.

2. **Productivity Impact**
   - ✅ **Code Acceptance Rate**: Tracks AI relevance and developer trust in suggestions.
   - 📏 **Code Adoption Efficiency**: Measures AI's direct contribution to production code.
   - 🤖 **AI Chat Engagement**: Determines if chat is enhancing workflows.

3. **ROI & Cost Efficiency**
   - 💰 **Cost per Engaged User**: Evaluates per-user ROI.
   - 🔄 **Custom Model Efficiency**: Tracks the impact of fine-tuned models vs. default AI.

4. **Workflow Acceleration**
   - ⚡ **PR Automation Impact**: Measures AI-driven automation in code review.
   - ⌛ **AI-Driven Code Speed**: Tracks how much Copilot accelerates dev cycles.

5. **Strategic Growth Metrics**
   - 📣 **Expansion Potential**: Gauges organic adoption growth.
   - 📌 **Editor Preference Index**: Identifies IDE preference trends (VSCode vs. JetBrains, Neovim).

## Installation

To install the GitHub Copilot Insights plugin as a GitHub CLI extension, follow these steps:

1. Install the extension:
   ```sh
   gh extension install chkp-roniz/gh-copilot-insights
   ```

## Usage

To use the GitHub Copilot Insights plugin, run the following command:

```sh
gh copilot-insights --scope <scope> --output <output> [--extended] [--debug]
```

- `--scope`: The name of the organization or enterprise for which to retrieve insights.
- `--output`: The output format, either `json`, `summary`, or `table`.
- `--extended`: Include extended metrics in the output (optional).
- `--debug`: Enable debug mode (optional).

## Example

Here is an example of how to use the plugin:

```sh
gh copilot-insights --scope my-org --output summary --extended --debug
```

This command retrieves the GitHub Copilot insights for the organization `my-org` and outputs a summary with extended metrics in debug mode.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
