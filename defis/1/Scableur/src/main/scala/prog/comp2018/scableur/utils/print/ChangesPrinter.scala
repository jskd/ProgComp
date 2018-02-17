package prog.comp2018.scableur.utils.print

import java.io.{File, PrintWriter}

import prog.comp2018.scableur.data.ChangesData

class ChangesPrinter(changes : ChangesData, filename : String) {
   val f = new File(filename)
  val writer = new PrintWriter(f)

  def print(): Unit = {
    for((coord,item) <- changes.modificationList){
      item match {
        case Some(i: Int) => println(i);writer.append(coord._1 + " " + coord._2 + " " + i+"\n")
        case None => writer.append(coord._1 + " " + coord._2 + " " + "P\n" )
      }
      writer.flush()

    }
  }
}
