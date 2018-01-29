package prog.comp2018.scableur.utils

import java.io.{FileNotFoundException}

import scala.io.Source

object FileReader {

  @throws[NoSuchElementException ]("Out of Bounds")
  def getNthLine(filename : String, lineNumber : Int) : String = {
    val src = io.Source.fromFile(filename).getLines
    src drop(lineNumber) next
  }

  @throws[FileNotFoundException]
  def getSource(filename : String ) : Source = io.Source.fromFile(filename)
}
