package prog.comp2018.scableur.parsor

import prog.comp2018.scableur.data.Value

import scala.io.Source

class UserActionsParsor {

  private def parseFile(filename : String): List[((Int,Int),String)] ={
    var l:List[((Int,Int),String)]=List()
    var a:Array[String]=Array()
    var i:Int=0
    var j:Int=0
    var s:String=""
    val filename = "test.txt"

    for (line <- Source.fromFile(filename).getLines) {
      a=line.split(' ')
      i=a(0).toInt
      j=a(1).toInt
      s=a(2)
      l=l:+((i,j),s)
    }
    l
  }

  def buildList(filename : String) : List[((Int,Int),Value)] = {
    null
  }

}
