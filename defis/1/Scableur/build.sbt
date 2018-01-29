name := "Scableur"

version := "0.1"

scalaVersion := "2.12.4"

libraryDependencies += "com.github.tototoshi" %% "scala-csv" % "1.3.5"

mainClass in assembly := Some("prog.comp2018.scableur.Scableur")

assemblyJarName in assembly := "Scableur.jar"

assemblyOutputPath in assembly := file("out/Scableur.jar")

enablePlugins(DockerPlugin)

dockerfile in docker := {
  val jarfile =(assemblyOutputPath in assembly).value
  val appDirPath = "/app"
  val jarTargetPath = s"$appDirPath/${jarfile.name}"

  new Dockerfile {
    from("java")
    add(jarfile,jarTargetPath)
    workDir(appDirPath)
    entryPoint("java", "-jar", jarTargetPath)
  }
}

