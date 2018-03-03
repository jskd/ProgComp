package scableur.eval

import scableur.data.Data._
import scableur.utils._
import org.apache.spark.SparkContext


class Evaluator(sparkContext: SparkContext) {
	//val formulasRDD = new sparkContext.ListRDD()

	def evalPValue(s: String) : String = {
		//parse String to Value
		//evaluate Value
		// match
			// Number => Number
			// Formula => addToFormulaList
		"P"
	}

	//def getFormulas() : = formulasRDD
}


object StringToPValue{
	def convert(str: String) : PValue = {

		val parts = Conf.Patterns.pointValuePattern.findAllIn(str)

		try {
			val i = parts.group(1).toInt
			val j = parts.group(2).toInt
			//val v = stringToValue( parts.group(3) )
			null
		} catch {
			case _: IllegalStateException => null
			case _: Exception => null
		}
	}

	def stringToFormula(i: Int,j: Int, str: String): PValue = {
		if(str(1)=='=' && str(2)=='#'){
	    	
	    	try{
	        	val parts=Conf.Patterns.countFormulaPattern.findAllIn(str)
	        	val i1=parts.group(1)
	        	val j1=parts.group(2)
	        	val i2=parts.group(3)
	        	val j2=parts.group(4)
	        	val value=parts.group(5)
        
          new PCountFormula((i,j),((i1.toInt,j1.toInt),(i2.toInt,j2.toInt)), value.toInt)
          null
        }catch {
          case _ : Exception => PConstant((i,j), None)
        }

      }else{
        PConstant((i,j), None)
      }

	}
}