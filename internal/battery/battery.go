package battery

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func GetBatteryStatus() (string, error) {
	var filePath string

	// Check BAT1 first
	filePath = "/sys/class/power_supply/BAT1/"
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		// Use BAT1 if it exists
		return getBatteryData(filePath)
	}

	// Check BAT0 as fallback
	filePath = "/sys/class/power_supply/BAT0/"
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		// Use BAT0 if it exists
		return getBatteryData(filePath)
	}

	// No battery information found
	return "", fmt.Errorf("battery information not found")
}

func getBatteryData(filePath string) (string, error) {
	// Read battery capacity
	capacityData, err := os.Open(filePath + "capacity")
	if err != nil {
		return "", fmt.Errorf("failed to open capacity file: %w", err)
	}
	defer capacityData.Close()

	capacityBytes, err := io.ReadAll(capacityData)
	if err != nil {
		return "", fmt.Errorf("failed to read capacity data: %w", err)
	}
	capacity, err := strconv.Atoi(strings.TrimSpace(string(capacityBytes)))
	if err != nil {
		return "", fmt.Errorf("failed to convert capacity data: %w", err)
	}

	// Read battery status (Charging/Discharging)
	statusData, err := os.Open(filePath + "status")
	if err != nil {
		return "", fmt.Errorf("failed to open status file: %w", err)
	}
	defer statusData.Close()

	statusBytes, err := io.ReadAll(statusData)
	if err != nil {
		return "", fmt.Errorf("failed to read status data: %w", err)
	}
	status := strings.TrimSpace(string(statusBytes))

	// Determine charging state
	charging := "🔌"
	if status == "Charging" || status == "Full" {
		charging = "🔋"
	}

	// Format battery status information
	return fmt.Sprintf("Battery: %d%% (%s)", capacity, charging), nil

}
