package prog.comp2018.scableur.eval

import prog.comp2018.scableur.data._
import prog.comp2018.scableur.data.functions.NbrIteration

import scala.collection.mutable.ListBuffer

class UserActionEvaluator(val matrix: Matrix, val source : UserData, var resultMatrix: EvaluatedMatrix)
    extends Evaluator[ChangesData] {

  private val changesData : ChangesData = new ChangesData()

  def isInRange(pos : (Int,Int), x1 : (Int,Int), x2: (Int, Int)) : Boolean = {
    pos._1 >= x1._1 && pos._1 <= x2._1 && pos._2 >= x1._2 && pos._2 <= x2._2
  }

  def testFunctionChanges(coord: (Int,Int)) : ListBuffer[(Option[FunctionType])] = {
    var functionStack: ListBuffer[(Option[FunctionType])] = ListBuffer[(Option[FunctionType])]()
    var lstfun : List[Int] = List[Int]()

    for(i <- Range(0,matrix.functionStack.length,1)) {
      val fun = matrix.functionStack(i)
      fun match{
        case Some(f) =>{
          if(isInRange(coord, f.from, f.to) && !lstfun.contains(i) ){
            functionStack.append(fun)
            lstfun = i :: lstfun
          }
        }
        case None => 
      }
    }
    functionStack
  }

  def constantEval(v: ConstantType) : Option[Int] = v.value

  def evalFunctionStack(functionStack: ListBuffer[(Option[FunctionType])])
                                   : ListBuffer[((Int,Int), Option[Int])] = {

    var changes: ListBuffer[((Int,Int), Option[Int])] =  ListBuffer[((Int,Int), Option[Int])]()
  
    for(x <- functionStack){
      x match  {
        case Some(f) => {
          f match{
            case f1 : NbrIteration => 
              changes.+=:(f1.coordinates, NbrIterationEval.eval(f1,matrix))
            case _ =>
          }
        }
        case None =>
      }
    }

    changes.reverse
  }
  

  override def eval: ChangesData = {
    for((coord,v) <- source.userActionStack){
      
      var functionStack = ListBuffer[(Option[FunctionType])]()
      var actionChanges = ListBuffer[((Int,Int), Option[Int])]()
      
      matrix.set(v, coord._1, coord._2)
      
      v match {
        case v : ConstantType => {
          actionChanges.+=:( (coord, constantEval(v)) )
        }
        case f : FunctionType => {
          functionStack.+=:(Some(f))
        }
      }
      functionStack = (testFunctionChanges(coord)).++=:(functionStack)
      actionChanges = (evalFunctionStack(functionStack)).++:(actionChanges)
      changesData.modificationList.+=:((coord,v), actionChanges)
    }
    changesData
  }
}
