package prog.comp2018.scableur.utils

object FileReader {
  @throws[NoSuchElementException ]("Out of Bounds")
  def getNthLine(filename : String, lineNumber : Int) : String = {
    val src = io.Source.fromFile(filename).getLines
    src drop(lineNumber) next
  }
}
