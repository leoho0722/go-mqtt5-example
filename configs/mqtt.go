package configs

import (
	mqtt "github.com/eclipse/paho.golang/paho"
)

// MQTTConfig 定義 MQTT client 的基本設定
type MQTTConfig struct {

	// Broker 定義 MQTT broker 的設定
	Broker BrokerBaseConfig `json:"broker"`

	// ClientID 定義 MQTT client 的 ID
	ClientID string `json:"clientID"`

	// Topic 定義 MQTT 訊息的 topic
	Topic string `json:"topic"`

	// QosLevel 定義 MQTT 訊息的 QoS level
	QosLevel uint `json:"qosLevel"`

	// IsSharedSubscriptionEnabled 定義是否啟用 MQTT Shared Subscription
	IsSharedSubscriptionEnabled bool `json:"isSharedSubscriptionEnabled"`
}

// PublisherConfig 定義 MQTT publisher 的設定
type PublisherConfig struct {
	MQTTConfig

	// Message 定義要發佈的 MQTT 訊息
	Message string `json:"message"`
}

// SubscriberConfig 定義 MQTT subscriber 的設定
type SubscriberConfig struct {
	MQTTConfig
}

type BaseClient interface {
	// NewClientConfig 建立 MQTT client config
	NewClientConfig() *mqtt.ClientConfig

	// NewClient 建立 MQTT client
	NewClient() *mqtt.Client

	// Connect 連線至 MQTT broker
	Connect() (*mqtt.Client, error)

	// Disconnect 中斷與 MQTT broker 的連線
	Disconnect() error
}

type PublisherClient interface {
	BaseClient

	// Publish 發佈訊息至 MQTT broker
	Publish(message string) error
}

type SubscriberClient interface {
	BaseClient

	// Subscribe 訂閱 MQTT 訊息
	Subscribe() error
}
