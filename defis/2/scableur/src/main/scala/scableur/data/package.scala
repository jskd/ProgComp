package scableur

import scala.io.Source
import java.io.File
import java.io.FileWriter
import java.io.BufferedWriter

package object data {
	type Point = (Int, Int)
	type Area = ((Int,Int), (Int,Int))


	def getNumberOfCells(a: Area) : Int = {
		//TODO hocine
		0
	}

	def isPointInArea(p: Point, a:Area ) : Boolean = {
		//TODO hocine
		false
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