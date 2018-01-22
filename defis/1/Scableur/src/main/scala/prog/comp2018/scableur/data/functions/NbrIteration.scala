package prog.comp2018.scableur.data.functions

import prog.comp2018.scableur.data._

class NbrIteration(args: Array[String])  extends FunctionType(args) {
  
  
  def nbrIteration(r1:Int, c1:Int, r2:Int, c2:Int, v:Value, view:View) : Int =
  {
  	var count = 0
  	for(i <- r1 to r2){
  		for(j <- c1 to c2){
  			//if(v == view.getValue(i,j)) count += 1
  		}
  	}
  	count
  } 

   def eval(view: View) : Option[Int] = {
    //TODO
    try{
       var itr = nbrIteration(args(0).toInt, args(1).toInt, args(2).toInt
       				, args(3).toInt, ViewBuilder.evaluate(args(4)), view)

       Some(itr)
    }catch {
      case _ : IndexOutOfBoundsException =>  println("mutch arguments, functions expected 5 arguments.")
      None
    }

  }
}