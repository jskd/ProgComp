import scala.io.Source
import java.io.File
import java.io.FileWriter
import java.io.BufferedWriter

object FileTrackReader {


   def read_write(filename_source:String,filename_dest:String): Unit = {

    val file = Source.fromFile(filename_source)

    val outputFile = new File(filename_dest)

    val writer = new BufferedWriter(new FileWriter(outputFile))

    for ( (line,index) <- file.getLines().zipWithIndex){

      println(line)
      var lin = line.split(";")
      var i=lin.length

      for( a <- 0 to lin.length-1){
      writer.write("("+ (index) + ","+a+")" +" " +lin(a)+";")
      }


      writer.newLine()

    }
    writer.flush()
    writer.close()
  }

}
