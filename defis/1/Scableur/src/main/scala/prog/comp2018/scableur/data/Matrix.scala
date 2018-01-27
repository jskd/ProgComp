package prog.comp2018.scableur.data

import com.github.tototoshi.csv.CSVReader
import scala.collection.mutable.ArrayBuffer

class Matrix(private var filename: String) {
  var matrix : ArrayBuffer[ArrayBuffer[Value]] = new ArrayBuffer[ArrayBuffer[Value]]()

  def load() : Unit = {
     val reader = CSVReader.open(filename)
     val unevaluatedMatrix : List[List[String]] = reader.all()

    //TODO read list and create values in matrix
  }

  def getMatrix : ArrayBuffer[ArrayBuffer[Value]] = matrix
}
