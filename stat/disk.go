package stat

import (
	"time"

	"github.com/shirou/gopsutil/disk"
)

type DiskStat struct {
	TotalDisk  int
	UsedDisk   int
	UsePercent float64
	ReadSize   int
	WriteSize  int
	ReadDiff   int
	WriteDiff  int
	DisksInfo  map[string]*disk.UsageStat
	DiskParts  map[string]disk.PartitionStat
	DiskIOStat map[string]disk.IOCountersStat
	DiskNames  []string
}

func NewDiskStat(timing int) *DiskStat {
	totalDisk := 0
	disksNames := []string{}
	parts, _ := disk.Partitions(false)
	partsReal := make(map[string]disk.PartitionStat)
	for _, part := range parts {
		if part.Fstype == "NTFS" && []rune(part.Mountpoint)[0] < []rune("N")[0] {
			diskInfo, _ := disk.Usage(part.Mountpoint)
			totalDisk += int(diskInfo.Total)
			partsReal[part.Mountpoint] = part
			disksNames = append(disksNames, part.Mountpoint)
		}
	}

	diskStat := DiskStat{
		TotalDisk:  totalDisk,
		UsedDisk:   0,
		UsePercent: 0,
		ReadSize:   0,
		WriteSize:  0,
		ReadDiff:   0,
		WriteDiff:  0,
		DisksInfo:  make(map[string]*disk.UsageStat),
		DiskParts:  partsReal,
		DiskIOStat: make(map[string]disk.IOCountersStat),
		DiskNames:  disksNames,
	}
	go diskStat.TimingStat(timing)
	return &diskStat
}

func GetDiskInfo() DiskStat {
	totalDisk := 0
	parts, _ := disk.Partitions(false)
	partsReal := make(map[string]disk.PartitionStat)
	for _, part := range parts {
		if part.Fstype == "NTFS" && []rune(part.Mountpoint)[0] < []rune("N")[0] {
			diskInfo, _ := disk.Usage(part.Mountpoint)
			totalDisk += int(diskInfo.Total)
			partsReal[part.Mountpoint] = part
		}
	}
	diskStat := DiskStat{
		TotalDisk:  totalDisk,
		UsedDisk:   0,
		UsePercent: 0,
		ReadSize:   0,
		WriteSize:  0,
		ReadDiff:   0,
		WriteDiff:  0,
		DisksInfo:  nil,
		DiskParts:  nil,
		DiskNames:  []string{},
	}
	return diskStat

}

func (s *DiskStat) TimingStat(timing int) {
	for {
		readSize := 0
		writeSize := 0
		usedDisk := 0
		for _, name := range s.DiskNames {
			diskInfo, _ := disk.Usage(name)
			s.DisksInfo[name] = diskInfo
			usedDisk += int(diskInfo.Used)
		}
		ioStat, _ := disk.IOCounters(s.DiskNames...)
		for _, v := range ioStat {
			s.DiskIOStat[v.Name] = v
			readSize += int(v.ReadBytes)
			writeSize += int(v.WriteBytes)
		}
		s.ReadDiff = readSize - s.ReadSize
		s.WriteDiff = writeSize - s.WriteSize
		s.ReadSize = readSize
		s.WriteSize = writeSize
		s.UsedDisk = usedDisk
		s.UsePercent = float64(usedDisk) / float64(s.TotalDisk)
		time.Sleep(time.Second * time.Duration(timing))
	}
}
