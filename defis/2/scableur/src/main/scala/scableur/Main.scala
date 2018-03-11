package scableur

import org.apache.log4j.Logger;
import org.apache.log4j.PropertyConfigurator;
import org.apache.spark.{SparkConf, SparkContext}
import org.apache.spark.rdd.RDD
import org.apache.spark.sql.SQLContext
import org.apache.spark.sql.types.{StructField, StructType, IntegerType, StringType}
import scableur.utils.Conf
import scableur.data._
import scableur.eval._
import java.io._

object Scableur {
  var conf : SparkConf = null
  var sc : SparkContext = null
  var evaluator : Evaluator = null

  
  def main(args: Array[String]): Unit = {
    Conf.logger = Logger.getLogger("DEBUG");
    Conf.logger.info("Launching Scableur")
    check_args(args)
    config_environment()
    parse_input() 
    val file = spark_data_load(Conf.Arguments.dataFile)
    map_data(file)
    print_view_csv(file)
    print_changes()

    //cycle_detection()
    //evaluation_flow()
    //print_to_file()

    //TODO: users
    
    //Debug
    //file.foreach(line => line.foreach(println _))
    Conf.logger.info("Scableur : done.")
  }


  def parse_input() : Unit = {
    // Parse input to add position
    addPositionsToCSV(Conf.Arguments.dataFile, Conf.positionDataCSV)
    Conf.Arguments.dataFile(Conf.positionDataCSV)
  }


  def spark_data_load(filename: String) = {
    // Load data in spark file system
    val csvFile = sc.textFile(filename)
    csvFile.map(line => line.split(";").map(elem => elem.trim))
  }


  def map_data(file : RDD[Array[String]]) : Unit = {
    //file.foreach(line => line.foreach(println _))
    val mapped_data = file.map(line => line.map(elem => evaluator.evalPValue(elem) ))
    println("mapped data : " + mapped_data)
    //mapped_data.foreach(line => line.foreach(println _)) 
    val functionList = evaluator.getFormulasList()
    println("functionList : " + functionList)

  }

  def print_view_csv(file : RDD[Array[String]]): Unit = {
    val writable_file = file.map(line => line.reduce((accum,c) => accum + ";" + c ))
    writable_file.coalesce(1).saveAsTextFile(Conf.Arguments.viewFile)
  }

  def print_changes() : Unit = {
    val pw = new PrintWriter(new File(Conf.Arguments.changesFile))
    pw.write("Hello, world")
    pw.close
  }

  def config_environment(){
    conf = new SparkConf().setAppName(Conf.AppName)
    sc = new SparkContext(conf)
    //
    evaluator = new Evaluator(sc) 

    val dir1 = new File(Conf.outputFolder)
    var dir2 = new File(Conf.inputFolder)
    dir1.mkdir()
    dir2.mkdir()
  }

  def check_args(args : Array[String]) : Unit = {
    //Init Logger : 
    PropertyConfigurator.configure("log4j.properties");
    if(args.length < 4 || args.length > 4){
      println("use : ./ws <data.csv> <user.txt> <view0.csv> <changes.txt>")
      System.exit(0)
    }

    Conf.Arguments.dataFile(args(0))
    Conf.Arguments.userFile(args(1))
    Conf.Arguments.viewFile(args(2))
    Conf.Arguments.changesFile(args(3))
    Conf.Arguments.print(println)
  }
}