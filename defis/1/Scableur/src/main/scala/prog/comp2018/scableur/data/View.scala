package prog.comp2018.scableur.data

import java.util.ArrayList

class View(private var i: Int, private var j : Int, private var height : Int, private var width : Int) {
  private var matrix : ArrayList[ArrayList[Value]] = ViewBuilder.build(i,j,height,width)

  def getHeight : Int = matrix.size()
  def getWidth(j: Int) : Int = matrix.get(j).size()

  def getValue(i: Int, j: Int) : Option[Value] = {
    try{
       Some(matrix.get(i).get(j));
    }catch {
      case _ : IndexOutOfBoundsException =>  None
    }
  }

  def setValue(v : Value, i: Int, j: Int) : Unit = {
    try{
      matrix.get(i).set(j,v)
    }catch {
      case _ : IndexOutOfBoundsException =>
    }
  }
}