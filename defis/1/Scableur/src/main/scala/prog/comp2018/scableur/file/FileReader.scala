package prog.comp2018.scableur.file

import java.io.{FileInputStream, IOException}
import java.util.Scanner
import prog.comp2018.scableur.utils.Debug

class FileReader {
  var inputStream : FileInputStream = null
  var scanner : Scanner = null

  def FileReader (filename : String) : Unit = {
    inputStream = new FileInputStream(filename)
    scanner = new Scanner(inputStream, "UTF-8")
  }

  def line : String = {
    try{
      if(scanner.hasNextLine()){
        return scanner.nextLine()
      }
    }catch {
      case ioe: IOException =>
        Debug.e("Error in reading file")
        System.exit(0)
    }
    return ""
  }

  def close : Unit = {
    inputStream.close()
    scanner.close()

  }
}
