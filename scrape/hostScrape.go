package scrape

import (
	"context"
	"winbase_exporter/global"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Subsystem.
	hostScrapeSubs = "hostScrape"
)

var (
	hostScrapeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(global.Namespace, hostScrapeSubs, "host"),
		"hostInfo .",
		[]string{"cpu", "gpu", "os", "mac"}, nil,
	)
)

type HostScrape struct {
}

func (HostScrape) Name() string {
	return hostScrapeSubs
}

// Help describes the role of the Scraper.
func (HostScrape) Help() string {
	return "hostInfo scrape ."
}

func (HostScrape) Scrape(ctx context.Context, dc string, ch chan<- prometheus.Metric) error {

	//get some from datacentor
	//dc.dosomthing...
	//may be return error，will stop this scrape's register
	//return error
	ch <- prometheus.MustNewConstMetric(
		hostScrapeDesc, prometheus.GaugeValue, 1, global.HostStat.CPUName, global.HostStat.GPUName, global.HostStat.OSName, global.HostStat.MacAddress,
	)
	return nil
}
