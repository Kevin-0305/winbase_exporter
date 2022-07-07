package stat

import (
	"github.com/shirou/gopsutil/host"
)

type HostStat struct {
	BootTime    int
	CPUName     string
	CPUCores    int
	GPUName     string
	TotalMemory int
	TotalDisk   int
	MacAddress  string
	IPAddress   string
	OSName      string
}

func NewHostStat() *HostStat {
	bootTime, _ := host.BootTime()
	// cpuInfo, _ := cpu.Info()
	// cpuName := cpuInfo[0].ModelName
	//gpuInfo, _ := gpu.Info()

	hostStat := HostStat{
		BootTime:    int(bootTime),
		CPUName:     "",
		GPUName:     "",
		TotalMemory: 0,
		TotalDisk:   0,
		MacAddress:  "",
		IPAddress:   "",
		OSName:      "",
	}
	return &hostStat
}
