name := "Test_Spark"

version := "0.1"

scalaVersion in ThisBuild := "2.11.7"

sparkVersion := "2.2.0"
sparkComponents ++= Seq("sql")