package ossec

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Define a struct for you collector that contains pointers
// to prometheus descriptors for each metric you wish to expose.
// Note you can also include fields of other types if they provide utility
// but we just won't be exposing them as metrics.
type agentCollector struct {
	agents                      []*Client
	connectedMetric             *prometheus.GaugeVec
	queueSizeMetric             *prometheus.GaugeVec
	eventsTotalMetric           *prometheus.CounterVec
	messagesSentTotalMetric     *prometheus.CounterVec
	messagesReceivedTotalMetric *prometheus.CounterVec
	messageErrorsTotalMetric    *prometheus.CounterVec

	bytesSentTotalMetric     *prometheus.CounterVec
	bytesReceivedTotalMetric *prometheus.CounterVec

	connectionAttemptsTotalMetric *prometheus.CounterVec
	connectionsOpenedTotalMetric  *prometheus.CounterVec
	connectionsClosedTotalMetric  *prometheus.CounterVec
}

var agentLabels = []string{"agent_id", "agent_name", "agent_protocol", "agent_encryption"}

var AgentCollector = newAgentCollector()

func init() {
	prometheus.Register(AgentCollector)
}

// You must create a constructor for you collector that
// initializes every descriptor and returns a pointer to the collector
func newAgentCollector() *agentCollector {
	return &agentCollector{
		agents: make([]*Client, 0),
		connectedMetric: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "ossec",
			Subsystem: "agent",
			Name:      "connections",
			Help:      "agent connection count",
		}, agentLabels),

		queueSizeMetric: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "ossec",
			Subsystem: "agent",
			Name:      "queued_count",
			Help:      "total messages in queued",
		}, agentLabels),
		eventsTotalMetric: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "ossec",
			Subsystem: "agent",
			Name:      "events_total",
			Help:      "total events processed",
		}, agentLabels),

		messageErrorsTotalMetric: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "ossec",
			Subsystem: "agent",
			Name:      "message_errors_total",
			Help:      "total errors in messages",
		}, agentLabels),

		messagesSentTotalMetric: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "ossec",
			Subsystem: "agent",
			Name:      "messages_sent_total",
			Help:      "total messages sent",
		}, agentLabels),

		messagesReceivedTotalMetric: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "ossec",
			Subsystem: "agent",
			Name:      "messages_received_total",
			Help:      "total bytes received",
		}, agentLabels),

		bytesSentTotalMetric: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "ossec",
			Subsystem: "agent",
			Name:      "bytes_sent_total",
			Help:      "total bytes sent",
		}, agentLabels),

		bytesReceivedTotalMetric: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "ossec",
			Subsystem: "agent",
			Name:      "bytes_received_total",
			Help:      "total bytes received",
		}, agentLabels),

		connectionAttemptsTotalMetric: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "ossec",
			Subsystem: "agent",
			Name:      "connection_attempts_total",
			Help:      "total connection attempts",
		}, agentLabels),

		connectionsOpenedTotalMetric: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "ossec",
			Subsystem: "agent",
			Name:      "connections_opened_total",
			Help:      "succesful connection attempts",
		}, agentLabels),

		connectionsClosedTotalMetric: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "ossec",
			Subsystem: "agent",
			Name:      "connections_closed_total",
			Help:      "total connection actively closed",
		}, agentLabels),
	}
}

func (collector *agentCollector) getAgentLabels(agent *Client) []string {
	if agent != nil {
		var protocol string
		if agent.UDP {
			protocol = "UDP"
		} else {
			protocol = "TCP"
		}
		var encryption string
		switch agent.EncryptionMethod {
		case EncryptionMethodBlowFish:
			encryption = "BlowFish"
		case EncryptionMethodAES:
			encryption = "AES"
		default:
			encryption = "Unknown"
		}
		/* TODO: Implement session stats?
		var sessionID string
		if agent.sessionID == 0 {
			sessionID = "none"
		} else {
			sessionID = strconv.FormatInt(agent.sessionID, 16)
		}
		*/
		return []string{agent.AgentID, agent.AgentName, protocol, encryption}
	}
	return nil
}

func (collector *agentCollector) Register(agent *Client) {
	if agent == nil {
		return
	}
	labels := collector.getAgentLabels(agent)
	if labels == nil {
		return
	}
	collector.connectedMetric.WithLabelValues(labels...).Set(0)
	collector.queueSizeMetric.WithLabelValues(labels...).Set(0)

	collector.eventsTotalMetric.WithLabelValues(labels...).Add(0)
	collector.messagesSentTotalMetric.WithLabelValues(labels...).Add(0)
	collector.messagesReceivedTotalMetric.WithLabelValues(labels...).Add(0)
	collector.messageErrorsTotalMetric.WithLabelValues(labels...).Add(0)
	collector.bytesSentTotalMetric.WithLabelValues(labels...).Add(0)
	collector.bytesReceivedTotalMetric.WithLabelValues(labels...).Add(0)
	collector.connectionAttemptsTotalMetric.WithLabelValues(labels...).Add(0)
	collector.connectionsOpenedTotalMetric.WithLabelValues(labels...).Add(0)
	collector.connectionsClosedTotalMetric.WithLabelValues(labels...).Add(0)
}

func (collector *agentCollector) TryConnect(agent *Client) {
	if agent == nil {
		return
	}
	labels := collector.getAgentLabels(agent)
	if labels == nil {
		return
	}
	collector.connectionAttemptsTotalMetric.WithLabelValues(labels...).Add(1)
}

