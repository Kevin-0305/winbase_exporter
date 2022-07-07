package stat

import (
	"time"
	"winbase_exporter/utils"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

type CPUStats struct {
	Load0            float64 // CPU load
	Load5            float64 // CPU load
	Load15           float64 // CPU load
	Load30Min        float64
	Load1Hour        float64
	Load1Day         float64
	Load5Stack       *utils.FloatStack
	Load15Stack      *utils.FloatStack
	Load30MinStack   *utils.FloatStack
	Load1HourStack   *utils.FloatStack
	Load1DayStack    *utils.FloatStack
	VendorId         string // CPU vendor
	CPUName          string // CPU name
	CPUCores         int    // CPU cores
	CPUThreads       int    // CPU threads
	CPUSockets       int
	CPUModel         string
	CPUMHz           float64
	CPUCacheL1       float64
	CPUCacheL2       float64
	CPUCacheL3       float64
	timing           int // 1min, 5min, 15min, 30min, 1hour, 1day
	BootTime         int // boot time
	BootDuration     int // boot duration
	CPUTimeUser      float64
	CPUTimeSystem    float64
	CPUTimeIdle      float64
	CPUTimeNice      float64
	CPUTimeIOWait    float64
	CPUTimeSoftIRQ   float64
	CPUTimeSteal     float64
	CPUTimeGuest     float64
	CPUTimeGuestNice float64
}

func NewCPUStat(timing int) *CPUStats {
	cpuInfo, _ := cpu.Info()
	timesTamp, _ := host.BootTime()
	cpuStat := CPUStats{
		Load5Stack:     utils.NewFloatStack(5),
		Load15Stack:    utils.NewFloatStack(15),
		Load30MinStack: utils.NewFloatStack(30 * 60),
		Load1HourStack: utils.NewFloatStack(60 * 60),
		Load1DayStack:  utils.NewFloatStack(24 * 60 * 60),
		VendorId:       cpuInfo[0].VendorID,
		CPUName:        cpuInfo[0].ModelName,
		CPUCores:       int(cpuInfo[0].Cores),
		CPUThreads:     int(cpuInfo[0].Cores),
		CPUModel:       cpuInfo[0].ModelName,
		CPUMHz:         cpuInfo[0].Mhz,
		BootTime:       int(timesTamp),
	}
	go cpuStat.TimingStat()
	return &cpuStat

}

func GetCPUInfo() CPUStats {
	cpuInfo, _ := cpu.Info()
	timesTamp, _ := host.BootTime()
	cpuStat := CPUStats{
		Load5Stack:     utils.NewFloatStack(5),
		Load15Stack:    utils.NewFloatStack(15),
		Load30MinStack: utils.NewFloatStack(30 * 60),
		Load1HourStack: utils.NewFloatStack(60 * 60),
		Load1DayStack:  utils.NewFloatStack(24 * 60 * 60),
		VendorId:       cpuInfo[0].VendorID,
		CPUName:        cpuInfo[0].ModelName,
		CPUCores:       int(cpuInfo[0].Cores),
		CPUThreads:     int(cpuInfo[0].Cores),
		CPUModel:       cpuInfo[0].ModelName,
		CPUMHz:         cpuInfo[0].Mhz,
		BootTime:       int(timesTamp),
	}
	return cpuStat
}

//timing: 1min, 5min, 15min, 30min, 1hour, 1day
func (s *CPUStats) TimingStat() {
	for {
		timeNowSec := int(time.Now().Unix())
		timeStat, _ := cpu.Times(false)
		load, _ := cpu.Percent(time.Second, false)
		s.Load5Stack.Push(load[0])
		s.Load15Stack.Push(load[0])
		s.Load30MinStack.Push(load[0])
		s.Load1HourStack.Push(load[0])
		s.Load1DayStack.Push(load[0])
		s.Load0 = load[0]
		s.Load5 = s.Load5Stack.Avg()
		s.Load15 = s.Load15Stack.Avg()
		s.Load30Min = s.Load30MinStack.Avg()
		s.Load1Hour = s.Load1HourStack.Avg()
		s.Load1Day = s.Load1DayStack.Avg()
		s.BootDuration = timeNowSec - s.BootTime
		s.CPUTimeUser = timeStat[0].User
		s.CPUTimeSystem = timeStat[0].System
		s.CPUTimeIdle = timeStat[0].Idle
		s.CPUTimeNice = timeStat[0].Nice
		s.CPUTimeIOWait = timeStat[0].Iowait
		s.CPUTimeSoftIRQ = timeStat[0].Softirq
		s.CPUTimeSteal = timeStat[0].Steal
		s.CPUTimeGuest = timeStat[0].Guest
		s.CPUTimeGuestNice = timeStat[0].GuestNice
		time.Sleep(time.Duration(s.timing) * time.Second)
	}
}
