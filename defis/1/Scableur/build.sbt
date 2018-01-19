name := "Scableur"

version := "0.1"

scalaVersion := "2.12.4"

mainClass in assembly := Some("prog.comp2018.scableur.Scableur")

assemblyJarName in assembly := "Scableur.jar"

assemblyOutputPath in assembly := file("out/Scableur.jar")