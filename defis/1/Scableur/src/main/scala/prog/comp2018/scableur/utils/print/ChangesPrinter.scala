package prog.comp2018.scableur.utils.print

import java.io.{File, PrintWriter}

import prog.comp2018.scableur.data.ChangesData

class ChangesPrinter(changes : ChangesData, filename : String) {
   val f = new File(filename)
  val writer = new PrintWriter(f)

  def print(): Unit = {
    for((userAction,listChages) <- changes.modificationList){
     val (coord, value) = userAction
      writer.append("after "+"\""+coord._1+" "+coord._2+" "+value.to_String+"\":\n")
      writer.flush

      
      for ((coordChg, valueChg) <- listChages){

        valueChg match {
          case Some(i: Int) => 
            println(i);
            writer.append(coordChg._1 + " " + coordChg._2 + " " + i+"\n")
          case None => 
            writer.append(coordChg._1 + " " + coordChg._2 + " " + "P\n" )
        }
        writer.flush
      }
    }
  }

}
