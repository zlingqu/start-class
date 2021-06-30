package producer

import (
	kafkaInit "begin_training/kafka/init"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func ProductMessToKafka(topic, msg string) error {
	producer, _ := sarama.NewSyncProducerFromClient(*kafkaInit.Client)
	partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(msg)})
	if err != nil {
		log.Fatalf("无法生成消息，错误原因: %q", err)
		return err
	}
	fmt.Println("写入到partition: ", partition)
	fmt.Println("当前的offset是：", offset)
	return nil
}
