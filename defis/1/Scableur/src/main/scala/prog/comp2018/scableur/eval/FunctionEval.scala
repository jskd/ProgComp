package prog.comp2018.scableur.eval

import prog.comp2018.scableur.data.functions.NbrIteration
import prog.comp2018.scableur.data.{EvaluatedMatrix, FunctionType, Matrix, Value}
import prog.comp2018.scableur.utils.Conf

/**
  * Evaluates a FunctionType to Option[Int]
  */
trait FunctionEval[T] {
  def eval(f : FunctionType, m : T) : Option[Int]
}


object NbrIterationEval extends FunctionEval[Matrix] {
  var counter = 0

  def eval(f: NbrIteration, m: EvaluatedMatrix): Option[Int] =  {
    for{
      i <- Range(f.from._1,f.to._2)
      j <- Range(f.from._2, f.to._2)
    }{
      if(m.get(i,j) == f.value){
        counter += 1
      }
      if(counter > Conf.MAX_VALUE) return None
    }
    Some(counter)
  }

  override def eval(f: FunctionType, m: Matrix): Option[Int] = None
}



