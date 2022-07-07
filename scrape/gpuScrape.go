package scrape

import (
	"context"
	"winbase_exporter/global"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Subsystem.
	gpuScrapeSubs = "gpuScrape"
)

var (
	gpuScrapeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(global.Namespace, gpuScrapeSubs, "gpu"),
		"gpuScrape .",
		[]string{"gpuName", "version"}, nil,
	)
)

type GpuScrape struct {
}

func (GpuScrape) Name() string {
	return gpuScrapeSubs
}

// Help describes the role of the Scraper.
func (GpuScrape) Help() string {
	return "my scraper one"
}

func (GpuScrape) Scrape(ctx context.Context, dc string, ch chan<- prometheus.Metric) error {

	//get some from datacentor
	//dc.dosomthing...
	//may be return error，will stop this scrape's register
	//return error
	ch <- prometheus.MustNewConstMetric(
		gpuScrapeDesc, prometheus.GaugeValue, 1, global.GPUStat.GPUName, global.GPUStat.NvidiaDriver,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(gpuScrapeSubs, "GPUFanSpeed", "GPUFanSpeed"),
		prometheus.GaugeValue,
		float64(global.GPUStat.GPUFanSpeed),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(gpuScrapeSubs, "GPUUsePercent", "GPUUsePercent"),
		prometheus.GaugeValue,
		float64(global.GPUStat.GPUUsePercent),
	)

	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(gpuScrapeSubs, "GPUMemoryTotal", "GPUMemoryTotal"),
		prometheus.GaugeValue,
		float64(global.GPUStat.GPUMemoryTotal),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(gpuScrapeSubs, "GPUMemoryUsed", "GPUMemoryUsed"),
		prometheus.GaugeValue,
		float64(global.GPUStat.GPUMemoryUsed),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(gpuScrapeSubs, "GPUMemoryUsedMax", "GPUMemoryUsedMax"),
		prometheus.GaugeValue,
		float64(global.GPUStat.GPUMemoryUsedMax),
	)

	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(gpuScrapeSubs, "GPUPowerUsage", "GPUPowerUsage"),
		prometheus.GaugeValue,
		float64(global.GPUStat.GPUPowerUsage),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(gpuScrapeSubs, "GPUPowerUsageMax", "GPUPowerUsageMax"),
		prometheus.GaugeValue,
		float64(global.GPUStat.GPUPowerUsageMax),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(gpuScrapeSubs, "GPUPowerUsageAvg", "GPUPowerUsageAvg"),
		prometheus.GaugeValue,
		float64(global.GPUStat.GPUPowerUsageAvg),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(gpuScrapeSubs, "GPUTemperature", "GPUTemperature"),
		prometheus.GaugeValue,
		float64(global.GPUStat.GPUTemperature),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(gpuScrapeSubs, "GPUTemperatureMax", "GPUTemperatureMax"),
		prometheus.GaugeValue,
		float64(global.GPUStat.GPUTemperatureMax),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(gpuScrapeSubs, "GPUTemperatureAvg", "GPUTemperatureAvg"),
		prometheus.GaugeValue,
		float64(global.GPUStat.GPUTemperatureAvg),
	)

	return nil
}
