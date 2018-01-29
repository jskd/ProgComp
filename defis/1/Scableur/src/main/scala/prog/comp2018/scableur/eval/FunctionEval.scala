package prog.comp2018.scableur.eval

import prog.comp2018.scableur.data.{FunctionType, BuildMatrix, Value}
import scala.collection.mutable.ArrayBuffer


/**
  * Evaluates a FunctionType to Option[Int]
  */
trait FunctionEval[T] {
  def eval(f : FunctionType, m : T) : Option[Int]
}


object NbrIterationEval extends FunctionEval[BuildMatrix] {
  var counter = 0

  def count_iterations(v : Value, cmp: Value) : Unit = if( v == cmp) counter = counter + 1

  override def eval(f: FunctionType, m: BuildMatrix): Option[Int] =  {
    val table : ArrayBuffer[ArrayBuffer[Value]] = m.getMatrix
    for{
      i <- Range(f.coord(0), f.coord(1),1)
      j <- Range(f.coord(2),f.coord(3),1)
    }count_iterations(table(i)(j),f.v)
    Some(counter)
  }
}



