package prog.comp2018.scableur.data

import java.util.ArrayList

import prog.comp2018.scableur.file.FileReader
import prog.comp2018.scableur.utils.Conf

object ViewBuilder {
  // Build a Partial View from data
  def build(i1: Int, j1: Int, i2: Int, j2 : Int ) : ArrayList[ArrayList[Value]] = {
    //TODO load partie partiel de data.csv
    val f : FileReader = new FileReader(Conf.InputData.dataFile)
    var s : String  = f.line

    var cptline : Int =0
    

    while (s != null){
      println(s)
      s = f.line
    }
    null
  }
}
