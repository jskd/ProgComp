package prog.comp2018.scableur.eval

import prog.comp2018.scableur.data.functions.NbrIteration
import prog.comp2018.scableur.data.{EvaluatedMatrix, FunctionType, Matrix, Value, ConstantType}
import prog.comp2018.scableur.utils.Conf

/**
  * Evaluates a FunctionType to Option[Int]
  */
trait FunctionEval[T] {
  def eval(f: FunctionType, m: Matrix): Option[Int]
}

object FindCycle {

  /*Return "true" if cycle is detected, otherwise false
   arguments:
    f: is the calling function 
    g: is the called function
    m: the matrix to evaluate
    dep: dependency list between function indexes on function_stack
   for each call to the function find_cycle we add the index of the calling function
   to the dependency list*/

  def find_cycle(f:FunctionType, g:FunctionType, m:Matrix, dep: List[Int]): Boolean = {
    if(dep.exists(i => i==(m.functionStack.indexOf(g)))){
      m.set(ConstantType(None), f.coordinates._1, f.coordinates._2)
      println("\nWARNING: Cycle detected between positions\n"+
              "("+f.coordinates._1+","+f.coordinates._2+") and "+
              "("+g.coordinates._1+","+g.coordinates._2+")\n")
      return true
    }else{return false}
  }
}

object NbrIterationEval extends FunctionEval[Matrix] {
 
  def eval(f: NbrIteration, m: Matrix, dependencies: List[Int] = List()): Option[Int] = {
    var counter = 0
    for{
      i <- Range(f.from._1,(f.to._1 + 1))
      j <- Range(f.from._2, (f.to._2 + 1))
    }{

      (m.get(i,j)) match{

        case x:ConstantType => if(x == f.value){counter += 1}

        case g:NbrIteration =>
          try{//TEST IF CYCLE EXIST RETURN "None" ELSE EVALUATE g 
            if(FindCycle.find_cycle(f,g,m,dependencies) ){
              return None
            }else{             
              val reslt = eval(g,m,( dependencies.+: (m.functionStack.indexOf(f))))
              if(ConstantType(reslt) == f.value)
                counter += 1
            }
          }catch{
            case _ : Exception =>
              return None
            }
         case _ => 

       }
       if(counter > Conf.MAX_VALUE) return None
     }
     Some(counter)
   }

  override def eval(f: FunctionType, m: Matrix): Option[Int] = None
}



