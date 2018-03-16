package scableur

import scala.io.Source
import java.io.File
import java.io.FileWriter
import java.io.BufferedWriter

package object data {
	type Point = (Int, Int)
	type Area = ((Int,Int), (Int,Int))


	def getNumberOfCells(a: Area) : Int = {

		val x1 : Int = (a._1)._1
		val y1 : Int = (a._1)._2

		val x2 : Int = (a._2)._1
		val y2 : Int = (a._2)._2


		var ratiox : Int = x2 - x1
		var ratioy : Int = y2 - y1

		return (ratiox*ratioy)


	}

	def isPointInArea(p: Point, a:Area ) : Boolean = {

		val x = p._1
		val y = p._2

		val x1 : Int = (a._1)._1
		val y1 : Int = (a._1)._2

		val x2 : Int = (a._2)._1
		val y2 : Int = (a._2)._2



		return (  (x >= x1) && (x <= x2) && (y >= y1) && (y <= y2))


	}


	def print_iValue(x : ((Int,Int), Int)) : Unit = {
		println( "( " + x._1._1 + ", " + x._1._2 + ") " + x._2)
	}



	def addPositionsToCSV(filename_source:String,filename_dest: String): Unit = {
		try{
		    val file = Source.fromFile(filename_source)
		    val outputFile = new File(filename_dest)
		    val writer = new BufferedWriter(new FileWriter(outputFile))
		    for ( (line,index) <- file.getLines().zipWithIndex){
			
		      var lin = line.split(";")
		      var i=lin.length-1
		      for( a <- 0 to lin.length-2){
		        writer.write("("+ (index) + ","+a+")" +" " +lin(a)+";")
		      }
		      writer.write("("+ (index) + ","+i+")" +" " +lin(i))
		      writer.newLine()     

		    }

		    writer.flush()
		    writer.close()
		}catch{
			case e: java.io.FileNotFoundException => println("#file not found: " + filename_source)
		}
	} 
}