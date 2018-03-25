package scableur.data

import java.util.concurrent.CountDownLatch
import java.util.concurrent.atomic.AtomicInteger

// On définit un trait pour nos cellules ,c'est toutes des Pvalue de 
//type PConstant(valeur) ou PCountFormula(formule)
trait PValue{
	def toString() : String
}

//Nous avons des classes Pconstant et PcountFormula (qui sont de Pvalue )
//pour représenter nos valeurs et nos formules
case class PConstant(val p:Point, val value:Option[Int]) extends PValue {

	override def toString() : String = {
		value match {
			case Some(x:Int) => "(%d,%d) %d".format(p._1, p._2, x)
			case None => "(%d,%d) %s".format(p._1, p._2, "P")
		}
	}
}

case class PCountFormula(val p:Point, val defArea: Area, val value: Int ) extends PValue {
		private var NbrPoints = 0
		private var MAX_POINTS = getNumberOfCells(defArea)
		@volatile private var result = new AtomicInteger(0)//résultat de la formule
		@volatile private var countDownLatch = new CountDownLatch(MAX_POINTS);
		@volatile private var invalid = false//flag de validité

//la méthode receive reçoie un point (ses coordonnées et sa valeur)
	//s'il n'est pas invalide et qu'il est dans sa zone de "juridication" :p
	//il regarde si sa valeur correspond à la valeur concernée par la formule 
    //(si c'est le cas on incrémente le résultat
	//de la formule)

		def receive(p: Point, v:Option[Int]) : Unit = {
			if(!(invalid) && isPointInArea(p, defArea)){
				NbrPoints += 1 //Increment number of received cells
				v match { 
					case Some(v1) => 
						if(v1 == value){
							result.incrementAndGet() 	
						}
					case None => //do nothing
				}
				countDownLatch.countDown();
			}
		}

		//met le flag de validité de la formule à true
		def setInvalid() : Unit = { invalid = true }

		//débuggage
		override def toString() : String = {
			val f = "=#(%d,%d,%d,%d,%d)".format(defArea._1._1, defArea._1._2, defArea._2._1, 
              defArea._2._2, this.value)

			"(%d,%d) %s".format(p._1,p._2,f)
		}

		//affichage du résultat s'il en est un
		def getResult() : Option[Int] = {
			if(invalid) return None
			countDownLatch.await();
			return Some(result.get())
		}

	}
