package prog.comp2018.scableur

import java.io.FileNotFoundException

import prog.comp2018.scableur.data.Matrix
import prog.comp2018.scableur.eval.{ArrayBufferEvaluator, Evaluator}
import prog.comp2018.scableur.utils.print.MatrixToCSV
import prog.comp2018.scableur.utils.{Conf, Debug}

import scala.collection.mutable.ArrayBuffer

object Scableur extends App {
  private var matrix : Matrix = null

  override def main(args : Array[String]): Unit = {
    if(args.length < 4 || args.length > 4){
      println("use : ./ws <data.csv> <user.txt> <view0.csv> <changes.txt>")
      System.exit(0)
    }
    load_args(args)
    load_data_csv()
    val result_data : ArrayBuffer[ArrayBuffer[Option[Int]]] = evaluate_data_csv()
    load_users_txt()
    evaluate_user_actions()
    print_data(result_data)
    print_user_actions()
  }

  def load_data_csv() : Unit = {
    try{
      matrix = new Matrix(Conf.InputData.dataFile)
      matrix.load()
    }catch {
      case _ : FileNotFoundException =>
        Debug.e("Data file Not Found")
        System.exit(0)
    }
  }

  def load_users_txt() : Unit = {
    //TODO
  }

  def evaluate_data_csv() :  ArrayBuffer[ArrayBuffer[Option[Int]]] = {
    val eval : ArrayBufferEvaluator = new ArrayBufferEvaluator(matrix.getMatrix)
    eval.eval
  }

  def evaluate_user_actions() : Unit = {
    //TODO
  }

  def print_data(data : ArrayBuffer[ArrayBuffer[Option[Int]]]) : Unit = {
    //FIXME: val printer : MatrixToCSV = new MatrixToCSV(data)
  }

  def print_user_actions() : Unit = {
    //TODO
  }

  def load_args(args : Array[String]) : Unit = {
    Conf.InputData.dataFile(args(0))
    Conf.InputData.userFile(args(1))
    Conf.InputData.viewFile(args(2))
    Conf.InputData.changesFile(args(3))
  }
}
