package scableur

import scableur.data._
import scableur.utils._
import scableur.eval._


package object map {

	def mapPValue(str:String) : (Point,PValue) = {
		dumbMapping(str)
	}

	def dumbMapping(str:String) : (Point,PValue) = {
		try{
			val arr = str.split(' ')
			val pos = arr(0).stripPrefix("(").stripSuffix(")").split(',')
			val (i,j) = (pos(0).toInt,pos(1).toInt)
			val v = dumbStringToValue(i,j,arr(1))
			((i,j),v)
		}catch{
			case e: Exception => 
				println("BAD FORMATING")
				null
		}
		
	}

	

	def dumbStringToValue(i: Int,j: Int, str: String) : PValue = {
		try{
			if(str.contains("=#")){
				//count formula
				val values = str.stripPrefix("=#(").stripSuffix(")").split(',')
				val area = ((values(0).toInt,values(1).toInt), (values(2).toInt, values(3).toInt))
				val v = values(4).toInt
				
				val f = new PCountFormula((i,j),area, v)
				FlowController.addNewCountFormula(f)
				f
			}else{
				//Constant
				new PConstant((i,j),Some(str.trim().toInt))
			}
		}catch{
			case e:Exception => new PConstant((i,j),None)
		}
	}




	def stringToValue(i: Int,j: Int, str: String): PValue = {
		//TODO: fix :(
		if(str.contains("=#")){
	    	try{
	        	val parts= Conf.Patterns.countFormulaPattern.findAllIn(str)
	        	Conf.logger.info("::::wtf::" + parts) //WTF??
	        	val i1=parts.group(0)
	        	val j1=parts.group(1)
	        	val i2=parts.group(2)
	        	val j2=parts.group(3)
	        	val value=parts.group(4)
        
          val f = new PCountFormula((i,j),((i1.toInt,j1.toInt),(i2.toInt,j2.toInt)), value.toInt)
          /*Add to flowController*/
          FlowController.addNewCountFormula(f)
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

	def smartMapping(str: String): (Point,PValue) = {
		//TODO: fix :( 
		val parts = Conf.Patterns.pointValuePattern.findAllIn(str)
		
		try {
			val i = parts.group(0)
			val j = parts.group(1)
			val v = parts.group(2)
			println("::: result = { " + i + ", " + j + ", " + v + " }" )

			val intI = i.toInt
			val intJ = j.toInt
			((intI,intJ),stringToValue(intI , intJ, v))
		} catch {
			case e: Exception => 
				println("::::::" + e)
				null
		}
	}
}