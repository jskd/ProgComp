package prog.comp2018.scableur.eval

import prog.comp2018.scableur.data.functions.NbrIteration
import prog.comp2018.scableur.data.{EvaluatedMatrix, _}

trait Evaluator[T] {
  def eval : T
}

/**
  * Evaluates a Matrix given as parameter
  */
class MatrixEvaluator(private var matrix : Matrix  )
  extends Evaluator[EvaluatedMatrix] {
  private val result = new EvaluatedMatrix(matrix.height,matrix.width)

  private def eval_function(i: Int, j: Int): Unit = {
    matrix.get(i,j) match {
      case f1 : NbrIteration =>
        matrix.functionStack.append(f1)
        result.set(None,i,j)
      case _ => result.set(None,i,j)
    }
  }

  private def eval_cell(i: Int, j: Int) : Unit = {
    matrix.get(i,j) match {
      case v : ConstantType => result.set(v.value, i, j)
      case _ : FunctionType => eval_function(i,j)
      case _ => result.set(None,i,j)
    }
  }

  def eval_stack() = {
    //TODO read function stack and find cycles
    for(f <- matrix.functionStack){
      f match {
        case f1 : NbrIteration => result.set(NbrIterationEval.eval(f1,matrix), f.coordinates._1,f.coordinates._2)
        case _ =>
      }
    }
  }

  def eval : EvaluatedMatrix = {
    for{
      i <- Range(0,matrix.height)
      j <- Range(0, matrix.width)
    } eval_cell(i,j)
    eval_stack()
    result //return
  }
}
