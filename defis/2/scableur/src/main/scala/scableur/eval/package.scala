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

	// Evaluate a Count Formula
	def evalFCount(formula: PCountFormula) : Option[Int] = {
		/* Future & Promise do not work with spark ??
		val f: Future[Option[Int]] = Future {
  			
		}
		val res = f onComplete {
		  case Success(result) => {
		  	//FlowController.sendFunctionResult(....)
		  	result
		  }
		  case Failure(error) => {
		  	//FlowController.sendFunctionResult(....)
		  	None
		  }
		}
		Await.result(f, Duration.Inf)
		*/
		return FlowController.getResultOf(formula.p)
	}

	def evalConstant(c: PConstant) : Option[Int] = {
		None
	}

	def eval(pValue: PValue) : Option[Int] = {
		/*
		pValue match {
			...
		}
		*/
		None
	}
}