package scableur

import scableur.data.{PConstant, PCountFormula, PValue, Point}
import scableur.eval.FlowController
import scableur.eval.FlowController.formulaList
import scableur.data._

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

	private def parseFile(filename : String):scala.collection.mutable.Map[Point,PValue]={

		// var formulaListUser = scala.collection.mutable.Map[Point,PCountFormula]()
		// var pconstantListUser = scala.collection.mutable.Map[Point,PConstant]()

      var userActionStack = scala.collection.mutable.Map[Point,PValue]()
      
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
		  val v = values(4).toInt //Hocine ici c'est pas toujour un int ça peut etre une formule

		  val f = new PCountFormula((i,j),area, v)
		  userActionStack.put((i,j),f)

		}else{

		  val pc=new PConstant((i,j),Some(str.trim().toInt))
		  userActionStack.put((i,j),pc)
		}

	  }

      return userActionStack
    }

    def constantEval(v: PConstant) : Option[Int] = v.value

    def testFunctionChanges(p: Point): scala.collection.mutable.Map[Point,PCountFormula]={
      var functionStack = scala.collection.mutable.Map[Point,PCountFormula]()
      for((position,f) <- FlowController.formulaList){
        if(isPointInArea(p, f.defArea)){
          functionStack.put(position, f)
        }
      }

      return functionStack
    }

	def map_user_actions(data: Map[Point,PValue] , filename: String) : 
         ListBuffer[((Point,PValue),ListBuffer[(Point,Option[Int])])] = {

	  var lform= parseFile(filename)
      var changesData= ListBuffer[( (Point,PValue),ListBuffer[(Point,Option[Int])])]()

	  for ( (p,a) <- lform ) {
        var actionChanges = ListBuffer[(Point,Option[Int])]()
        a match {
          case v: PConstant =>{
            actionChanges.+=:((p,constantEval(v)))
          }
          case f: PCountFormula => {
            FlowController.addNewCountFormula(f)
            FlowController.addNewCountFormulaUser(f)
          }
        }

        cycles.setFormulaList(FlowController.formulaList)
        FlowController.formulaList = cycles.cycle_removal()
        FlowController.formulaListUser = FlowController.formulaListUser.++:(testFunctionChanges(p))
        //TODO: Eval formulaListUser and make result in actionchanges

        changesData.+=:((p,a),actionChanges)
	  }

      return changesData
	}

	def evaluate_changes() : Unit = {

	}

}
