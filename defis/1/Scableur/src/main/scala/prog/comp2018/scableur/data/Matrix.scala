package prog.comp2018.scableur.data

import scala.collection.mutable.ListBuffer

/**
  * Abstract container type
  */
trait TableType[Printable]  {
  def height : Int
  def width : Int
  def get (i: Int, j: Int) : Printable
  def set(v: Printable, i: Int, j: Int)

}

/**
  * Data Structure that containing parsed CSV file
  * @param h (height)
  * @param w (width)
  */
class Matrix(private var h : Int, private var w : Int) extends TableType[Value] {
  private val matrix = Array.ofDim[Value](h,w)
  var functionStack : ListBuffer[FunctionType] = ListBuffer[FunctionType]()
  override def height : Int = h
  override def width : Int = w
  override def get(i: Int, j: Int) : Value = matrix(i)(j)
  override def set(v: Value, i: Int, j: Int): Unit = matrix(i)(j) = v
}

class EvaluatedMatrix(private var h: Int, private var w : Int) extends TableType[Option[Int]] {
  private val matrix = Array.ofDim[Option[Int]](h,w)
  override def height : Int = h
  override def width : Int = w
  override def get(i: Int, j: Int) : Option[Int] = matrix(i)(j)
  override def set(v: Option[Int], i: Int, j: Int): Unit = matrix(i)(j) = v
}


object PrintTableType{

  def print(intro : String, outro: String, table : Matrix, printer: String => Unit ) : Unit = {
    var str = "\n" + intro + "\n"
    for{
      i <- Range(0, table.height)
      j <- Range(0,table.width)
    }{
      str += Value2String( table.get(i,j) )
      if(j < table.width) str += " |"
      if(i != 0 && j == 0) str += "\n"
    }
    str += "\n" + outro + "\n"
    printer(str)
  }

  def print(intro: String, outro: String, table: EvaluatedMatrix, printer: String => Unit) : Unit = {
    var str = "\n" + intro + "\n"
    for{
      i <- Range(0, table.height)
      j <- Range(0,table.width)
    }{
      str += Option2String( table.get(i,j) )
      if(j < table.width) str += " |"
      if(j == table.width - 1) str += "\n"
    }
    str += "\n" + outro + "\n"
    printer(str)
  }

  private def Value2String(v : Value) : String = {
      v match {
        case v : ConstantType => v.value.toString
        case f : FunctionType => f.getClass.getName
      }
  }

  private def Option2String(v : Option[Int]) : String = {
    v match {
      case Some(x : Int) => x.toString
      case None => "P"
    }
  }
}
