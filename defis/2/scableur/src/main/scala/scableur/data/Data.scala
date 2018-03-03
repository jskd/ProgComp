package scableur.data

package object Data {
	type Point = (Int, Int)
	type Area = ((Int,Int), (Int,Int))

	trait PValue

	case class PConstant(val p:Point, val value:Option[Int]) extends PValue {

	}

	case class PCountFormula(val p:Point, val defArea: Area, val value: Int ) extends PValue {

	}	
}
