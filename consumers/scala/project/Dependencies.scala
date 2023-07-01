import sbt.*

object Dependencies {
  lazy val kafkaClientsDeps: List[ModuleID] =
    "org.apache.kafka" % "kafka-clients" % "3.4.0" :: Nil
}