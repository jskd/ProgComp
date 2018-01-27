name := "Scableur"

version := "0.1"

scalaVersion := "2.12.4"

libraryDependencies += "com.github.tototoshi" %% "scala-csv" % "1.3.5"

mainClass in assembly := Some("prog.comp2018.scableur.Scableur")

assemblyJarName in assembly := "Scableur.jar"

assemblyOutputPath in assembly := file("out/Scableur.jar")