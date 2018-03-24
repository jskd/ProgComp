package scableur.eval

import scableur.data._
import scableur.utils._
/**
	Object controlling the flow of data towards a list of formulas
*/
object FlowController {
	@volatile var formulaList = scala.collection.mutable.Map[Point,PCountFormula]()


	//si dans toutes les formules , le point - option(value ) en parametre est dans
	//l'une des formule de la formula list, si c'est le cas , faire receive des parametre

	def receiveValue(position: Point, value:Option[Int]): Unit = {

		for ( f <- formulaList ) {

			var actualfun : PCountFormula = f._2

			var a : Area = f._2.defArea

			if(isPointInArea(position,a)) {

				actualfun.receive(position,value)

			}

		}

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