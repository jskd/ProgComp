package scableur.eval

import scableur.data._
import scableur.utils._
import org.apache.spark.SparkContext
import org.apache.spark.rdd.RDD
import org.apache.log4j.Logger




class Evaluator(sparkContext: SparkContext) {
	//val formulasRDD : RDD[List] = sc.emptyRDD[List]
	var folmulasList: List[PValue] = List[PValue]()

	def evalPValue(s: String) : String = {
		val pvalue = StringToPValue.convert(s)
		pvalue match {
			case pc : PConstant => pc.toString()
			case count : PCountFormula => 
				folmulasList = count :: folmulasList
				count.toString()
			case _ => "P"
		}
	}

	def getFormulasList() : List[PValue] = folmulasList
}


object StringToPValue{

	def convert(str: String) : PValue = {
		val parts = Conf.Patterns.pointValuePattern.findAllIn(str)
		Conf.logger.info("::::wtf::" + parts)
		try {
			val i = parts.group(1)
			val j = parts.group(2)
			val v = parts.group(3)
			Conf.logger.info("::: result = { " + i + ", " + j + ", " + v + " }" )

			stringToValue(i.toInt ,j.toInt, v)
		} catch {
			case e: Exception => 
				Conf.logger.info("::::wtf::" + e)
				null
		}
	}


	def stringToValue(i: Int,j: Int, str: String): PValue = {
		
		if(str.contains("=#")){
	    	try{
	        	val parts= Conf.Patterns.countFormulaPattern.findAllIn(str)
	        	Conf.logger.info("::::wtf::" + parts) //WTF??
	        	val i1=parts.group(1)
	        	val j1=parts.group(2)
	        	val i2=parts.group(3)
	        	val j2=parts.group(4)
	        	val value=parts.group(5)
        
          new PCountFormula((i,j),((i1.toInt,j1.toInt),(i2.toInt,j2.toInt)), value.toInt)
        }catch {
          case _ : Exception => PConstant((i,j), None)
        }

      }else{
      	try{
      		val v = str.toInt
        	PConstant((i,j), Some(v))
        }catch{
        	case _ : Exception => PConstant((i,j), None)
        }
      }
	}
}