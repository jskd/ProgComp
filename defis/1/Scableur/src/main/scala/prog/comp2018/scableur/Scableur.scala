package prog.comp2018.scableur

import prog.comp2018.scableur.data.Matrix
import prog.comp2018.scableur.eval.Evaluator
import prog.comp2018.scableur.utils.Conf

object Scableur extends App {
  private var matrix : Matrix = null

  override def main(args : Array[String]): Unit = {
    if(args.length < 4 || args.length > 4){
      println("use : ./ws <data.csv> <user.txt> <view0.csv> <changes.txt>")
      System.exit(0)
    }
    load_args(args)
    load_data_csv()
    evaluate_data_csv()
    load_users_txt()
    evaluate_user_actions()
    print_data()
    print_user_actions()
  }

  def load_data_csv() : Unit = {
    matrix = new Matrix(Conf.InputData.dataFile)
    matrix.load()
  }

  def load_users_txt() : Unit = {
    //TODO
  }

  def evaluate_data_csv() : Unit = {
    val eval : Evaluator = new Evaluator
    eval.eval(matrix.getMatrix)
  }

  def evaluate_user_actions() : Unit = {
    //TODO
  }

  def print_data() : Unit = {
    //TODO
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
