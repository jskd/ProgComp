package prog.comp2018.scableur.utils.print

import java.io.{BufferedWriter, FileOutputStream, OutputStreamWriter}

import scala.collection.mutable.ArrayBuffer


class MatrixToCSV(private var matrix : ArrayBuffer[ArrayBuffer[Option[Int]]],
                  private var filename : String,
                  private var rows: Int,
                  private var cols:Int ) {
  val writer = new BufferedWriter(new OutputStreamWriter(new FileOutputStream(filename)))

  def echo(col : Int, row : Int) : Unit = {
    if(row != 0 && col == 0)  writer.write("\n")
    matrix(row)(col) match {
      case Some(x: Int) => writer.write(x + ";")
      case None => writer.write("P" + ";")
    }
  }

  def print() : Unit = {
    for{
      i <- 0 until rows
      j <- 0 until cols
    } echo(i,j)
  }
}
