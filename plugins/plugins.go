package plugins

import (
	"github.com/ibm/opentalaria/protocol"
	"github.com/ibm/opentalaria/utils"
	"github.com/spf13/viper"
)

type PluginInterface interface {
	Init(env *viper.Viper) error

	// settings
	GetSetting(key string) (string, error)
	SetSetting(key, value string) error

	// topics
	ListTopics(topicName []string) ([]protocol.MetadataResponseTopic, error)
	GetTopic(topicName string) (protocol.MetadataResponseTopic, error)
	AddTopic(topic protocol.CreatableTopic) utils.KError
	DeleteTopic(topic string) utils.KError

	// partitions
	CreatePartitions(topicName string, partitionCount int) error
}
