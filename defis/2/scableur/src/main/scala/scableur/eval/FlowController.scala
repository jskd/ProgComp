package scableur.eval

import scableur.data._
import scableur.utils._

object FlowController {
	//List of formulas 


	def receiveValueCell(position: Point, value:Option[Int]): Unit = {
		//TODO: map list of formulas to send them the value 
	}

	def receiveResult(position: Point, value:Option[Int]): Unit = {
		//TODO map list of formulas to send them the value
	}


	def addNewCountFormula(position: Point, area:Area, v:Int) : Unit = {
		/*TODO*/
	}
	
	def getResultOf(p:Point) : Option[Int] = {
		/*TODO*/
		None 
	}
}