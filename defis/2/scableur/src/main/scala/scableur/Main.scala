package scableur

import org.apache.log4j.Logger;
import org.apache.log4j.PropertyConfigurator;
import org.apache.spark.{SparkConf, SparkContext}
import org.apache.spark.rdd.RDD
import org.apache.spark.sql.SQLContext
import org.apache.spark.sql.types.{StructField, StructType, IntegerType, StringType}
import scableur.utils.Conf
import scableur.data._
import scableur.data.PValue
import scableur.eval._
import java.io._
import scableur.map._

object Scableur {
  type Data = RDD[Array[String]]
  type MapedData = RDD[Array[(Point,PValue)]]
  type ResultData = RDD[Array[String]]
  var sc : SparkContext = null

  
  def main(args: Array[String]): Unit = {
    Conf.logger = Logger.getLogger("DEBUG");
    Conf.logger.info("Launching Scableur")
    check_args(args)
    config_environment()
    parse_input() 
    val file = spark_data_load(Conf.Arguments.dataFile)
    val mapped_data = map_data(file)

    //TODO: cycle_detection()
    val result_data = eval_mapped_data(mapped_data)

    print_view_csv(file)
    print_changes()

    //TODO: users
    
    Conf.logger.info("Scableur : done.")
  }


  def parse_input() : Unit = {
    // Parse input to add position
    addPositionsToCSV(Conf.Arguments.dataFile, Conf.positionDataCSV)
    Conf.Arguments.dataFile(Conf.positionDataCSV)
  }


  def spark_data_load(filename: String) : Data = {
    // Load data in spark file system
    val csvFile = sc.textFile(filename)
    csvFile.map(line => line.split(";").map(elem => elem.trim))
  }


  def map_data(file : Data) : MapedData = {
    //Map all the data to create formula list
    val mapped_data = file.map(line => line.map(elem => mapPValue(elem)))
    /*
    println("mapped data : " + mapped_data)
    val functionList = evaluator.getFormulasList()
    println("functionList : " + functionList)
    */
    return mapped_data
  }

  def print_view_csv(file : ResultData): Unit = {
    //Print the data in result csv file
    val writable_file = file.map(line => line.reduce((accum,c) => accum + ";" + c))
    /*TODO: format output file*/
    writable_file.coalesce(1).saveAsTextFile(Conf.Arguments.viewFile)
  }

  def print_changes() : Unit = {
    //Print the user changes in changes file
    val pw = new PrintWriter(new File(Conf.Arguments.changesFile))
    pw.write("Hello, world")
    pw.close
  }


  def eval_mapped_data(data: MapedData) : ResultData = {
    //Evaluate all the cells and the formulas
    null
  }

  def config_environment(){
    //Set up spark and other enviroment values 
    val conf = new SparkConf().setAppName(Conf.AppName)
    conf.set("spark.scheduler.mode", "FAIR")
    sc = new SparkContext(conf)
    
    val dir1 = new File(Conf.outputFolder)
    var dir2 = new File(Conf.inputFolder)
    dir1.mkdir()
    dir2.mkdir()
  }

  def check_args(args : Array[String]) : Unit = {
    //Init Logger : 
    PropertyConfigurator.configure("log4j.properties");
    //Verify the number of arguments
    if(args.length < 4 || args.length > 4){
      println("use : ./ws <data.csv> <user.txt> <view0.csv> <changes.txt>")
      System.exit(0)
    }
    //Set up argument variables in Conf
    Conf.Arguments.dataFile(args(0))
    Conf.Arguments.userFile(args(1))
    Conf.Arguments.viewFile(args(2))
    Conf.Arguments.changesFile(args(3))
    Conf.Arguments.print(println)
  }
}