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
    // Load data in Spark
    val csvFile = sc.textFile(Conf.Arguments.dataFile)
    //Split data
    val splitedRDD = csvFile.map(line => line.split(";").map(elem => elem.trim))
    //Parse data to Objects
    val mapped_data = splitedRDD.map(line => line.map(elem => mapPValue(elem)))
    //TODO: cycle_detection()
    //Evaluate data
    //TODO: val evaluated_data = mapped_data.map(line => line.map((p,elem) => eval(elem)))

    //DEBUG PRINTS
    mapped_data.foreach(line => line.foreach(v => println(v.toString())))
    FlowController.printFormulaList()

    
    //TODO: eval

    //TODO: users

    //Print the data in result csv file
    val res_file = splitedRDD.map(line => line.reduce((accum,c) => accum + ";" + c))
    /*TODO: format output file*/
    res_file.coalesce(1).saveAsTextFile(Conf.Arguments.viewFile)

    //Print the user changes in changes file
    val pw = new PrintWriter(new File(Conf.Arguments.changesFile))
    pw.write("Hello, world")
    pw.close
    Conf.logger.info("Scableur : done.")
  }


  def parse_input() : Unit = {
    // Parse input to add position
    addPositionsToCSV(Conf.Arguments.dataFile, Conf.positionDataCSV)
    Conf.Arguments.dataFile(Conf.positionDataCSV)
  }

  def config_environment(){
    //Set up spark and other enviroment values 
    val conf = new SparkConf().setAppName(Conf.AppName)
    //conf.set("spark.scheduler.mode", "FAIR")
    sc = new SparkContext(conf)
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