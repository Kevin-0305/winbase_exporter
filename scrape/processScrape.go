package scrape

import (
	"context"
	"winbase_exporter/global"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Subsystem.
	processScrapeSubs = "processScrape"
)

var (
	processScrapeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(global.Namespace, processScrapeSubs, "process"),
		"processScrape .",
		[]string{"ip", "mac"}, nil,
	)
)

type processScrape struct {
}

func (processScrape) Name() string {
	return processScrapeSubs
}

// Help describes the role of the Scraper.
func (processScrape) Help() string {
	return "my scraper one"
}

func (processScrape) Scrape(ctx context.Context, dc string, ch chan<- prometheus.Metric) error {

	//get some from datacentor
	//dc.dosomthing...
	//may be return error，will stop this scrape's register
	//return error
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(processScrapeSubs, "ProcessNum", "ProcessNum"),
		prometheus.GaugeValue,
		float64(global.ProcessStat.ProcessNum),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(processScrapeSubs, "ProcessNumMax", "ProcessNumMax"),
		prometheus.GaugeValue,
		float64(global.ProcessStat.ProcessNumMax),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(processScrapeSubs, "ProcessNumMin", "ProcessNumMin"),
		prometheus.GaugeValue,
		float64(global.ProcessStat.ProcessNumMin),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(processScrapeSubs, "ProcessNumAvg", "ProcessNumAvg"),
		prometheus.GaugeValue,
		float64(global.ProcessStat.ProcessNumAvg),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(processScrapeSubs, "ThreadNum", "ThreadNum"),
		prometheus.GaugeValue,
		float64(global.ProcessStat.ThreadNum),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(processScrapeSubs, "ThreadNumMax", "ThreadNumMax"),
		prometheus.GaugeValue,
		float64(global.ProcessStat.ThreadNumMax),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(processScrapeSubs, "ThreadNumMin", "ThreadNumMin"),
		prometheus.GaugeValue,
		float64(global.ProcessStat.ThreadNumMin),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(processScrapeSubs, "ThreadNumAvg", "ThreadNumAvg"),
		prometheus.GaugeValue,
		float64(global.ProcessStat.ThreadNumAvg),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(processScrapeSubs, "TotalHandle", "TotalHandle"),
		prometheus.GaugeValue,
		float64(global.ProcessStat.TotalHandle),
	)

	return nil
}
