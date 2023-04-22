package kafkaconfig

import "github.com/segmentio/kafka-go"

func NewKafkaReaderConfig(topic string,brokers []string,)(kafka.ReaderConfig){
    config:= kafka.ReaderConfig{
        Brokers: brokers,
        Topic: topic,
        WatchPartitionChanges: true,
        Partition: 0,
        MinBytes: 10e3,
        MaxBytes: 10e6,
    }
    
    return config;
}
