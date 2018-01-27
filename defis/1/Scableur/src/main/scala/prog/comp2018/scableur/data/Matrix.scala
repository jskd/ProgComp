package prog.comp2018.scableur.data

import com.github.tototoshi.csv.CSVReader
import com.sun.crypto.provider.AESCipher.AES128_CBC_NoPadding

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


    return matrix

  }

  def getMatrix : ArrayBuffer[ArrayBuffer[Value]] = matrix
}
