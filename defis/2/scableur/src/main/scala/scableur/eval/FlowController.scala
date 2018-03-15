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


	def addNewCountFormula(f: PCountFormula) : Unit = {
		formulaList.synchronized {
			formulaList.put(f.p, f)
		}
	}
	
	def getResultOf(p:Point) : Option[Int] = {
		formulaList.get(p) match {
			case Some(f) => f.getResult()
			case None => None
		}
	}

	def printFormulaList() : Unit = {
		println("Formula list :: ")
		formulaList foreach {case (key, value) => println (key + "-->" + value)}
	}
}