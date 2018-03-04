package scableur.utils

import org.apache.log4j.Logger

object Conf { 
	val AppName = "scableur"  
	val MAX_VALUE : Int = 255
  	val MIN_VALUE : Int = 0
  	val TempFileName = "temp.csv"
  	var logger : Logger = _ 	

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


	object Patterns {
		/* Regex pattern to match count formula in a string*/
		val countFormulaPattern = """^.*[#=(](\d+),(\d+),(\d+),(\d+),(\d+)[)].*$""".r
		/* Regex to match point Value pair in a string */ 
		val pointValuePattern = """^.*\((\d+),(\d+)\) (.+).*$""".r
	}

}