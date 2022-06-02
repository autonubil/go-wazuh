package ossec

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

//Define a struct for you collector that contains pointers
//to prometheus descriptors for each metric you wish to expose.
//Note you can also include fields of other types if they provide utility
//but we just won't be exposing them as metrics.
type agentCollector struct {
	agent                       *Client
	connectedMetric             *prometheus.Desc
	eventsTotalMetric           *prometheus.Desc
	messagesSentTotalMetric     *prometheus.Desc
	messagesReceivedTotalMetric *prometheus.Desc

	sentBytesTotalMetric        *prometheus.Desc
	sentBytesSessionTotalMetric *prometheus.Desc

	receivedBytesTotalMetric        *prometheus.Desc
	receivedBytesSessionTotalMetric *prometheus.Desc

	connectionAttemptsTotalMetric *prometheus.Desc
	connectionsOpenedTotalMetric  *prometheus.Desc
	connectionsClosedTotalMetric  *prometheus.Desc
}

var agentLabels = []string{"agent_id", "agent_name", "agent_ip", "agent_protocol", "agent_encryption", "session_id"}

//You must create a constructor for you collector that
//initializes every descriptor and returns a pointer to the collector
func newAgentCollector(agent *Client) *agentCollector {
	return &agentCollector{
		agent: agent,
		connectedMetric: prometheus.NewDesc("ossec_agent_connections_total",
			"agent connection count (normally 0 or 1)",
			agentLabels, nil,
		),

		eventsTotalMetric: prometheus.NewDesc("ossec_agent_events_total",
			"total events processes",
			agentLabels, nil,
		),
		messagesSentTotalMetric: prometheus.NewDesc("ossec_agent_messages_sent_total",
			"total messages sent",
			agentLabels, nil,
		),
		messagesReceivedTotalMetric: prometheus.NewDesc("ossec_agent_messages_received_total",
			"total messages received",
			agentLabels, nil,
		),

		sentBytesTotalMetric: prometheus.NewDesc("ossec_agent_bytes_sent_total",
			"total bytes sent",
			agentLabels, nil,
		),
		sentBytesSessionTotalMetric: prometheus.NewDesc("ossec_agent_session_bytes_sent_total",
			"bytes sent during current session",
			agentLabels, nil,
		),

		receivedBytesTotalMetric: prometheus.NewDesc("ossec_agent_bytes_received_total",
			"total bytes received",
			agentLabels, nil,
		),
		receivedBytesSessionTotalMetric: prometheus.NewDesc("ossec_agent_session_bytes_received_total",
			"bytes received during current session",
			agentLabels, nil,
		),

		connectionAttemptsTotalMetric: prometheus.NewDesc("ossec_agent_connection_attempts_total",
			"total connection attempts",
			agentLabels, nil,
		),

		connectionsOpenedTotalMetric: prometheus.NewDesc("ossec_agent_connections_openend_total",
			"total succesfull connection attempts",
			agentLabels, nil,
		),

		connectionsClosedTotalMetric: prometheus.NewDesc("ossec_agent_connections_closed_total",
			"total connection activly closed ",
			agentLabels, nil,
		),
	}
}

//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *agentCollector) Describe(ch chan<- *prometheus.Desc) {
	//Update this section with the each metric you create for a given collector
	ch <- collector.connectedMetric
	ch <- collector.eventsTotalMetric
	ch <- collector.messagesSentTotalMetric
	ch <- collector.messagesReceivedTotalMetric

	ch <- collector.sentBytesTotalMetric
	ch <- collector.sentBytesSessionTotalMetric

	ch <- collector.receivedBytesTotalMetric
	ch <- collector.receivedBytesSessionTotalMetric

	ch <- collector.connectionAttemptsTotalMetric
	ch <- collector.connectionsOpenedTotalMetric
	ch <- collector.connectionsClosedTotalMetric
}

//Collect implements required collect function for all promehteus collectors
func (collector *agentCollector) Collect(ch chan<- prometheus.Metric) {
	//Implement logic here to determine proper metric value to return to prometheus
	//for each descriptor or call other functions that do so.
	if collector.agent != nil {
		var protocol string
		if collector.agent.UDP {
			protocol = "UDP"
		} else {
			protocol = "TCP"
		}
		var encryption string
		switch collector.agent.EncryptionMethod {
		case EncryptionMethodBlowFish:
			encryption = "BlowFish"
		case EncryptionMethodAES:
			encryption = "AES"
		default:
			encryption = "Unknown"
		}
		var sessionID string
		if collector.agent.sessionID == 0 {
			sessionID = "none"
		} else {
			sessionID = strconv.FormatInt(collector.agent.sessionID, 16)
		}
		agentLabeValues := []string{collector.agent.AgentID, collector.agent.AgentName, collector.agent.AgentIP, protocol, encryption, sessionID}

		var connections float64
		if collector.agent.IsConencted() {
			connections = 1
		}
		ch <- prometheus.MustNewConstMetric(collector.connectedMetric, prometheus.CounterValue, connections, agentLabeValues...)
		ch <- prometheus.MustNewConstMetric(collector.eventsTotalMetric, prometheus.CounterValue, float64(collector.agent.evtCount), agentLabeValues...)
		ch <- prometheus.MustNewConstMetric(collector.messagesSentTotalMetric, prometheus.CounterValue, float64(collector.agent.sentCount), agentLabeValues...)
		ch <- prometheus.MustNewConstMetric(collector.messagesReceivedTotalMetric, prometheus.CounterValue, float64(collector.agent.receivedCount), agentLabeValues...)

		ch <- prometheus.MustNewConstMetric(collector.sentBytesTotalMetric, prometheus.CounterValue, float64(collector.agent.sentBytes), agentLabeValues...)
		ch <- prometheus.MustNewConstMetric(collector.sentBytesSessionTotalMetric, prometheus.CounterValue, float64(collector.agent.sentBytesTotal), agentLabeValues...)

		ch <- prometheus.MustNewConstMetric(collector.receivedBytesTotalMetric, prometheus.CounterValue, float64(collector.agent.receivedBytes), agentLabeValues...)
		ch <- prometheus.MustNewConstMetric(collector.receivedBytesSessionTotalMetric, prometheus.CounterValue, float64(collector.agent.receivedBytesTotal), agentLabeValues...)

		ch <- prometheus.MustNewConstMetric(collector.connectionAttemptsTotalMetric, prometheus.CounterValue, float64(collector.agent.connectionAttempts), agentLabeValues...)
		ch <- prometheus.MustNewConstMetric(collector.connectionsOpenedTotalMetric, prometheus.CounterValue, float64(collector.agent.connectionsOpened), agentLabeValues...)
		ch <- prometheus.MustNewConstMetric(collector.connectionsClosedTotalMetric, prometheus.CounterValue, float64(collector.agent.connectionsOpened), agentLabeValues...)

	}
}
