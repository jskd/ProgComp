package scableur

import scableur.data._
import scableur.utils._


package object map {

	def mapPValue(str: String): (Point,PValue) = {
		val parts = Conf.Patterns.pointValuePattern.findAllIn(str)
		Conf.logger.info("::::wtf::" + parts)
		try {
			val i = parts.group(1)
			val j = parts.group(2)
			val v = parts.group(3)
			Conf.logger.info("::: result = { " + i + ", " + j + ", " + v + " }" )

			val intI = i.toInt
			val intJ = j.toInt
			((intI,intJ),stringToValue(intI , intJ, v))
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
        
          val f = new PCountFormula((i,j),((i1.toInt,j1.toInt),(i2.toInt,j2.toInt)), value.toInt)
          /*Add to flowController*/
          return f
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