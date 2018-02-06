package prog.comp2018.scableur.parsor

import com.github.tototoshi.csv.{CSVReader, DefaultCSVFormat}
import prog.comp2018.scableur.data.functions.NbrIteration
import prog.comp2018.scableur.data.{ConstantType, Matrix, Value}


class BuildMatrix(private var filename: String) {

  implicit object MyFormat extends DefaultCSVFormat {
    override val delimiter = ';'
  }

  def inInterval(x:Int, min:Int = 0, max:Int = 255):Int = {
   if(x<min) min
    else
    if(x>max) max
      else x
  }

  def load() : Matrix = {
     val reader = CSVReader.open(filename)
     val unevaluatedMatrix : List[List[String]] = reader.all()
     val matrix : Matrix = new Matrix(unevaluatedMatrix.size,unevaluatedMatrix.head.size)
     StringToValue(unevaluatedMatrix, matrix)
  }

  def StringToValue(lls : List[List[String]], matrix: Matrix) : Matrix = {
    for{
      i <- Range(0,matrix.height)
      j <- Range(0,matrix.width)
    }{
      val str = lls(i)(j)
      if(str(0) == '#'){
        matrix.set(parseFunction(str,i,j), i, j)
      }else{
        try{
          matrix.set(ConstantType(Some(inInterval(str.toInt)) ), i,j)
        }catch {
          case _ : Exception =>
            matrix.set(ConstantType(None), i,j)
        }
      }
    }
    matrix
  }

  def parseFunction(str: String, i: Int, j: Int): Value = {
    val pattern = """[#=(](\d+),(\d+),(\d+),(\d+),(\d+)[)]""".r
    if(str(1)=='='){
      val parts=pattern.findAllIn(str)
      val i1=parts.group(1)
      val j1=parts.group(2)
      val i2=parts.group(3)
      val j2=parts.group(4)
      val value=parts.group(5)
      try{
        new NbrIteration((i,j),(i1.toInt,j1.toInt),(i2.toInt,j2.toInt), ConstantType(Some(value.toInt)))
      }catch {
        case _ : Exception => ConstantType(None)
      }

    }else{
      ConstantType(None)
    }
  }
}
