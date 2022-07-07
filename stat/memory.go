package stat

import (
	"time"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type MemoryStat struct {
	BootTime          int
	TotalMemory       int
	UsedMemory        int
	UsedMemoryMax     int
	UsedMemoryAvg     float64
	UsePercent        float64
	FreeMemory        int
	FreeMemoryMax     int
	FreeMemoryAvg     float64
	ActiveMemory      int
	ActiveMemoryMax   int
	ActiveMemoryAvg   float64
	InactiveMemory    int
	InactiveMemoryMax int
	InactiveMemoryAvg float64
	WiredMemory       int
	WiredMemoryMax    int
	WiredMemoryAvg    float64
	BuffersMemory     int
	BuffersMemoryMax  int
	BuffersMemoryAvg  float64
	CachedMemory      int
	CachedMemoryMax   int
	CachedMemoryAvg   float64
	TotalSwap         int
	TotalSwapMax      int
	TotalSwapAvg      float64
	UsedSwap          int
	UsedSwapMax       int
	UsedSwapAvg       float64
}

func NewMemoryStat(timing int) *MemoryStat {
	bootTime, _ := host.BootTime()
	memoryStat := MemoryStat{
		BootTime:          int(bootTime),
		TotalMemory:       0,
		UsedMemory:        0,
		UsedMemoryMax:     0,
		UsedMemoryAvg:     0,
		FreeMemory:        0,
		FreeMemoryMax:     0,
		FreeMemoryAvg:     0,
		ActiveMemory:      0,
		ActiveMemoryMax:   0,
		ActiveMemoryAvg:   0,
		InactiveMemory:    0,
		InactiveMemoryMax: 0,
		InactiveMemoryAvg: 0,
		WiredMemory:       0,
		WiredMemoryMax:    0,
		WiredMemoryAvg:    0,
		BuffersMemory:     0,
		BuffersMemoryMax:  0,
		BuffersMemoryAvg:  0,
		CachedMemory:      0,
		CachedMemoryMax:   0,
		CachedMemoryAvg:   0,
		TotalSwap:         0,
		TotalSwapMax:      0,
		TotalSwapAvg:      0,
		UsedSwap:          0,
		UsedSwapMax:       0,
		UsedSwapAvg:       0,
	}
	go memoryStat.TimingStat(timing)
	return &memoryStat
}

// func GetMemoryInfo() *MemoryStat {
// 	memory, _ := mem.VirtualMemory()
// 	return memoryStat
// }

func (s *MemoryStat) TimingStat(timing int) {
	for {
		nowTime := int(time.Now().Unix())
		memory, _ := mem.VirtualMemory()
		s.TotalMemory = int(memory.Total)
		s.UsedMemory = int(memory.Used)
		if s.UsedMemoryMax < s.UsedMemory {
			s.UsedMemoryMax = s.UsedMemory
		}
		s.UsePercent = memory.UsedPercent
		s.UsedMemoryAvg = (s.UsedMemoryAvg*float64(nowTime-s.BootTime-1) + float64(s.UsedMemoryAvg)) / float64(nowTime-s.BootTime)
		s.FreeMemory = int(memory.Free)
		if s.FreeMemoryMax < s.FreeMemory {
			s.FreeMemoryMax = s.FreeMemory
		}
		s.FreeMemoryAvg = (s.FreeMemoryAvg*float64(nowTime-s.BootTime-1) + float64(s.FreeMemory)) / float64(nowTime-s.BootTime)
		s.ActiveMemory = int(memory.Active)
		if s.ActiveMemoryMax < s.ActiveMemory {
			s.ActiveMemoryMax = s.ActiveMemory
		}
		s.ActiveMemoryAvg = (s.ActiveMemoryAvg*float64(nowTime-s.BootTime-1) + float64(s.ActiveMemory)) / float64(nowTime-s.BootTime)
		s.InactiveMemory = int(memory.Inactive)
		if s.InactiveMemoryMax < s.InactiveMemory {
			s.InactiveMemoryMax = s.InactiveMemory
		}
		s.InactiveMemoryAvg = (s.InactiveMemoryAvg*float64(nowTime-s.BootTime-1) + float64(s.InactiveMemory)) / float64(nowTime-s.BootTime)
		s.WiredMemory = int(memory.Wired)
		if s.WiredMemoryMax < s.WiredMemory {
			s.WiredMemoryMax = s.WiredMemory
		}
		s.WiredMemoryAvg = (s.WiredMemoryAvg*float64(nowTime-s.BootTime-1) + float64(s.WiredMemory)) / float64(nowTime-s.BootTime)
		s.BuffersMemory = int(memory.Buffers)
		if s.BuffersMemoryMax < s.BuffersMemory {
			s.BuffersMemoryMax = s.BuffersMemory
		}
		s.BuffersMemoryAvg = (s.BuffersMemoryAvg*float64(nowTime-s.BootTime-1) + float64(s.BuffersMemory)) / float64(nowTime-s.BootTime)
		s.CachedMemory = int(memory.Cached)
		if s.CachedMemoryMax < s.CachedMemory {
			s.CachedMemoryMax = s.CachedMemory
		}
		s.CachedMemoryAvg = (s.CachedMemoryAvg*float64(nowTime-s.BootTime-1) + float64(s.CachedMemory)) / float64(nowTime-s.BootTime)
		swap, _ := mem.SwapMemory()
		s.TotalSwap = int(swap.Total)
		if s.TotalSwapMax < s.TotalSwap {
			s.TotalSwapMax = s.TotalSwap
		}
		s.TotalSwapAvg = (s.TotalSwapAvg*float64(nowTime-s.BootTime-1) + float64(s.TotalSwap)) / float64(nowTime-s.BootTime)
		s.UsedSwap = int(swap.Used)
		if s.UsedSwapMax < s.UsedSwap {
			s.UsedSwapMax = s.UsedSwap
		}
		s.UsedSwapAvg = (s.UsedSwapAvg*float64(nowTime-s.BootTime-1) + float64(s.UsedSwap)) / float64(nowTime-s.BootTime)
		time.Sleep(time.Duration(timing) * time.Second)
	}
}
