package cli

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mnsdojo/gofetch/internal/ascii"
	"github.com/mnsdojo/gofetch/internal/battery"
	"github.com/mnsdojo/gofetch/internal/system"
)

var (
	labelColor = lipgloss.Color("145")
	style      = lipgloss.NewStyle().
			Bold(true).
			Padding(5).
			Border(lipgloss.ThickBorder(), true, true, true, true).Align(lipgloss.Center).
			BorderForeground(lipgloss.Color("228"))

	asciiStyle = lipgloss.NewStyle().
			Align(lipgloss.Center)

	separatorStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(lipgloss.Color("201")).
			MarginTop(1).
			MarginBottom(1)

	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("228")).Align(lipgloss.Center)

	labelStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(labelColor).
			Width(15)

	valueStyle = lipgloss.NewStyle().Bold(true).
			Width(30)
)

// genOutput generates the formatted system information output
func GenOutput() (string, error) {
	info, err := fetchSystemInfo()
	if err != nil {
		return "", fmt.Errorf("failed to fetch system info: %w", err)
	}

	char, err := fetchASCIIArt()
	if err != nil {
		return "", fmt.Errorf("failed to fetch ASCII art: %w", err)
	}

	batteryStatus, err := fetchBatteryStatus()
	if err != nil {
		return "", fmt.Errorf("failed to fetch battery status: %w", err)
	}

	header := headerStyle.Render("gofetch")

	labels := []string{"Username:", "Battery:", "OS Release:", "Kernel:", "Default Shell:", "Uptime:", "Memory Usage:"}
	values := []string{info.UserName, batteryStatus, info.OSRelease, info.Kernel, info.DefaultShell, info.Uptime, info.MemoryUsage}

	output := formatOutput(labels, values)
	separator := separatorStyle.Render(strings.Repeat("─", 50))
	asciiArt := asciiStyle.Render(char)

	combinedOutput := fmt.Sprintf("%s\n%s\n%s\n%s", header, asciiArt, separator, output)
	return style.Render(combinedOutput), nil
}

// fetchSystemInfo retrieves system information
func fetchSystemInfo() (*system.SysInfo, error) {
	return system.GetSysInfo()
}

// fetchASCIIArt retrieves random ASCII art
func fetchASCIIArt() (string, error) {
	return ascii.GetRandomAsciiArts(), nil
}

// fetchBatteryStatus retrieves the current battery status
func fetchBatteryStatus() (string, error) {
	return battery.GetBatteryStatus()
}

// formatOutput formats labels and values into a structured output
func formatOutput(labels, values []string) string {
	var builder strings.Builder
	for i, label := range labels {
		paddedLabel := labelStyle.Render(label)
		paddedValue := valueStyle.Render(values[i])
		builder.WriteString(fmt.Sprintf("%s %s\n", paddedLabel, paddedValue))
	}
	return builder.String()
}
