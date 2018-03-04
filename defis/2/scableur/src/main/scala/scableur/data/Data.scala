package scableur.data

trait PValue{
	def toString() : String
}

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
		private var result = 0


		def receice(cell: PConstant) : Unit = {
			if(isPointInArea(cell.p, defArea)){
				NbrPoints += 1 //Increment number of received cells
				cell.value match { 
					case Some(v) => if(v == value) result += 1
					case None => //do nothing
				}
				if(NbrPoints >= MAX_POINTS){
					//TODO: remove_function_from_list
					//TODO: return result to controler
					println("FCount : " + p + " done.")
				}
			}
		}


		override def toString() : String = {
			val f = "=#(%d,%d,%d,%d,%d)".format(defArea._1._1, defArea._1._2, defArea._2._1, defArea._2._2, this.value)
			"(%d,%d) %s".format(p._1,p._2,f)
		}

	}
