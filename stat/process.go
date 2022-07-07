package stat

import (
	"time"

	"github.com/shirou/gopsutil/process"
)

type ProcessStat struct {
	ProcessNum    int
	ProcessNumMax int
	ProcessNumMin int
	ProcessNumAvg float64
	ThreadNum     int
	ThreadNumMax  int
	ThreadNumMin  int
	ThreadNumAvg  float64
	TotalHandle   int
}

// type ProcessInfo struct {
// 	Pid           int
// 	Name          string
// 	Cmdline       string
// 	CreateTime    int64
// 	CpuTime       float64
// 	Memory        int64
// 	MemoryPercent float64
// 	Status        string
// 	Nice          int
// 	Uid           int
// 	Gid           int
// 	Username      string
// 	Groups        []string
// 	Threads       int
// 	OpenFiles     []string
// }

func NewProcessStat(timing int) *ProcessStat {
	processStat := ProcessStat{
		ProcessNum:    0,
		ProcessNumMax: 0,
		ProcessNumMin: 0,
		ThreadNum:     0,
		ThreadNumMax:  0,
		ThreadNumMin:  0,
		TotalHandle:   0,
	}
	go processStat.TimingStat(timing)
	return &processStat
}

func (s *ProcessStat) TimingStat(timing int) {
	for {
		processes, _ := process.Processes()
		s.ProcessNum = len(processes)
		if s.ProcessNumMax < s.ProcessNum {
			s.ProcessNumMax = s.ProcessNum
		}
		if s.ProcessNumMin > s.ProcessNum {
			s.ProcessNumMin = s.ProcessNum
		}
		s.ThreadNum = StartThreadNum(processes)
		if s.ThreadNumMax < s.ThreadNum {
			s.ThreadNumMax = s.ThreadNum
		}
		if s.ThreadNumMin > s.ThreadNum {
			s.ThreadNumMin = s.ThreadNum
		}
		time.Sleep(time.Second * time.Duration(timing))
	}
}

func StartThreadNum(processes []*process.Process) int {
	threadNum := 0
	for _, process := range processes {
		num, _ := process.NumThreads()
		threadNum += int(num)
	}
	return threadNum
}

// func (s *ProcessStat) FindProcessByName(processes []*process.Process, name string) {
// 	for _, process := range processes {
// 		processName, _ := process.Name()
// 		if processName == name {
// 			s.UnityProcess = process
// 			return
// 		}
// 	}
// 	return
// }