func (collector *agentCollector) Connect(agent *Client) {
	if agent == nil {
		return
	}
	labels := collector.getAgentLabels(agent)
	if labels == nil {
		return
	}
	collector.connectedMetric.WithLabelValues(labels...).Set(1)
	collector.connectionsOpenedTotalMetric.WithLabelValues(labels...).Add(1)
}

func (collector *agentCollector) Disconnect(agent *Client) {
	if agent == nil {
		return
	}
	labels := collector.getAgentLabels(agent)
	if labels == nil {
		return
	}
	collector.connectedMetric.WithLabelValues(labels...).Set(0)
}

func (collector *agentCollector) Unregister(agent *Client) {
	if agent == nil {
		return
	}
	labels := collector.getAgentLabels(agent)
	if labels == nil {
		return
	}
	collector.connectedMetric.DeleteLabelValues(labels...)
}

func (collector *agentCollector) EventsReceived(agent *Client, count int) {
	if count == 0 {
		return
	}
	if agent == nil {
		return
	}
	labels := collector.getAgentLabels(agent)
	if labels == nil {
		return
	}
	collector.eventsTotalMetric.WithLabelValues(labels...).Add(float64(count))
}

func (collector *agentCollector) MessagesReceived(agent *Client, count int) {
	if count == 0 {
		return
	}
	if agent == nil {
		return
	}
	labels := collector.getAgentLabels(agent)
	if labels == nil {
		return
	}
	collector.messagesReceivedTotalMetric.WithLabelValues(labels...).Add(float64(count))
}

func (collector *agentCollector) MessageError(agent *Client, count int) {
	if count == 0 {
		return
	}
	if agent == nil {
		return
	}
	labels := collector.getAgentLabels(agent)
	if labels == nil {
		return
	}
	collector.messageErrorsTotalMetric.WithLabelValues(labels...).Add(float64(count))
}

func (collector *agentCollector) BytesSent(agent *Client, count int) {
	if count == 0 {
		return
	}
	if agent == nil {
		return
	}
	labels := collector.getAgentLabels(agent)
	if labels == nil {
		return
	}
	collector.bytesSentTotalMetric.WithLabelValues(labels...).Add(float64(count))
}

func (collector *agentCollector) BytesRead(agent *Client, count int) {
	if count == 0 {
		return
	}
	if agent == nil {
		return
	}
	labels := collector.getAgentLabels(agent)
	if labels == nil {
		return
	}
	collector.bytesReceivedTotalMetric.WithLabelValues(labels...).Add(float64(count))
}

func (collector *agentCollector) MessagesSent(agent *Client, count int) {
	if count == 0 {
		return
	}
	if agent == nil {
		return
	}
	labels := collector.getAgentLabels(agent)
	if labels == nil {
		return
	}
	collector.messagesSentTotalMetric.WithLabelValues(labels...).Add(float64(count))
}

func (collector *agentCollector) SetQueueSize(agent *Client, size int) {
	if agent == nil {
		return
	}
	labels := collector.getAgentLabels(agent)
	if labels == nil {
		return
	}
	collector.queueSizeMetric.WithLabelValues(labels...).Set(float64(size))
}

func (collector *agentCollector) Enqueue(agent *Client) {
	if agent == nil {
		return
	}
	labels := collector.getAgentLabels(agent)
	if labels == nil {
		return
	}
	collector.queueSizeMetric.WithLabelValues(labels...).Inc()
}

func (collector *agentCollector) Dequeue(agent *Client) {
	if agent == nil {
		return
	}
	labels := collector.getAgentLabels(agent)
	if labels == nil {
		return
	}
	collector.queueSizeMetric.WithLabelValues(labels...).Dec()
}

// Each and every collector must implement the Describe function.
// It essentially writes all descriptors to the prometheus desc channel.
func (collector *agentCollector) Describe(ch chan<- *prometheus.Desc) {
	//Update this section with the each metric you create for a given collector
	collector.connectedMetric.Describe(ch)
	collector.queueSizeMetric.Describe(ch)
	collector.eventsTotalMetric.Describe(ch)
	collector.messagesSentTotalMetric.Describe(ch)
	collector.messagesReceivedTotalMetric.Describe(ch)
	collector.messageErrorsTotalMetric.Describe(ch)
	collector.bytesSentTotalMetric.Describe(ch)
	collector.bytesReceivedTotalMetric.Describe(ch)
	collector.connectionAttemptsTotalMetric.Describe(ch)
	collector.connectionsOpenedTotalMetric.Describe(ch)
	collector.connectionsClosedTotalMetric.Describe(ch)

}

// Collect implements required collect function for all promehteus collectors
func (collector *agentCollector) Collect(ch chan<- prometheus.Metric) {
	collector.connectedMetric.Collect(ch)
	collector.queueSizeMetric.Collect(ch)
	collector.eventsTotalMetric.Collect(ch)
	collector.messagesSentTotalMetric.Collect(ch)
	collector.messagesReceivedTotalMetric.Collect(ch)
	collector.messageErrorsTotalMetric.Collect(ch)
	collector.bytesSentTotalMetric.Collect(ch)
	collector.bytesReceivedTotalMetric.Collect(ch)
	collector.connectionAttemptsTotalMetric.Collect(ch)
	collector.connectionsOpenedTotalMetric.Collect(ch)
	collector.connectionsClosedTotalMetric.Collect(ch)
}
