import prog.comp2018.scableur.data.Value

import scala.collection.mutable.ArrayBuffer

trait TableType[T]  {
  def length : Int
  def width : Int
  def get (i: Int, j: Int) : T
  def set(v: T, i: Int, j: Int)

}

class Matrix extends TableType[Value] {
  var matrix = new ArrayBuffer[ArrayBuffer[Value]]()
  override def length = matrix.length
  override def width = matrix(0).length
  override def get(i: Int, j: Int) = matrix(i)(j)
  override def set(v: Value, i: Int, j: Int): Unit = matrix(i)(j) = v
}