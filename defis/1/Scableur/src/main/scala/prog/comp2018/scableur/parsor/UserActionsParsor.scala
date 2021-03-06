package prog.comp2018.scableur.parsor

import scala.collection.mutable.ListBuffer

import prog.comp2018.scableur.data.Value

import scala.io.Source

object UserActionsParsor {
  private def parseFile(filename : String): ListBuffer[((Int,Int),String)] ={
    var l:ListBuffer[((Int,Int),String)]=ListBuffer[((Int, Int), String)]()
    var a:Array[String]=Array()
    var i:Int=0
    var j:Int=0
    var s:String=""

    for (line <- Source.fromFile(filename).getLines) {
      a=line.split(' ')
      i=a(0).toInt
      j=a(1).toInt
      s=a(2)
      l=l:+((i,j),s)
    }
    l
  }

  def buildList(filename : String) : ListBuffer[((Int,Int),Value)] = {
    var l: ListBuffer[((Int, Int), String)] = ListBuffer[((Int, Int), String)]()
    var lv: ListBuffer[((Int, Int), Value)] = ListBuffer[((Int, Int), Value)]()

    l = parseFile(filename)

    for (k <- 0 to l.length - 1 ) {

      var v: Value = StringToValue.stv(l(k)._2, l(k)._1._1, l(k)._1._2)
      var i = l(k)._1._1
      var j = l(k)._1._2
      lv = lv :+ ((i, j), v)


    }

    return lv
  }


}
