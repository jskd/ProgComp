package prog.comp2018.scableur.data

import com.github.tototoshi.csv.CSVReader
import prog.comp2018.scableur.data.functions.NbrIteration

import scala.util.matching.Regex
import scala.collection.mutable.ArrayBuffer

class Matrix(private var filename: String) {

  var matrix : ArrayBuffer[ArrayBuffer[Value]] = new ArrayBuffer[ArrayBuffer[Value]]()

  def load() : Unit = {
     val reader = CSVReader.open(filename)
     val unevaluatedMatrix : List[List[String]] = reader.all()
     StringToValue(unevaluatedMatrix)

    //TODO read list and create values in matrix
  }

  /*

  val pattern = """[#=(](\d+),(\d+),(\d+),(\d+),(\d+)[)]""".r

  val string = "#=(1,2,3,4,5)"

  pattern.findAllIn(string).matchData foreach {
     |    m => println(m.group(1)) ; println(m.group(2)) ; println(m.group(3));
     | println(m.group(4)); println(m.group(5)) }



   */

  def StringToValue(lls : List[List[String]]) : ArrayBuffer[ArrayBuffer[Value]] = {

    val pattern = """[#=(](\d+),(\d+),(\d+),(\d+),(\d+)[)]""".r

    var str= ""

    for(i<-0 to (lls.length)){
      for(j<-0 to (lls(0).length)){

        if(lls(i)(j)(0)=='#'){
          str=lls(i)(j)
          val parts=pattern.findAllIn(str)
          val i1=parts.group(1)
          val j1=parts.group(2)
          val i2=parts.group(3)
          val j2=parts.group(4)
          val value=parts.group(5)

          matrix(i)(j)=new NbrIteration(Array(i1.toInt,j1.toInt,i2.toInt,j2.toInt),new ConstantType(Some(value.toInt)))

        }

        else{
          matrix(i)(j)=new ConstantType(Some(lls(i)(j).toInt))

        }

      }


    }

    return matrix

  }

  def getMatrix : ArrayBuffer[ArrayBuffer[Value]] = matrix
}
