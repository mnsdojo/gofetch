package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/mnsdojo/gofetch/internal/ascii"
	"github.com/mnsdojo/gofetch/internal/system"
)



var labelColor = lipgloss.Color("145") // You can choose any ANSI color code for the labels
var style = lipgloss.NewStyle().
	Bold(true).
	Padding(5).
	Border(lipgloss.ThickBorder(), true, true, true, true).Align(lipgloss.Center).
	BorderForeground(lipgloss.Color("228"))

var asciiStyle = lipgloss.NewStyle().
	Align(lipgloss.Center)

var separatorStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), false, false, true, false).
	BorderForeground(lipgloss.Color("201")).
	MarginTop(1).
	MarginBottom(1)

var headerStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("228")).Align(lipgloss.Center)

var labelStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(labelColor).
	Width(15)

var valueStyle = lipgloss.NewStyle().Bold(true).
	Width(30)

func main() {
	info := system.GetSysInfo()
	char := ascii.GetRandomAsciiArts()

	// Center align the header text
	header := headerStyle.Render("gofetch")

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
	asciiArt := asciiStyle.Render(char)

	combinedOutput := fmt.Sprintf("%s\n%s\n%s\n%s", header, asciiArt, separator, output)
	fmt.Println(style.Render(combinedOutput))
}
