package scableur

package object data {
	type Point = (Int, Int)
	type Area = ((Int,Int), (Int,Int))


	def getNumberOfCells(a: Area) : Int = {
		//TODO
		0
	}

	def isPointInArea(p: Point, a:Area ) : Boolean = {
		false
	}


	  def print_iValue(x : ((Int,Int), Int)) : Unit = {
    	println( "( " + x._1._1 + ", " + x._1._2 + ") " + x._2)
	  } 


}