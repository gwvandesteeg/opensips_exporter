package processors

import (
	"github.com/VoIPGRID/opensips_exporter/opensips"
	"github.com/prometheus/client_golang/prometheus"
)

type CoreProcessor struct {
	statistics map[string]opensips.Statistic
}

var coreMetrics = map[string]metric{
	"rcv_requests":        newMetric("core", "requests_total", "Total number of received requests by OpenSIPS.", []string{}, prometheus.CounterValue),
	"rcv_replies":         newMetric("core", "replies_total", "Total number of received replies by OpenSIPS.", []string{}, prometheus.CounterValue),
	"fwd_requests":        newMetric("core", "requests", "Number of requests by OpenSIPS.", []string{"kind"}, prometheus.CounterValue),
	"fwd_replies":         newMetric("core", "replies", "Number of received replies by OpenSIPS.", []string{"kind"}, prometheus.CounterValue),
	"drop_requests":       newMetric("core", "requests", "Number of requests by OpenSIPS.", []string{"kind"}, prometheus.CounterValue),
	"drop_replies":        newMetric("core", "replies", "Number of received replies by OpenSIPS.", []string{"kind"}, prometheus.CounterValue),
	"err_requests":        newMetric("core", "requests", "Number of requests by OpenSIPS.", []string{"kind"}, prometheus.CounterValue),
	"err_replies":         newMetric("core", "replies", "Number of received replies by OpenSIPS.", []string{"kind"}, prometheus.CounterValue),
	"bad_URIs_rcvd":       newMetric("core", "bad_URIs_rcvd", "Number of URIs that OpenSIPS failed to parse.", []string{}, prometheus.CounterValue),
	"unsupported_methods": newMetric("core", "unsupported_methods", "Number of non-standard methods encountered by OpenSIPS while parsing SIP methods.", []string{}, prometheus.CounterValue),
	"bad_msg_hdr":         newMetric("core", "bad_msg_hdr", "Number of SIP headers that OpenSIPS failed to parse.", []string{}, prometheus.CounterValue),
	"timestamp":           newMetric("core", "uptime_seconds", "Number of seconds elapsed from OpenSIPS starting.", []string{}, prometheus.CounterValue),
}

func init() {
	for metric := range coreMetrics {
		Processors[metric] = coreProcessorFunc
	}
	Processors["core:"] = coreProcessorFunc
}

func coreProcessorFunc(s map[string]opensips.Statistic) prometheus.Collector {
	return &CoreProcessor{
		statistics: s,
	}
}

func (c CoreProcessor) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range coreMetrics {
		ch <- metric.Desc
	}
}

func (c CoreProcessor) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(
		coreMetrics["rcv_requests"].Desc,
		coreMetrics["rcv_requests"].ValueType,
		c.statistics["rcv_requests"].Value,
	)
	ch <- prometheus.MustNewConstMetric(
		coreMetrics["rcv_replies"].Desc,
		coreMetrics["rcv_replies"].ValueType,
		c.statistics["rcv_replies"].Value,
	)
	ch <- prometheus.MustNewConstMetric(
		coreMetrics["fwd_requests"].Desc,
		coreMetrics["fwd_requests"].ValueType,
		c.statistics["fwd_requests"].Value,
		"forwarded",
	)
	ch <- prometheus.MustNewConstMetric(
		coreMetrics["fwd_replies"].Desc,
		coreMetrics["fwd_replies"].ValueType,
		c.statistics["fwd_replies"].Value,
		"forwarded",
	)
	ch <- prometheus.MustNewConstMetric(
		coreMetrics["drop_requests"].Desc,
		coreMetrics["drop_requests"].ValueType,
		c.statistics["drop_requests"].Value,
		"dropped",
	)
	ch <- prometheus.MustNewConstMetric(
		coreMetrics["drop_replies"].Desc,
		coreMetrics["drop_replies"].ValueType,
		c.statistics["drop_replies"].Value,
		"dropped",
	)
	ch <- prometheus.MustNewConstMetric(
		coreMetrics["err_requests"].Desc,
		coreMetrics["err_requests"].ValueType,
		c.statistics["err_requests"].Value,
		"error",
	)
	ch <- prometheus.MustNewConstMetric(
		coreMetrics["err_replies"].Desc,
		coreMetrics["err_replies"].ValueType,
		c.statistics["err_replies"].Value,
		"error",
	)
	ch <- prometheus.MustNewConstMetric(
		coreMetrics["bad_URIs_rcvd"].Desc,
		coreMetrics["bad_URIs_rcvd"].ValueType,
		c.statistics["bad_URIs_rcvd"].Value,
	)
	ch <- prometheus.MustNewConstMetric(
		coreMetrics["unsupported_methods"].Desc,
		coreMetrics["unsupported_methods"].ValueType,
		c.statistics["unsupported_methods"].Value,
	)
	ch <- prometheus.MustNewConstMetric(
		coreMetrics["bad_msg_hdr"].Desc,
		coreMetrics["bad_msg_hdr"].ValueType,
		c.statistics["bad_msg_hdr"].Value,
	)
	ch <- prometheus.MustNewConstMetric(
		coreMetrics["timestamp"].Desc,
		coreMetrics["timestamp"].ValueType,
		c.statistics["timestamp"].Value,
	)
}