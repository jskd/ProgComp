package prog.comp2018.scableur.utils.print

import java.io._

import prog.comp2018.scableur.data.EvaluatedMatrix

class MatrixToCSV(private var matrix : EvaluatedMatrix,
                  private var filename : String) {
  val writer = new PrintWriter(new File(filename))

  def print() : Unit = {
    for{
      i <- Range(0,matrix.height)
      j <- Range(0,matrix.width)
    }{
      if(i != 0 && j == 0)  writer.append("\n")

      matrix.get(i,j) match {
        case Some(x: Int) => writer.append(x.toString)
        case None => writer.append("P")
      }
      if(j < matrix.width-1) writer.append(";")
      writer.flush()
    }
    writer.close()
  }
}
