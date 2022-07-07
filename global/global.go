package global

import (
	"winbase_exporter/stat"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Exporter Namespace.
	Namespace = "winbase"
)

func NewDesc(subsystem, name, help string) *prometheus.Desc {
	return prometheus.NewDesc(
		prometheus.BuildFQName(Namespace, subsystem, name),
		help, nil, nil,
	)
}

var (
	CPUStat     *stat.CPUStats
	ProcessStat *stat.ProcessStat
	HostStat    *stat.HostStat
	NetStat     *stat.NetStat
	DiskStat    *stat.DiskStat
	MemoryStat  *stat.MemoryStat
	GPUStat     *stat.GPUStat
)
