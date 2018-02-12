package prog.comp2018.scableur.utils.print

import prog.comp2018.scableur.data.ChangesData

class ChangesPrinter(changes : ChangesData, filename : String) {

  def print(): Unit = {
    for((coord,item) <- changes.modificationList){
      //TODO print line in file according to convention
    }
  }
}
