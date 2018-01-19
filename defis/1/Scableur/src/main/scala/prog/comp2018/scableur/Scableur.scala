package prog.comp2018.scableur

object Scableur extends App {
  override def main(args : Array[String]): Unit = {
    if(args.length < 4 || args.length > 4){
      println("use : ./ws <data.csv> <user.txt> <view0.csv> <changes.txt>")
      System.exit(1)
    }else{
      println(args)
    }
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
