package scableur

import org.apache.spark.{SparkConf, SparkContext}
import org.apache.spark.sql.SQLContext
import org.apache.spark.sql.types.{StructField, StructType, IntegerType, StringType}
import scableur.utils.Conf

/**
  * https://medium.com/@mrpowers/creating-a-spark-project-with-sbt-intellij-sbt-spark-package-and-friends-cc9108751c28
  */
object Scableur {
  // Spark Init
  val conf = new SparkConf().setAppName(Conf.AppName)
  val sc = new SparkContext(conf)
  
  def main(args: Array[String]): Unit = {
    //parse_input()
    spark_data_load(args(0))
    //println(s"there are $c lines in $filename")
  }


  def string_to_iValue(str: String): ((Int,Int), Int) = {
    val pattern = """[(](\d+),(\d+)[)] (\d+)""".r
    val parts=pattern.findAllIn(str)
    val i = parts.group(1).toInt
    val j = parts.group(2).toInt
    val v = parts.group(3).toInt
    return ((i,j), v)
  }

  def print_iValue(x : ((Int,Int), Int)) : Unit = {
    println( "( " + x._1._1 + ", " + x._1._2 + ") " + x._2)
  } 


  def spark_data_load(filename: String) : Unit = {
    val csvFile = sc.textFile(filename)
    val data = csvFile.map(line => line.split(";").map(elem => elem.trim))
    val data2 = data.map(line => line.map(x => string_to_iValue _ ))
    //val res = data2.collect().foreach(line => line.foreach(x => print_iValue _))
    //val data3 = data2.map(line => line.fold( to a string))
  }




  def check_args(args : Array[String]) : Unit = {
    if(args.length < 4 || args.length > 4){
      println("use : ./ws <data.csv> <user.txt> <view0.csv> <changes.txt>")
      System.exit(0)
    }

    Conf.Arguments.dataFile(args(0))
    Conf.Arguments.userFile(args(1))
    Conf.Arguments.viewFile(args(2))
    Conf.Arguments.changesFile(args(3))

    Conf.Arguments.print(Conf.DEBUG)
  }

  

}