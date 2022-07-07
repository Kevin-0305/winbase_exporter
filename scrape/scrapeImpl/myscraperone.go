package scrapeImpl

import (
	"context"
	"winbase_exporter/global"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Subsystem.
	myscraperonesubs = "myscraperonesubs"
)

var (
	myscraperoneDesc = prometheus.NewDesc(
		prometheus.BuildFQName(global.Namespace, myscraperonesubs, "args"),
		"my scraper one test args .",
		[]string{"args1", "args2"}, nil,
	)
)

type MyScraperOne struct{}

// Name of the Scraper. Should be unique.
func (MyScraperOne) Name() string {
	return myscraperonesubs
}

// Help describes the role of the Scraper.
func (MyScraperOne) Help() string {
	return "my scraper one"
}
func (MyScraperOne) Scrape(ctx context.Context, dc string, ch chan<- prometheus.Metric) error {
	//get some from datacentor
	//dc.dosomthing...
	//may be return error，will stop this scrape's register
	//return error
	ch <- prometheus.MustNewConstMetric(
		myscraperoneDesc, prometheus.CounterValue, 1, "argsone1", "argsone2",
	)
	ch <- prometheus.MustNewConstMetric(
		global.NewDesc(myscraperonesubs, "subsystemnameone", "Generic metric"),
		prometheus.UntypedValue,
		0.01,
	)
	return nil
}
