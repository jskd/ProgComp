package prog.comp2018.scableur.eval

import prog.comp2018.scableur.data.functions.NbrIteration
import prog.comp2018.scableur.data.{EvaluatedMatrix, FunctionType, Matrix, Value, ConstantType}
import prog.comp2018.scableur.utils.Conf

/**
  * Evaluates a FunctionType to Option[Int]
  */
trait FunctionEval[T] {
  def eval(f : FunctionType, m : T) : Option[Int]
}


object NbrIterationEval extends FunctionEval[Matrix] {
  var counter = 0

  def eval(f: NbrIteration, m: Matrix, dependencies: List[Int] = List()): Option[Int] =  {
    dependencies.+: (m.functionStack.indexOf(f))
    for{
      i <- Range(f.from._1,f.to._1)
      j <- Range(f.from._2, f.to._2)
    }{

      (m.get(i,j)) match{

        case x:ConstantType => if(m.get(i,j) == f.value){counter += 1}

        case y:NbrIteration =>
          if(dependencies.exists(i => i == (m.functionStack.indexOf(y)))){
              m.set(ConstantType(None), y.coordinates._1, y.coordinates._2)
          }else{
            //TODO FOR ALL Functiontype THIS IS JUSTE ONE CASE ÉVALUATE "Nbriteration"
            val reslt = NbrIterationEval.eval(y,m,dependencies)
            if(ConstantType(reslt) == f.value)
              counter += 1
          }

        case _ => 

      }

      if(counter > Conf.MAX_VALUE) return None
    }
    Some(counter)
  }

  override def eval(f: FunctionType, m: Matrix): Option[Int] = None
}



