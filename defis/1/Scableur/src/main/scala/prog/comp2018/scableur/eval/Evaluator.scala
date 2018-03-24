package prog.comp2018.scableur.eval

import prog.comp2018.scableur.data.functions.NbrIteration
import prog.comp2018.scableur.data.{EvaluatedMatrix, _}

import scala.collection.mutable.ListBuffer
import scala.util.control.Breaks._


abstract class Evaluator[T](private var matrix: Matrix) {

  private var cycle_functions: List[Int] = List()

  def eval : T


  /*
    Find_cycle: Calculates the list of the functions 
    of the cycle arguments:
      f: is the calling function 
      g: is the called function
      m: the matrix to evaluate
      dep: dependency list between function indexes on function_stack
   For each call to the function find_cycle we add the index of the 
   calling function to the dependency list
   */
  def find_cycle(f:FunctionType,matrix:Matrix, dep: List[Int]): Unit = {

    for{
      i <- Range(f.from._1, (f.to._1 + 1))
      j <- Range(f.from._2, (f.to._2 + 1))
    }{

      (matrix.get(i,j)) match{

        case x:ConstantType =>
        case g:FunctionType => {
          if(dep.exists(i => i==(matrix.functionStack.indexOf(Some(g))))){

            for(i<-(dep.+: (matrix.functionStack.indexOf(Some(f))))) 
              cycle_functions =
              (cycle_functions.:::(
                dep.+:(matrix.functionStack.indexOf(Some(f))))).distinct
            
           return 
          }else{ 
           find_cycle(g,matrix,(dep.+: (matrix.functionStack.indexOf(Some(f)))))
          }
        }
      }
    }
  }

   def eliminateCycle() = {
     for( x <- matrix.functionStack){
       x match {
         case Some(f1) => {
           find_cycle(f1,matrix, List())
         }
         case None =>
       }
     }

     for(i <-  cycle_functions) {
       val h = matrix.functionStack(i);
       h match{
         case Some(f1) => {
           matrix.set(ConstantType(None),
             f1.coordinates._1, f1.coordinates._2)
         }
         case None =>
       }
       matrix.functionStack.update(i,None)
     }

   }

}

 
/**
  * Evaluates a Matrix given as parameter
  */
class MatrixEvaluator(matrix : Matrix  )
  extends Evaluator[EvaluatedMatrix](matrix) {
  private val result = new EvaluatedMatrix(matrix.height,matrix.width)

  private def eval_function(i: Int, j: Int): Unit = {
    matrix.get(i,j) match {
      case f1 : NbrIteration =>
        matrix.functionStack.append(Some(f1))
        result.set(None,i,j)
      case _ => result.set(None,i,j)
    }
  }

  private def eval_cell(i: Int, j: Int) : Unit = {
    matrix.get(i,j) match {
      case v : ConstantType => result.set(v.value, i, j)
      case _ : FunctionType => eval_function(i,j)
      case _ => result.set(None,i,j)
    }
  }


  def eval_stack() = {
    for(x <- matrix.functionStack){
      x match {
        case Some(f1) => 
          f1 match{
            case f: NbrIteration => {
              result.set(NbrIterationEval.eval(f,matrix),
                                              f.coordinates._1,f.coordinates._2)
            }
            case _ =>
          }
        case None =>
      }
    }
  }
  

  def eval : EvaluatedMatrix = {
    for{
      i <- Range(0,matrix.height)
      j <- Range(0, matrix.width)
    }{
      eval_cell(i,j)
    }
    //Eliminate the cycle
    eliminateCycle()
    //Evaluate all the functions
    eval_stack()
    result
  }
}
