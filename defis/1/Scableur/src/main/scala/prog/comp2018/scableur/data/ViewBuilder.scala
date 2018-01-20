package prog.comp2018.scableur.data

import java.util.ArrayList

import prog.comp2018.scableur.file.FileReader
import prog.comp2018.scableur.utils.Conf

object ViewBuilder {
  // Build a Partial View from data
  def build(i: Int, j: Int, height: Int, width : Int ) : ArrayList[ArrayList[Value]] = {
    //TODO load partie partiel de data.csv
    val f : FileReader = new FileReader(Conf.InputData.dataFile)
    var s : String  = f.line
    while (s != null){
      println(s)
      s = f.line
    }
    null
  }
}
