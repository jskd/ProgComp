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

object Scableur {
  // Spark Init
  val conf = new SparkConf().setAppName(Conf.AppName)
  val sc = new SparkContext(conf)
  //
  val evaluator = new Evaluator(sc) 

  
  def main(args: Array[String]): Unit = {
    Conf.logger = Logger.getLogger("DEBUG");
    Conf.logger.info("Launching Scableur")
    //check_args(args)
    //parse_input() 
    val file = spark_data_load(args(0))
    map_data(file)

    //cycle_detection()
    //evaluation_flow()
    //print_to_file()

    //TODO: users
    
    //Debug
    //file.foreach(line => line.foreach(println _))
    Conf.logger.info("Scableur : done.")
  }

  def spark_data_load(filename: String) = {
    // Load data in spark file system
    val csvFile = sc.textFile(filename)
    csvFile.map(line => line.split(";").map(elem => elem.trim))
  }


  
  def map_data(file : RDD[Array[String]]) : Unit = {
    file.foreach(line => line.foreach(println _))
    val mapped_data = file.map(line => line.map(elem => evaluator.evalPValue(elem) ))
    mapped_data.foreach(line => line.foreach(println _)) 
    val functionList = evaluator.getFormulasList()
    println("functionList : " + functionList)
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
    //Conf.Arguments.print(Conf.DEBUG)
  }
}