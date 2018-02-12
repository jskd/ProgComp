package prog.comp2018.scableur.utils.print

import java.io.{File, PrintWriter}

import prog.comp2018.scableur.data.ChangesData

class ChangesPrinter(changes : ChangesData, filename : String) {
  val writer = new PrintWriter(new File(filename))

  def print(): Unit = {
    for((coord,item) <- changes.modificationList){
      item match {
        case i : Some[Int] => writer.println(coord._1 + " " + coord._2 + " " + i )
        case None => writer.println(coord._1 + " " + coord._2 + " " + "P" )
      }

    }
  }
}
