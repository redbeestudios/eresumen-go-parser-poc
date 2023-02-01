package pkg

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Config struct {
	DatabaseService    *DatabaseServiceConfig    `json:"database_service"`
	EmailConsumer      *EmailConsumerConfig      `json:"email_consumer"`
	FileParserProducer *FileParserProducerConfig `json:"file_parser_producer"`
	EmailMockServer    *EmailMockServerConfig    `json:"email_mock_server"`
}

type DatabaseServiceConfig struct {
	Server  serverConfig        `json:"server"`
	Kafka   kafkaConsumerConfig `json:"kafka"`
	Mongodb mongoDbConfig       `json:"mongodb"`
	Sql     sqlDbConfig         `json:"sql"`
}

type EmailConsumerConfig struct {
	Kafka kafkaConsumerConfig `json:"kafka"`
	Smtp  smtpClientConfig    `json:"smtp"`
}

type FileParserProducerConfig struct {
	Server serverConfig        `json:"server"`
	Kafka  kafkaProducerConfig `json:"kafka"`
}

type EmailMockServerConfig struct {
	Enabled           bool        `json:"enabled"`
	Address           string      `json:"address"`
	Port              string      `json:"port"`
	ReadTimeout       int         `json:"read_timeout"`
	WriteTimeout      int         `json:"write_timeout"`
	MaxMessageBytes   int         `json:"max_message_bytes"`
	MaxRecipients     int         `json:"max_recipients"`
	AllowInsecureAuth bool        `json:"allow_insecure_auth"`
	Redis             redisConfig `json:"redis"`
}

type serverConfig struct {
	Address      string `json:"address"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
	IdleTimeout  int    `json:"idle_timeout"`
}

type restClientConfig struct {
	BaseUrl               string `json:"base_url"`
	TimeoutMilliseconds   int    `json:"timeout_milliseconds"`
	RetryCount            int    `json:"retry_count"`
	RetryWaitMilliseconds int    `json:"retry_wait_milliseconds"`
}

type kafkaConsumerConfig struct {
	BootstrapServers      string `json:"bootstrap_servers"`
	GroupId               string `json:"group_id"`
	SessionTimeoutMs      int    `json:"session_timeout_ms"`
	AutoOffsetReset       string `json:"auto_offset_reset"`
	EnableAutoOffsetStore bool   `json:"enable_auto_offset_store"`
	Topic                 string `json:"topic"`
}

type kafkaProducerConfig struct {
	BootstrapServers string `json:"bootstrap_servers"`
	Topic            string `json:"topic"`
}

type mongoDbConfig struct {
	Uri                 string `json:"uri"`
	SessionTimeoutMs    int    `json:"session_timeout_ms"`
	DatabaseName        string `json:"database_name"`
	IndexCollectionName string `json:"index_collection_name"`
	UserCollectionName  string `json:"user_collection_name"`
}

type sqlDbConfig struct {
	Uri              string `json:"uri"`
	SessionTimeoutMs int    `json:"session_timeout_ms"`
	IndexTableName   string `json:"index_table_name"`
	UserTableName    string `json:"user_table_name"`
}

type redisConfig struct {
	Address string `json:"address"`
}

type smtpClientConfig struct {
	Url string `json:"url"`
}

func (k *kafkaConsumerConfig) ConfigMap() *kafka.ConfigMap {

	return &kafka.ConfigMap{
		"bootstrap.servers":        k.BootstrapServers,
		"group.id":                 k.GroupId,
		"session.timeout.ms":       k.SessionTimeoutMs,
		"auto.offset.reset":        k.AutoOffsetReset,
		"enable.auto.offset.store": k.EnableAutoOffsetStore,
	}
}

func (k *kafkaProducerConfig) ConfigMap() *kafka.ConfigMap {

	return &kafka.ConfigMap{
		"bootstrap.servers": k.BootstrapServers,
	}
}

func (k *kafkaConsumerConfig) Topics() []string {
	return []string{k.Topic}
}
