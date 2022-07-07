package scrape

import (
	"context"
	"winbase_exporter/global"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Subsystem.
	cpuScrapeSubs = "cpuScrape"
)

var (
	cpuScrapeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(global.Namespace, cpuScrapeSubs, "cpu"),
		"cpuScrape  .",
		[]string{"gpu"}, nil,
	)
)

type CpuScrape struct {
}

func (CpuScrape) Name() string {
	return cpuScrapeSubs
}

// Help describes the role of the Scraper.
func (CpuScrape) Help() string {
	return "my scraper one"
}

func (CpuScrape) Scrape(ctx context.Context, dc string, ch chan<- prometheus.Metric) error {

	//get some from datacentor
	//dc.dosomthing...
	//may be return error，will stop this scrape's register
	//return error
	ch <- prometheus.MustNewConstMetric(

		// cpuScrapeDesc, prometheus.CounterValue, global.CPUStat.CPUMHz, "test1",
		cpuScrapeDesc, prometheus.GaugeValue, 1, global.CPUStat.CPUName,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "cpuMHZ", "cpuMHZ"),
		prometheus.GaugeValue,
		global.CPUStat.CPUMHz,
	)
	// ch <- prometheus.MustNewConstMetric(
	// 	global.NewDesc(cpuScrapeSubs, "cpuName", "cpuName"),
	// 	prometheus.GaugeValue,
	// 	global.CPUStat.CpuName,
	// )
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "Load0", "Load0"),
		prometheus.GaugeValue,
		global.CPUStat.Load0,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "Load5", "Load5"),
		prometheus.GaugeValue,
		global.CPUStat.Load5,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "Load15", "Load15"),
		prometheus.GaugeValue,
		global.CPUStat.Load15,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "Load30Min", "Load30Min"),
		prometheus.GaugeValue,
		global.CPUStat.Load30Min,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "Load1Hour", "Load1Hour"),
		prometheus.GaugeValue,
		global.CPUStat.Load1Hour,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "Load1Day", "Load1Day"),
		prometheus.GaugeValue,
		global.CPUStat.Load1Day,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "CpuCores", "CpuCores"),
		prometheus.GaugeValue,
		float64(global.CPUStat.CPUCores),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "CPUThreads", "CPUThreads"),
		prometheus.GaugeValue,
		float64(global.CPUStat.CPUThreads),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "CPUCacheL1", "CPUCacheL1"),
		prometheus.GaugeValue,
		float64(global.CPUStat.CPUCacheL1),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "CPUCacheL2", "CPUCacheL2"),
		prometheus.GaugeValue,
		float64(global.CPUStat.CPUCacheL2),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "CPUCacheL3", "CPUCacheL3"),
		prometheus.GaugeValue,
		float64(global.CPUStat.CPUCacheL3),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "BootDuration", "BootDuration"),
		prometheus.GaugeValue,
		float64(global.CPUStat.BootDuration),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "CPUTimeUser", "CPUTimeUser"),
		prometheus.GaugeValue,
		global.CPUStat.CPUTimeUser,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "CPUTimeSystem", "CPUTimeSystem"),
		prometheus.GaugeValue,
		global.CPUStat.CPUTimeSystem,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "CPUTimeIdle", "CPUTimeIdle"),
		prometheus.GaugeValue,
		global.CPUStat.CPUTimeIdle,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "CPUTimeIOWait", "CPUTimeIOWait"),
		prometheus.GaugeValue,
		global.CPUStat.CPUTimeIOWait,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "CPUTimeSoftIRQ", "CPUTimeSoftIRQ"),
		prometheus.GaugeValue,
		global.CPUStat.CPUTimeSoftIRQ,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "CPUTimeSteal", "CPUTimeSteal"),
		prometheus.GaugeValue,
		global.CPUStat.CPUTimeSteal,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "CPUTimeGuest", "CPUTimeGuest"),
		prometheus.GaugeValue,
		global.CPUStat.CPUTimeGuest,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(cpuScrapeSubs, "CPUTimeGuestNice", "CPUTimeGuestNice"),
		prometheus.GaugeValue,
		global.CPUStat.CPUTimeGuestNice,
	)

	return nil
}
