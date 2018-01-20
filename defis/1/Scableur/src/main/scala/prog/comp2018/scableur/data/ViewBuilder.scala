package prog.comp2018.scableur.data

import java.util
import java.util.ArrayList


import prog.comp2018.scableur.utils.{Conf, FileReader}

object ViewBuilder {
  // Build a Partial View from data

  def getsection(str: String, i: Int, j:Int):ArrayList[String] ={
    return new util.ArrayList[String]()//////TODO
  }

  def build(i1: Int, j1: Int, i2: Int, j2 : Int ) : ArrayList[ArrayList[String]] = {
    //TODO load partie partiel de data.csv
    val matrix : ArrayList[ArrayList[String]] = new ArrayList[ArrayList[String]]()

    var s : String =""
    
    for(i <- i1  to i2+1 ){
      s=FileReader.getNthLine(Conf.InputData.dataFile,i)
      matrix.add(getsection(s, j1,j2))

    }


    

    matrix
  }
}
