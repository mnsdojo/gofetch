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
	shell, err := getCmdOutput("bash", "-c", "echo $SHELL")
	if err != nil {
		return "", err
	}
	shellPath := strings.TrimSpace(string(shell))
	parts := strings.Split(shellPath, "/")
	return parts[len(parts)-1], nil
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

// getOsRelease retrieves the OS release name from /etc/os-release
func getOsRelease() (string, error) {
	// Use awk to extract the value after the '=' sign, and remove surrounding quotes and spaces
	
	// copied from the internet !
	cmd := "awk -F= '/^PRETTY_NAME=/ {gsub(/^\"|\"$/, \"\", $2); print $2}' /etc/os-release"
	return getCmdOutput("bash", "-c", cmd)
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

	return fmt.Sprintf("%d MB/%d MB", (availableMb)/1024, (totalMb)/1024), nil
}

func GetSysInfo() (*SysInfo, error) {
	systemInfo := &SysInfo{}

	userName, err := getUserInfo()
	if err != nil {
		return systemInfo, fmt.Errorf("error retrieving user name: %v", err)
	}

	osRelease, err := getOsRelease()
	if err != nil {
		return systemInfo, fmt.Errorf("error retrieving OS release: %v", err)
	}

	kernel, err := getKernel()
	if err != nil {
		return systemInfo, fmt.Errorf("error retrieving kernel version: %v", err)
	}

	defaultShell, err := getDefaultShell()
	if err != nil {
		return systemInfo, fmt.Errorf("error retrieving default shell: %v", err)
	}

	memoryUsage, err := getMemoryUsage()
	if err != nil {
		return systemInfo, fmt.Errorf("error retrieving memory usage: %v", err)
	}

	uptime, err := getUptime()
	if err != nil {
		return systemInfo, fmt.Errorf("error retrieving uptime: %v", err)
	}

	systemInfo = &SysInfo{
		UserName:     userName,
		OSRelease:    osRelease,
		Kernel:       kernel,
		DefaultShell: defaultShell,
		Uptime:       uptime,
		MemoryUsage:  memoryUsage,
	}

	return systemInfo, nil
}
