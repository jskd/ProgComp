package prog.comp2018.scableur

import prog.comp2018.scableur.utils.Conf

object Scableur extends App {

  override def main(args : Array[String]): Unit = {
    if(args.length < 4 || args.length > 4){
      println("use : ./ws <data.csv> <user.txt> <view0.csv> <changes.txt>")
      System.exit(0)
    }
    load_args(args)
    Conf.InputData.print

  }

  def load_args(args : Array[String]) : Unit = {
    Conf.InputData.dataFile(args(0))
    Conf.InputData.userFile(args(1))
    Conf.InputData.viewFile(args(2))
    Conf.InputData.changesFile(args(3))
  }


  /* main method :
   *  parse arguments
   *
   *    - build view (partials ...)
   *    - run evaluator
   *    - compile result
   */

  /*
  objects:

    - evaluator : class (interface + abstract for extensibility different evaluation methods)
    - function parser ( contains algorithm to do lazy and optimised evaluation )
    - Data structure for matrix => partial view (!! dependencies)
   */
}
