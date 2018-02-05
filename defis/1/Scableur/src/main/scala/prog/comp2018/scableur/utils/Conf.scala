package prog.comp2018.scableur.utils

object Conf {
  val MAX_VALUE : Int = 255

  object Arguments {
    private var _dataCSVPath:String = ""
    private var _userTXTPath:String = ""
    private var _viewCSVPath:String = ""
    private var _changesTXTPath:String = ""

    def dataFile (filepath : String): Unit = {
      _dataCSVPath = filepath
    }

    def userFile (filepath : String) : Unit = {
      _userTXTPath = filepath
    }

    def viewFile (filepath : String) : Unit = {
      _viewCSVPath = filepath
    }

    def changesFile (filepath : String ) : Unit = {
      _changesTXTPath = filepath
    }

    def dataFile : String = _dataCSVPath
    def userFile : String = _userTXTPath
    def viewFile : String = _viewCSVPath
    def changesFile : String = _changesTXTPath

    def print(printer: String => Unit) : Unit = {
      var str = "===== Arguments ===== \n"
      str += "data=" + _dataCSVPath + "\n"
      str += "user=" + _userTXTPath + "\n"
      str += "view=" + _viewCSVPath + "\n"
      str += "changes=" + _changesTXTPath + "\n"
      str += "==========================\n"
      printer(str)
    }
  }
}

