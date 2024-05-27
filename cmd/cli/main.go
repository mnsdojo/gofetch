package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/mnsdojo/gofetch/internal/ascii"
	"github.com/mnsdojo/gofetch/internal/system"
)

var style = lipgloss.NewStyle().
	Bold(true).
	Padding(5).
	Border(lipgloss.ThickBorder(), true, true, true, true).
	BorderForeground(lipgloss.Color("228"))

var separatorStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), false, false, true, false).
	BorderForeground(lipgloss.Color("201")).
	MarginTop(1).
	MarginBottom(1)

var headerStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("228"))

var labelStyle = lipgloss.NewStyle().
	Bold(true).
	Width(15)

var valueStyle = lipgloss.NewStyle().Bold(true).
	Width(30)

func main() {
	info := system.GetSysInfo()
	char := ascii.GetRandomAsciiArts()

	header := headerStyle.Render("treefetch")

	// Define the labels and their corresponding values
	labels := []string{"Username:", "OS Release:", "Kernel:", "Default Shell:", "Uptime:", "Memory Usage:"}
	values := []string{info.UserName, info.OSRelease, info.Kernel, info.DefaultShell, info.Uptime, info.MemoryUsage}

	// Create the output string with aligned labels and values using Lipgloss
	var output string
	for i, label := range labels {
		paddedLabel := labelStyle.Render(label)
		paddedValue := valueStyle.Render(values[i])
		output += fmt.Sprintf("%s %s\n", paddedLabel, paddedValue)
	}

	separator := separatorStyle.Render("──────────────────────────────────────────────────")

	combinedOutput := fmt.Sprintf("%s\n%s\n%s\n%s", header, char, separator, output)
	fmt.Println(style.Render(combinedOutput))
}
