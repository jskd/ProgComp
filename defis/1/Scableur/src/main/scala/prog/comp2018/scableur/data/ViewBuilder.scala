package prog.comp2018.scableur.data

import java.util
import java.util.ArrayList


import prog.comp2018.scableur.utils.{Conf, FileReader}
import prog.comp2018.scableur.data.functions._

object ViewBuilder {
  // Build a Partial View from data

  def evaluate(s:String):Value ={
    //case value is a function
    if(s.substring(0,2).equals("=#")){
      val subS = s.substring(s.indexOf("(")+1,s.length-1)
      val args = subS.split(",")
      val f = new NbrIteration(args); f
    }
    else{
      //case value is constante or P
      val i = if(s.equals("P")) None 
      else {val ii = s.toInt; 
            if(ii >= 0 && ii <= 255) Some(ii) else None }

      val c = new ConstantType(i); c

    }
  }

  def getsection(str: String, i: Int, j:Int):ArrayList[Value] ={      // 1;2;3;4;5;6
    //on appel evaluate que je met dans le arrayList. pour chaque
    //case entre i et j ,et on stock la Value retournée dans le ArrayList[Value] et on le renvoi
    var al = new util.ArrayList[Value]()

    var arr = str.split(";")

    for(z<-i to j+1){
      var value=evaluate(arr(z))
      al.add(value)

    }



    return al
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
