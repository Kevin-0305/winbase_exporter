package scrape

import (
	"context"
	"winbase_exporter/global"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Subsystem.
	memoryScrapeSubs = "memoryScrape"
)

var (
	memoryScrapeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(global.Namespace, memoryScrapeSubs, "memory"),
		"memoryScrape  .",
		[]string{"memory"}, nil,
	)
)

type MemoryScrape struct {
}

func (MemoryScrape) Name() string {
	return memoryScrapeSubs
}

// Help describes the role of the Scraper.
func (MemoryScrape) Help() string {
	return "my scraper one"
}

func (MemoryScrape) Scrape(ctx context.Context, dc string, ch chan<- prometheus.Metric) error {
	//get some from datacentor
	//dc.dosomthing...
	//may be return error，will stop this scrape's register
	//return error
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "cpuMHZ", "cpuMHZ"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.TotalMemory/1024/1024),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "UsedMemory", "UsedMemory"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.UsedMemory/1024/1024),
	)

	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "UsedMemoryMax", "UsedMemoryMax"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.UsedMemoryMax/1024/1024),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "UsedMemoryAvg", "UsedMemoryAvg"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.UsedMemoryAvg/1024/1024),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "UsePercent", "UsePercent"),
		prometheus.GaugeValue,
		global.MemoryStat.UsePercent,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "FreeMemory", "FreeMemory"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.FreeMemory/1024/1024),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "FreeMemoryMax", "FreeMemoryMax"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.FreeMemoryMax/1024/1024),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "FreeMemoryAvg", "FreeMemoryAvg"),
		prometheus.GaugeValue,
		global.MemoryStat.FreeMemoryAvg/1024/1024,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "ActiveMemory", "ActiveMemory"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.ActiveMemory/1024/1024),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "ActiveMemoryMax", "ActiveMemoryMax"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.ActiveMemoryMax/1024/1024),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "ActiveMemoryAvg", "ActiveMemoryAvg"),
		prometheus.GaugeValue,
		global.MemoryStat.ActiveMemoryAvg/1024/1024,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "InactiveMemory", "InactiveMemory"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.InactiveMemory/1024/1024),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "InactiveMemoryMax", "InactiveMemoryMax"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.InactiveMemoryMax/1024/1024),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "InactiveMemoryAvg", "InactiveMemoryAvg"),
		prometheus.GaugeValue,
		global.MemoryStat.InactiveMemoryAvg/1024/1024,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "WiredMemory", "WiredMemory"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.WiredMemory/1024/1024),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "WiredMemoryMax", "WiredMemoryMax"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.WiredMemoryMax/1024/1024),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "WiredMemoryAvg", "WiredMemoryAvg"),
		prometheus.GaugeValue,
		global.MemoryStat.WiredMemoryAvg/1024/1024,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "BuffersMemory", "BuffersMemory"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.BuffersMemory/1024/1024),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "BuffersMemoryMax", "BuffersMemoryMax"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.BuffersMemoryMax/1024/1024),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "CachedMemory", "CachedMemory"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.CachedMemory/1024/1024),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "CachedMemoryMax", "CachedMemoryMax"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.CachedMemoryMax/1024/1024),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "CachedMemoryAvg", "CachedMemoryAvg"),
		prometheus.GaugeValue,
		global.MemoryStat.CachedMemoryAvg/1024/1024,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "TotalSwap", "TotalSwap"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.TotalSwap/1024/1024),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "TotalSwapMax", "TotalSwapMax"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.TotalSwapMax/1024/1024),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "TotalSwapAvg", "TotalSwapAvg"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.TotalSwapAvg/1024/1024),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(memoryScrapeSubs, "UsedSwap", "UsedSwap"),
		prometheus.GaugeValue,
		float64(global.MemoryStat.UsedSwap/1024/1024),
	)

	return nil
}
