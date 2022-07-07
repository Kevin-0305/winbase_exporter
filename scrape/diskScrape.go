package scrape

import (
	"context"
	"winbase_exporter/global"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Subsystem.
	diskScrapeSubs = "diskScrape"
)

var (
	diskScrapeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(global.Namespace, diskScrapeSubs, "disk"),
		"diskScrape  .",
		[]string{"path", "type"}, nil,
	)
)

type DiskScrape struct {
}

func (DiskScrape) Name() string {
	return diskScrapeSubs
}

// Help describes the role of the Scraper.
func (DiskScrape) Help() string {
	return "my scraper one"
}

func (DiskScrape) Scrape(ctx context.Context, dc string, ch chan<- prometheus.Metric) error {

	//get some from datacentor
	//dc.dosomthing...
	//may be return error，will stop this scrape's register
	//return error
	// ch <- prometheus.MustNewConstMetric(
	// 	// diskScrapeDesc, prometheus.CounterValue, global.DiskStat.diskMHz, "test1",
	// 	diskScrapeDesc, prometheus.GaugeValue, 1, global.DiskStat.diskName,
	// )
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(diskScrapeSubs, "TotalDisk", "TotalDisk"),
		prometheus.GaugeValue,
		float64(global.DiskStat.TotalDisk),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(diskScrapeSubs, "UsedDisk", "UsedDisk"),
		prometheus.GaugeValue,
		float64(global.DiskStat.UsedDisk),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(diskScrapeSubs, "UsePercent", "UsePercent"),
		prometheus.GaugeValue,
		float64(global.DiskStat.UsePercent),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(diskScrapeSubs, "ReadSize", "ReadSize"),
		prometheus.GaugeValue,
		float64(global.DiskStat.ReadSize),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(diskScrapeSubs, "WriteSize", "WriteSize"),
		prometheus.GaugeValue,
		float64(global.DiskStat.WriteSize),
	)

	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(diskScrapeSubs, "ReadDiff", "ReadDiff"),
		prometheus.GaugeValue,
		float64(global.DiskStat.ReadDiff),
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(diskScrapeSubs, "WriteDiff", "WriteDiff"),
		prometheus.GaugeValue,
		float64(global.DiskStat.WriteDiff),
	)
	for _, diskName := range global.DiskStat.DiskNames {
		ch <- prometheus.MustNewConstMetric(
			diskScrapeDesc,
			prometheus.GaugeValue,
			float64(global.DiskStat.DiskIOStat[diskName].ReadCount),
			diskName, "ReadCount",
		)
		ch <- prometheus.MustNewConstMetric(
			diskScrapeDesc,
			prometheus.GaugeValue,
			float64(global.DiskStat.DiskIOStat[diskName].MergedReadCount),
			diskName, "MergedReadCount",
		)
		ch <- prometheus.MustNewConstMetric(
			diskScrapeDesc,
			prometheus.GaugeValue,
			float64(global.DiskStat.DiskIOStat[diskName].WriteCount),
			diskName, "WriteCount",
		)
		ch <- prometheus.MustNewConstMetric(
			diskScrapeDesc,
			prometheus.GaugeValue,
			float64(global.DiskStat.DiskIOStat[diskName].MergedWriteCount),
			diskName, "MergedWriteCount",
		)
		ch <- prometheus.MustNewConstMetric(
			diskScrapeDesc,
			prometheus.GaugeValue,
			float64(global.DiskStat.DiskIOStat[diskName].ReadBytes),
			diskName, "ReadBytes",
		)
		ch <- prometheus.MustNewConstMetric(
			diskScrapeDesc,
			prometheus.GaugeValue,
			float64(global.DiskStat.DiskIOStat[diskName].WriteBytes),
			diskName, "WriteBytes",
		)
		ch <- prometheus.MustNewConstMetric(
			diskScrapeDesc,
			prometheus.GaugeValue,
			float64(global.DiskStat.DiskIOStat[diskName].ReadTime),
			diskName, "ReadTime",
		)
		ch <- prometheus.MustNewConstMetric(
			diskScrapeDesc,
			prometheus.GaugeValue,
			float64(global.DiskStat.DiskIOStat[diskName].WriteTime),
			diskName, "WriteTime",
		)
		ch <- prometheus.MustNewConstMetric(
			diskScrapeDesc,
			prometheus.GaugeValue,
			float64(global.DiskStat.DiskIOStat[diskName].IoTime),
			diskName, "IoTime",
		)
		ch <- prometheus.MustNewConstMetric(
			diskScrapeDesc,
			prometheus.GaugeValue,
			float64(global.DiskStat.DiskIOStat[diskName].WeightedIO),
			diskName, "WeightedIO",
		)
		ch <- prometheus.MustNewConstMetric(
			diskScrapeDesc,
			prometheus.GaugeValue,
			float64(global.DiskStat.DisksInfo[diskName].Total),
			diskName, "Total",
		)
		ch <- prometheus.MustNewConstMetric(
			diskScrapeDesc,
			prometheus.GaugeValue,
			float64(global.DiskStat.DisksInfo[diskName].Free),
			diskName, "Free",
		)
		ch <- prometheus.MustNewConstMetric(
			diskScrapeDesc,
			prometheus.GaugeValue,
			float64(global.DiskStat.DisksInfo[diskName].Used),
			diskName, "Used",
		)
		ch <- prometheus.MustNewConstMetric(
			diskScrapeDesc,
			prometheus.GaugeValue,
			global.DiskStat.DisksInfo[diskName].UsedPercent,
			diskName, "UsePercent",
		)
		ch <- prometheus.MustNewConstMetric(
			diskScrapeDesc,
			prometheus.GaugeValue,
			float64(global.DiskStat.DisksInfo[diskName].InodesTotal),
			diskName, "InodeTotal",
		)
		ch <- prometheus.MustNewConstMetric(
			diskScrapeDesc,
			prometheus.GaugeValue,
			float64(global.DiskStat.DisksInfo[diskName].InodesUsed),
			diskName, "InodesUsed",
		)
		ch <- prometheus.MustNewConstMetric(
			diskScrapeDesc,
			prometheus.GaugeValue,
			float64(global.DiskStat.DisksInfo[diskName].InodesFree),
			diskName, "InodesFree",
		)
		ch <- prometheus.MustNewConstMetric(
			diskScrapeDesc,
			prometheus.GaugeValue,
			float64(global.DiskStat.DisksInfo[diskName].InodesUsedPercent),
			diskName, "InodesUsedPercent",
		)
	}

	return nil
}
