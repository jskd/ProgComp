package prog.comp2018.scableur

import java.io.FileNotFoundException

import prog.comp2018.scableur.data.{EvaluatedMatrix, Matrix, PrintTableType}
import prog.comp2018.scableur.eval.MatrixEvaluator
import prog.comp2018.scableur.parsor.BuildMatrix
import prog.comp2018.scableur.utils.print.MatrixToCSV
import prog.comp2018.scableur.utils.{Conf, Debug}


object Scableur {
  var matrix : Matrix = _
  var resultMatrix : EvaluatedMatrix = _

  def main(args : Array[String]): Unit = {
    check_args(args)
    load_data_csv()
    evaluate_data_csv()
    load_users_txt()
    evaluate_user_actions()
    print_data()
    print_user_actions()
  }

  def load_data_csv() : Unit = {
    try{
      Debug.i("Loading data file...")
      val builder = new BuildMatrix(Conf.Arguments.dataFile)
      matrix = builder.load()
      //PrintTableType.print("=== CSV Parsed File === ","===================", matrix, Debug.d(_))
    }catch {
      case _ : FileNotFoundException =>
        Debug.e("Data file Not Found")
        System.exit(0)
    }
  }

  def evaluate_data_csv() : Unit  = {
    Debug.i("Evaluating data file...")
    val evaluator = new MatrixEvaluator(matrix)
    resultMatrix = evaluator.eval
    PrintTableType.print("=== Evaluated Matrix === ","===================",resultMatrix,Debug.d(_))
  }

  def print_data() : Unit = {
    val printer = new MatrixToCSV(resultMatrix,Conf.Arguments.viewFile)
    printer.print()
  }

  def load_users_txt() : Unit = {
    //TODO
  }

  def evaluate_user_actions() : Unit = {
    //TODO
  }

  def print_user_actions() : Unit = {
    //TODO
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

    Conf.Arguments.print(Debug.i)
  }
}
