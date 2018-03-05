

import scala.io.Source
import java.io.File
import java.io.FileWriter
import java.io.BufferedWriter

object FileReader {


  def read_write(filename_source:String,filename_dest): Unit = {
    val file = Source.fromFile(filename_source)
    val outputFile = new File(filename_dest)
    val writer = new BufferedWriter(new FileWriter(outputFile))
    for ( (line,index) <- file.getLines().zipWithIndex){
	
      var lin = line.split(";")
      var i=lin.length-1
      for( a <- 0 to lin.length-2){
      writer.write("("+ (index) + ","+a+")" +" " +lin(a)+";")
      }
      writer.write("("+ (index) + ","+i+")" +" " +lin(i))
      writer.newLine()     

    }

    writer.flush()
    writer.close()
  }



}
