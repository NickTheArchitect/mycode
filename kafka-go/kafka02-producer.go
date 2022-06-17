/* Alta3 Research | RZFeeser
   Kafka - Producer talking to kafka broker  */

package main

import (
    "context"
    "fmt"
    "time"

    "github.com/google/uuid"
    kafka "github.com/segmentio/kafka-go"
)

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
    return &kafka.Writer{
        Addr:     kafka.TCP(kafkaURL),
        Topic:    topic,
        Balancer: &kafka.LeastBytes{},
    }
}

func main() {
    // get kafka writer using environment variables.
    //kafkaURL := os.Getenv("kafkaURL")
    //topic := os.Getenv("topic")
    kafkaURL := "localhost:9093"
    topic := "mymessagelog"

    writer := newKafkaWriter(kafkaURL, topic)
    defer writer.Close()
    fmt.Println("start producing ... !!")
    for i := 0; ; i++ {
        key := fmt.Sprintf("Key-%d", i)
        msg := kafka.Message{
            Key:   []byte(key),
            Value: []byte(fmt.Sprint(uuid.New())),
        }
        err := writer.WriteMessages(context.Background(), msg)
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println("produced", key)
        }
        time.Sleep(1 * time.Second)
    }
}

