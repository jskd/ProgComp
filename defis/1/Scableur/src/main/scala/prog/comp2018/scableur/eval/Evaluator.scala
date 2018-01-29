package prog.comp2018.scableur.eval

import prog.comp2018.scableur.data.functions.NbrIteration
import prog.comp2018.scableur.data.{ConstantType, FunctionType, Value}

import scala.collection.mutable.ArrayBuffer

trait Evaluator[T] {
  def eval : T
}

/**
  * Evaluates a Matrix given as parameter
  */
class ArrayBufferEvaluator(private var matrix : ArrayBuffer[ArrayBuffer[Value]]  )
  extends Evaluator[ArrayBuffer[ArrayBuffer[Option[Int]]]] {

  private val result = new ArrayBuffer[ArrayBuffer[Option[Int]]] //FIXME: size of array

  private def eval_function(i: Int, j: Int): Unit = {
    matrix(i)(j) match {
      case f1 : NbrIteration => result(i)(j) = None //TODO : eval(f1, matrix)
      case _ => result(i)(j) = None
    }
  }

  private def eval_cell(i: Int, j: Int) : Unit = {
    matrix(i)(j) match {
      case v : ConstantType => result(i)(j) = v.value
      case _ : FunctionType => eval_function(i,j)
      case _ => result(i)(j) = None
    }
  }

  def eval : ArrayBuffer[ArrayBuffer[Option[Int]]] = {
    for{
      i <- 0 until matrix.length
      j <- 0 until matrix(0).length
    } eval_cell(i,j)
    result //return
  }
}
