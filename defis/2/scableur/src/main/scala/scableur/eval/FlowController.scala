package scableur.eval

import scableur.data._
import scableur.utils._
/**
	Object controlling the flow of data towards a list of formulas
*/
object FlowController {
	// listeMap de TOUTES les formules du CSV mappée selon les coordonnées des formules
	@volatile var formulaList = scala.collection.mutable.Map[Point,PCountFormula]()
	@volatile var formulaListUser = scala.collection.mutable.Map[Point,PCountFormula]()


	//Si ,pour toutes les formules , le point donné en parametre est dans
	//l'une des formule de la formulalist,on fait vérifier ce point par receive

	def receiveValue(position: Point, value:Option[Int]): Unit = {

		for ( f <- formulaList ) {

			var actualfun : PCountFormula = f._2

			var a : Area = f._2.defArea

			if(isPointInArea(position,a)) {

				actualfun.receive(position,value)

			}

		}

	}

	//ajoute une nouvelle formule dans la liste des formules
	def addNewCountFormula(f: PCountFormula) : Unit = {
		formulaList.synchronized {
			formulaList-=(f.p)
			formulaList.put(f.p, f)
		}
	}

	//Enleve l'ancienne formule s'il y en avait une à ces coordonnées
	def addNewCountFormulaUser(f: PCountFormula) : Unit = {
		formulaList.synchronized {
			formulaList.remove(f.p)
			formulaList.put(f.p, f)
		}
	}


	//renvoi le résultat d'une formule selon ou elle se situe dans le fichier
	def getResultOf(p:Point) : Option[Int] = {
		formulaList.get(p) match {
			case Some(f) => f.getResult()
			case None => None
		}
	}

	//fonction débuggage
	def printFormulaList() : Unit = {
		println("Formula list :: ")
		formulaList foreach {case (key, value) => println (key + "-->" + value)}
	}
}
