package scrape

import (
	"context"
	"winbase_exporter/global"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Subsystem.
	netScrapeSubs = "netScrape"
)

var (
	netScrapeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(global.Namespace, netScrapeSubs, "net"),
		"netScrape .",
		[]string{"ip", "mac"}, nil,
	)
)

type netScrape struct {
}

func (netScrape) Name() string {
	return netScrapeSubs
}

// Help describes the role of the Scraper.
func (netScrape) Help() string {
	return "my scraper one"
}

func (netScrape) Scrape(ctx context.Context, dc string, ch chan<- prometheus.Metric) error {

	//get some from datacentor
	//dc.dosomthing...
	//may be return error，will stop this scrape's register
	//return error
	ch <- prometheus.MustNewConstMetric(
		netScrapeDesc, prometheus.GaugeValue, 1, global.NetStat.IP, global.NetStat.MAC,
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(netScrapeSubs, "TotalBytesSent", "TotalBytesSent"),
		prometheus.GaugeValue,
		float64(global.NetStat.TotalBytesSent),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(netScrapeSubs, "TotalBytesSentAvg", "TotalBytesSentAvg"),
		prometheus.GaugeValue,
		float64(global.NetStat.TotalBytesSentAvg),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(netScrapeSubs, "TotalBytesRecv", "TotalBytesRecv"),
		prometheus.GaugeValue,
		float64(global.NetStat.TotalBytesRecv),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(netScrapeSubs, "TotalBytesRecvAvg", "TotalBytesRecvAvg"),
		prometheus.GaugeValue,
		float64(global.NetStat.TotalBytesRecvAvg),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(netScrapeSubs, "InBytesRecv", "InBytesRecv"),
		prometheus.GaugeValue,
		float64(global.NetStat.InBytesRecv),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(netScrapeSubs, "InBytesRecvAvg", "InBytesRecvAvg"),
		prometheus.GaugeValue,
		float64(global.NetStat.InBytesRecvAvg),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(netScrapeSubs, "OutBytesRecv", "OutBytesRecv"),
		prometheus.GaugeValue,
		float64(global.NetStat.OutBytesRecv),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(netScrapeSubs, "InBytesRecvDiff", "InBytesRecvDiff"),
		prometheus.GaugeValue,
		float64(global.NetStat.InBytesRecvDiff),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(netScrapeSubs, "OutBytesRecv", "OutBytesRecv"),
		prometheus.GaugeValue,
		float64(global.NetStat.OutBytesRecv),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(netScrapeSubs, "OutBytesRecvAvg", "OutBytesRecvAvg"),
		prometheus.GaugeValue,
		float64(global.NetStat.OutBytesRecvAvg),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(netScrapeSubs, "OutBytesRecvDiff", "OutBytesRecvDiff"),
		prometheus.GaugeValue,
		float64(global.NetStat.OutBytesRecvDiff),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(netScrapeSubs, "InBytesSent", "InBytesSent"),
		prometheus.GaugeValue,
		float64(global.NetStat.InBytesSent),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(netScrapeSubs, "InBytesSentAvg", "InBytesSentAvg"),
		prometheus.GaugeValue,
		float64(global.NetStat.InBytesSentAvg),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(netScrapeSubs, "InBytesSentDiff", "InBytesSentDiff"),
		prometheus.GaugeValue,
		float64(global.NetStat.InBytesSentDiff),
	)
	return nil
}
