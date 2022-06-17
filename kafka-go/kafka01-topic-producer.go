/* Alta3 Research | RZFeeser
   Kafka - Kafka Producer - topic creator   */

package main

import (

    kafka "github.com/segmentio/kafka-go"
)


func main() {

        // to create topics when auto.create.topics.enable='false'
        topic := "mymessagelog"   // topic to be created

        conn, err := kafka.Dial("tcp", "localhost:9093")
        if err != nil {
            panic(err.Error())
        }
        defer conn.Close()

        var controllerConn *kafka.Conn
        controllerConn, err = kafka.Dial("tcp", "localhost:9093")            
        // error handling
        if err != nil {
            panic(err.Error())
        }
        defer controllerConn.Close()


        topicConfigs := []kafka.TopicConfig{
            kafka.TopicConfig{
                Topic:             topic,
                NumPartitions:     1,
                ReplicationFactor: 1,
            },
        }

        err = controllerConn.CreateTopics(topicConfigs...)
        if err != nil {
            panic(err.Error())
        }
}

