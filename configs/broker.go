package configs

// BrokerBaseConfig 定義 Broker 的設定
type BrokerBaseConfig struct {
	// Host 定義 Broker 的 host
	Host string `json:"host"`

	// Port 定義 Broker 的 port
	Port uint `json:"port"`
}
