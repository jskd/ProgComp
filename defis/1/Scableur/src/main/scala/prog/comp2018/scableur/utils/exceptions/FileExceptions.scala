package prog.comp2018.scableur.utils.exceptions

import prog.comp2018.scableur.utils.Debug

class FileException(private var filename : String) extends Exception {
  Debug.e("thrown by" +  filename);
}

class NotCSVFileException(filename: String) extends FileException(filename) {

}

class NotTXTFileException(filename: String) extends FileException(filename) {

}
