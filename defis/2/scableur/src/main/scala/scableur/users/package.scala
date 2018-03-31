package scableur

import scableur.data.{PConstant, PCountFormula, PValue, Point}
import scableur.eval.FlowController
import scableur.eval.FlowController.formulaList

import scala.collection.mutable.ListBuffer
import scala.io.Source

package object users {

	/**

		- parsing du fichier users
		- filtrer les constantes
		- cree une liste des fonction
		- ajouter les fonctions a reevalue
		- detection des cycles

		- Evaluation des modification (liste changes)


		- à voir : 
			- pusher les fonction dans spark pour evaluation parallele
			- sectionner les donnee en block (pour garder que les info qui nous interessse)


	*/

	private def parseFile(filename : String):(scala.collection.mutable.Map[Point,PCountFormula],
																										scala.collection.mutable.Map[Point,PConstant])={


		var formulaListUser = scala.collection.mutable.Map[Point,PCountFormula]()
		var pconstantListUser = scala.collection.mutable.Map[Point,PConstant]()


		var a:Array[String]=Array()
		var i:Int=0
		var j:Int=0
		var str:String=""

		for (line <- Source.fromFile(filename).getLines) {
			a=line.split(' ')
			i=a(0).toInt
			j=a(1).toInt
			str=a(2)

			if(str.contains("=#")){

				val values = str.stripPrefix("=#(").stripSuffix(")").split(',')
				val area = ((values(0).toInt,values(1).toInt), (values(2).toInt, values(3).toInt))
				val v = values(4).toInt

				val f = new PCountFormula((i,j),area, v)
				formulaListUser.put((i,j),f)

			}else{

				val pc=new PConstant((i,j),Some(str.trim().toInt))
				pconstantListUser.put((i,j),pc)
			}

	}


	def map_user_actions() : Unit = {
		var lform=users.parseFile(filename : String)._1

		for ( i <- lform ) {

		FlowController.addNewCountFormulaUser(i._2)

		}

	}

	def evaluate_changes() : Unit = {

	}

}