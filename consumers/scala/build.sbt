ThisBuild / scalaVersion := "2.13.10"

lazy val root = (project in file("."))
  .settings(
    name := "Consumer",
    version := "1.0",
    libraryDependencies ++= Seq(
      "org.apache.kafka" % "kafka-clients" % "2.8.1",
      "org.apache.spark" %% "spark-core" % "3.4.1",
    )
  )
