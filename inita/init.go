package inita

import (
	"winbase_exporter/global"
	"winbase_exporter/stat"

	"github.com/shirou/gopsutil/host"
)

func InitStat() {
	// Initialize the global variables.
	hostInfo, _ := host.Info()
	bootTime, _ := host.BootTime()
	global.CPUStat = stat.NewCPUStat(1)
	global.GPUStat = stat.NewGPUStat()
	global.ProcessStat = stat.NewProcessStat(1)
	global.MemoryStat = stat.NewMemoryStat(1)
	global.NetStat = stat.NewNetStat(10)
	global.DiskStat = stat.NewDiskStat(10)
	h := &stat.HostStat{
		BootTime:    int(bootTime),
		CPUName:     global.CPUStat.CPUName,
		CPUCores:    global.CPUStat.CPUCores,
		GPUName:     global.GPUStat.GPUName,
		TotalMemory: global.MemoryStat.TotalMemory,
		TotalDisk:   global.DiskStat.TotalDisk,
		MacAddress:  global.NetStat.MAC,
		IPAddress:   global.NetStat.IP,
		OSName:      hostInfo.OS,
	}
	global.HostStat = h

}
