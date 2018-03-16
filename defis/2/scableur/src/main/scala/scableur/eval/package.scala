package scableur

import scableur.data._
import scableur.utils._
import org.apache.spark.SparkContext
import org.apache.spark.rdd.RDD
import org.apache.log4j.Logger
import scala.concurrent._
import scala.concurrent.duration._
import scala.util.{Success, Failure}


package object eval {

	// Evaluate a PValue
	def eval(pValue: PValue) : Option[Int] = {
		pValue match {
			case c:PConstant => 
				FlowController.receiveValue(c.p, c.value)
				c.value
			case f:PCountFormula => 
				val res = f.getResult()
				FlowController.receiveValue(f.p,res)
				res
		}
		None
	}
}