package init

import (
	"log"

	"github.com/Shopify/sarama"
)

var (
	Client *sarama.Client
)

func init() {
	config := sarama.NewConfig()
	config.Version = sarama.V0_11_0_2                                                                                //实例化个sarama的Config
	config.Producer.Return.Successes = true                                                                          //是否开启消息发送成功后通知 successes channel
	config.Producer.Partitioner = sarama.NewRandomPartitioner                                                        //随机分区器
	client, err := sarama.NewClient([]string{"192.168.3.147:9092", "192.168.3.148:9092", "192.168.3.149:9092"}, config) //初始化客户端
	// defer client.Close()
	if err != nil {
		log.Fatal("连接kafka失败，程序停止。错误原因: ", err)
		panic(err)
	}
	Client = &client
}
