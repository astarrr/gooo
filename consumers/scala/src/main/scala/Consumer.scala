import java.util.Properties
import org.apache.kafka.clients.consumer.{ConsumerConfig, KafkaConsumer}
import org.apache.kafka.common.serialization.StringDeserializer

import java.time.Duration
import scala.util.control.Exception

object Consumer {
  def main(args: Array[String]): Unit = {
    val brokers = "kafka:9092"
    val topic = sys.env("CHANNEL")

    val props = new Properties()
    props.put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, brokers)
    props.put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, brokers)
    props.put(
      ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG,
      classOf[StringDeserializer].getName
    )
    props.put(
      ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG,
      classOf[StringDeserializer].getName
    )
    props.put(ConsumerConfig.GROUP_ID_CONFIG, "consumer")

    val consumer = new KafkaConsumer(props)
    try {
      consumer.subscribe(java.util.Collections.singletonList(topic))

      while (true) {
        val records = consumer.poll(Duration.ofMillis(10))
        records.forEach(record => println(s"Received message: ${record.value()}"))
      }
    }
    finally {
      consumer.close()
    }
  }
}
