package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	fmt.Println("MyFetch - A Neofetch Alternative in Go")
	fmt.Println("--------------------------------------")

	// Hostname
	hostname, _ := os.Hostname()
	fmt.Printf("Hostname: %s\n", hostname)

	// OS Info
	info, _ := host.Info()
	fmt.Printf("OS: %s %s\n", info.Platform, info.PlatformVersion)
	fmt.Printf("Kernel: %s\n", info.KernelVersion)

	// Uptime
	uptime := time.Duration(info.Uptime) * time.Second
	fmt.Printf("Uptime: %s\n", uptime)

	// CPU
	cpuInfo, _ := cpu.Info()
	fmt.Printf("CPU: %s (%d cores)\n", cpuInfo[0].ModelName, runtime.NumCPU())

	// Memory
	vmem, _ := mem.VirtualMemory()
	fmt.Printf("Memory: %.2f GB / %.2f GB\n", float64(vmem.Used)/1e9, float64(vmem.Total)/1e9)

	// Shell and Terminal
	shell := os.Getenv("SHELL")
	term := os.Getenv("TERM")
	fmt.Printf("Shell: %s\n", shell)
	fmt.Printf("Terminal: %s\n", term)
}
