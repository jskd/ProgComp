package scableur.eval

import scableur.data.Data._
import scableur.utils._
import org.apache.spark.SparkContext



class Evaluator(sparkContext: SparkContext) {

	/**
		
	*/
	def evalPValue(s: String) : String = {
		
		//parse String to Value
		//evaluate Value
		//
		"P"
	}
}


object StringToPValue{
	def sToV(str: String) : PValue = {

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

	def stringToFormula(str: String): PFormula = {

		if(str(1)=='='){
	        val parts=Conf.Patterns.countFormulaPattern.findAllIn(str)
	        val i1=parts.group(1)
	        val j1=parts.group(2)
	        val i2=parts.group(3)
	        val j2=parts.group(4)
	        val value=parts.group(5)

        try{
          //new NbrIteration((i,j),(i1.toInt,j1.toInt),(i2.toInt,j2.toInt), ConstantType(Some(value.toInt)))
          null
        }catch {
          case _ : Exception => null
        }

      }else{
        null
      }

	}
}