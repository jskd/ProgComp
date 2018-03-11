package scableur.eval

import scableur.data._
import scableur.utils._

object FlowController {
	@volatile var formulaList = scala.collection.mutable.Map[Point,PCountFormula]()


	def receiveValueCell(position: Point, value:Option[Int]): Unit = {
		//TODO: map list of formulas to send them the value 
	}

	def receiveResult(position: Point, value:Option[Int]): Unit = {
		//TODO map list of formulas to send them the value
	}


	def addNewCountFormula(position: Point, f: PCountFormula) : Unit = {
		formulaList.synchronized {
			formulaList.put(position, f)
		}
	}
	
	def getResultOf(p:Point) : Option[Int] = {
		formulaList.get(p) match {
			case Some(f) => f.getResult()
			case None => None
		}
	}
}