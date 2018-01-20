package prog.comp2018.scableur.data

import java.util
import java.util.ArrayList


import prog.comp2018.scableur.utils.{Conf, FileReader}

object ViewBuilder {
  // Build a Partial View from data

  def evaluate(s:String):Value ={
    null
  }

  def getsection(str: String, i: Int, j:Int):ArrayList[Value] ={//on appel evaluate que je met dans le arrayList. pour chaque
    //case entre i et j ,et on stock la Value retournée dans le ArrayList[Value] et on le renvoi
    return new util.ArrayList[Value]()//////TODO
  }

  def build(i1: Int, j1: Int, i2: Int, j2 : Int ) : ArrayList[ArrayList[Value]] = {
    //TODO load partie partiel de data.csv
    val matrix : ArrayList[ArrayList[Value]] = new ArrayList[ArrayList[Value]]()

    var s : String =""

    for(i <- i1  to i2+1 ){
      s=FileReader.getNthLine(Conf.InputData.dataFile,i)
      matrix.add(getsection(s, j1,j2))

    }

    matrix

  }


}
