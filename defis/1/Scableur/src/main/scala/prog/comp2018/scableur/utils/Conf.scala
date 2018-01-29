package prog.comp2018.scableur.utils
import prog.comp2018.scableur.utils.exceptions.{NotCSVFileException, NotTXTFileException}
import prog.comp2018.scableur.utils.RegexUtils._

object Conf {
  object Arguments {
    private var _dataCSVPath:String = ""
    private var _userTXTPath:String = ""
    private var _viewCSVPath:String = ""
    private var _changesTXTPath:String = ""

    def dataFile (filepath : String): Unit = {
     // if (! (  """\w*.csv""".r =~ filepath)) throw new NotCSVFileException(filepath)
      _dataCSVPath = filepath
    }

    def userFile (filepath : String) : Unit = {
      // if (!("""\w*.txt""".r =~ filepath )) throw new NotTXTFileException(filepath)
      _userTXTPath = filepath
    }

    def viewFile (filepath : String) : Unit = {
      // if (!("""\w*.csv""".r =~ filepath)) throw new NotCSVFileException(filepath)
      _viewCSVPath = filepath
    }

    def changesFile (filepath : String ) : Unit = {
      // if (!("""\w*.txt""".r =~ filepath)) throw new NotTXTFileException(filepath)
      _changesTXTPath = filepath
    }

    def dataFile : String = _dataCSVPath
    def userFile : String = _userTXTPath
    def viewFile : String = _viewCSVPath
    def changesFile : String = _changesTXTPath

    def print() : Unit = {
      Debug.d("data=" + _dataCSVPath)
      Debug.d("user=" + _userTXTPath)
      Debug.d("view=" + _viewCSVPath)
      Debug.d("changes=" + _changesTXTPath)
    }
  }

}

