package system

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type SysInfo struct {
	UserName     string
	OSRelease    string
	Kernel       string
	DefaultShell string
	Uptime       string
	MemoryUsage  string
}

// getDefaultShell retrieves the default shell
func getDefaultShell() (string, error) {
	return getCmdOutput("echo", "$SHELL")
}

// getCmdOutput executes a command and returns its trimmed output
func getCmdOutput(cmd string, args ...string) (string, error) {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// getKernel retrieves the kernel version
func getKernel() (string, error) {
	return getCmdOutput("uname", "-r")
}

// getUserInfo retrieves the username
func getUserInfo() (string, error) {
	return getCmdOutput("whoami")
}

func getOsRelease() (string, error) {
	return getCmdOutput("grep", "^PRETTY_NAME=", "/etc/os-release")
}

// getUptime retrieves the system uptime
func getUptime() (string, error) {
	return getCmdOutput("uptime", "-p")
}

func getMemoryUsage() (string, error) {
	memTotal, err := getCmdOutput("grep", "MemTotal", "/proc/meminfo")
	if err != nil {
		return "", err
	}
	memAvailable, err := getCmdOutput("grep", "MemAvailable", "/proc/meminfo")
	if err != nil {
		return "", err
	}

	totalKb := strings.Fields(memTotal)[1]
	availableKb := strings.Fields(memAvailable)[1]

	totalMb, err := strconv.Atoi(totalKb)
	if err != nil {
		return "", err
	}
	availableMb, err := strconv.Atoi(availableKb)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%.2f MB/%.2f MB", float64(availableMb)/1024, float64(totalMb)/1024), nil
}

// GetSysInfo gathers the desired system information
func GetSysInfo() SysInfo {
	var systemInfo SysInfo

	userName, err := getUserInfo()
	if err != nil {
		fmt.Printf("Error retrieving user name: %v\n", err)
		return systemInfo
	}

	osRelease, err := getOsRelease()
	if err != nil {
		fmt.Printf("Error retrieving OS release: %v\n", err)
		return systemInfo
	}
	// Strip the PRETTY_NAME= prefix and quotes from osRelease
	osRelease = strings.TrimPrefix(osRelease, "PRETTY_NAME=")
	osRelease = strings.Trim(osRelease, `"`)

	kernel, err := getKernel()
	if err != nil {
		fmt.Printf("Error retrieving kernel version: %v\n", err)
		return systemInfo
	}

	defaultShell, err := getDefaultShell()
	if err != nil {
		fmt.Printf("Error retrieving default shell: %v\n", err)
		return systemInfo
	}

	memoryUsage, err := getMemoryUsage()
	if err != nil {
		fmt.Printf("Error retrieving memory usage: %v\n", err)
		return systemInfo
	}

	uptime, err := getUptime()
	if err != nil {
		fmt.Printf("Error retrieving uptime: %v\n", err)
		return systemInfo
	}

	systemInfo = SysInfo{
		UserName:     userName,
		OSRelease:    osRelease,
		Kernel:       kernel,
		DefaultShell: defaultShell,
		Uptime:       uptime,
		MemoryUsage:  memoryUsage,
	}

	return systemInfo
}

