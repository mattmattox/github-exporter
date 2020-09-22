package exporter

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

// AddMetrics - Add's all of the metrics to a map of strings, returns the map.
func AddMetrics() map[string]*prometheus.Desc {

	APIMetrics := make(map[string]*prometheus.Desc)

	APIMetrics["Stars"] = prometheus.NewDesc(
		prometheus.BuildFQName("github", "repo", "stars"),
		"Total number of Stars for given repository",
		[]string{"repo", "user", "private", "fork", "archived", "license", "language"}, nil,
	)
	APIMetrics["OpenIssues"] = prometheus.NewDesc(
		prometheus.BuildFQName("github", "repo", "open_issues"),
		"Total number of open issues for given repository",
		[]string{"repo", "user", "private", "fork", "archived", "license", "language"}, nil,
	)
	APIMetrics["PullRequestCount"] = prometheus.NewDesc(
		prometheus.BuildFQName("github", "repo", "pull_request_count"),
		"Total number of pull requests for given repository",
		[]string{"repo", "user", "private", "fork", "archived", "license", "language"}, nil,
	)
	APIMetrics["Watchers"] = prometheus.NewDesc(
		prometheus.BuildFQName("github", "repo", "watchers"),
		"Total number of watchers/subscribers for given repository",
		[]string{"repo", "user", "private", "fork", "archived", "license", "language"}, nil,
	)
	APIMetrics["Forks"] = prometheus.NewDesc(
		prometheus.BuildFQName("github", "repo", "forks"),
		"Total number of forks for given repository",
		[]string{"repo", "user", "private", "fork", "archived", "license", "language"}, nil,
	)
	APIMetrics["Size"] = prometheus.NewDesc(
		prometheus.BuildFQName("github", "repo", "size_kb"),
		"Size in KB for given repository",
		[]string{"repo", "user", "private", "fork", "archived", "license", "language"}, nil,
	)
	APIMetrics["ReleaseDownloads"] = prometheus.NewDesc(
		prometheus.BuildFQName("github", "repo", "release_downloads"),
		"Download count for a given release",
		[]string{"repo", "user", "release", "name", "created_at"}, nil,
	)
	APIMetrics["Limit"] = prometheus.NewDesc(
		prometheus.BuildFQName("github", "rate", "limit"),
		"Number of API queries allowed in a 60 minute window",
		[]string{}, nil,
	)
	APIMetrics["Remaining"] = prometheus.NewDesc(
		prometheus.BuildFQName("github", "rate", "remaining"),
		"Number of API queries remaining in the current window",
		[]string{}, nil,
	)
	APIMetrics["Reset"] = prometheus.NewDesc(
		prometheus.BuildFQName("github", "rate", "reset"),
		"The time at which the current rate limit window resets in UTC epoch seconds",
		[]string{}, nil,
	)

	return APIMetrics
}

// processMetrics - processes the response data and sets the metrics using it as a source
func (e *Exporter) processMetrics(ch chan<- prometheus.Metric) error {

	// TODO LIKE DIS
	// commonLabels := [*x.Base.Name, *x.Base.Owner.Name, strconv.FormatBool(*x.Base.Private), strconv.FormatBool(*x.Base.Fork), strconv.FormatBool(*x.Base.Archived), *x.Base.License.Key, *x.Base.Language]

	// Range through Repository metrics
	for _, x := range e.Repositories {
		ch <- prometheus.MustNewConstMetric(e.APIMetrics["Stars"], prometheus.GaugeValue, float64(*x.Base.StargazersCount), *x.Base.Name, *x.Base.Owner.Name, strconv.FormatBool(*x.Base.Private), strconv.FormatBool(*x.Base.Fork), strconv.FormatBool(*x.Base.Archived), *x.Base.License.Key, *x.Base.Language)
		ch <- prometheus.MustNewConstMetric(e.APIMetrics["Forks"], prometheus.GaugeValue, float64(*x.Base.ForksCount), *x.Base.Name, *x.Base.Owner.Name, strconv.FormatBool(*x.Base.Private), strconv.FormatBool(*x.Base.Fork), strconv.FormatBool(*x.Base.Archived), *x.Base.License.Key, *x.Base.Language)
		ch <- prometheus.MustNewConstMetric(e.APIMetrics["Watchers"], prometheus.GaugeValue, float64(*x.Base.WatchersCount), *x.Base.Name, *x.Base.Owner.Name, strconv.FormatBool(*x.Base.Private), strconv.FormatBool(*x.Base.Fork), strconv.FormatBool(*x.Base.Archived), *x.Base.License.Key, *x.Base.Language)
		ch <- prometheus.MustNewConstMetric(e.APIMetrics["Size"], prometheus.GaugeValue, float64(*x.Base.Size), *x.Base.Name, *x.Base.Owner.Name, strconv.FormatBool(*x.Base.Private), strconv.FormatBool(*x.Base.Fork), strconv.FormatBool(*x.Base.Archived), *x.Base.License.Key, *x.Base.Language)
		ch <- prometheus.MustNewConstMetric(e.APIMetrics["PullRequestCount"], prometheus.GaugeValue, float64(x.PullsCount), *x.Base.Name, *x.Base.Owner.Name, strconv.FormatBool(*x.Base.Private), strconv.FormatBool(*x.Base.Fork), strconv.FormatBool(*x.Base.Archived), *x.Base.License.Key, *x.Base.Language)
		ch <- prometheus.MustNewConstMetric(e.APIMetrics["OpenIssues"], prometheus.GaugeValue, float64(*x.Base.OpenIssuesCount), *x.Base.Name, *x.Base.Owner.Name, strconv.FormatBool(*x.Base.Private), strconv.FormatBool(*x.Base.Fork), strconv.FormatBool(*x.Base.Archived), *x.Base.License.Key, *x.Base.Language)
	}

	// Set Rate limit stats
	// ch <- prometheus.MustNewConstMetric(e.APIMetrics["Limit"], prometheus.GaugeValue, rate.Limit)
	// ch <- prometheus.MustNewConstMetric(e.APIMetrics["Remaining"], prometheus.GaugeValue, rate.Remaining)
	// ch <- prometheus.MustNewConstMetric(e.APIMetrics["Reset"], prometheus.GaugeValue, rate.Reset)

	return nil
}
