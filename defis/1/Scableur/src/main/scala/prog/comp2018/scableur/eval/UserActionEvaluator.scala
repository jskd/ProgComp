package prog.comp2018.scableur.eval

import prog.comp2018.scableur.data._

import scala.collection.mutable.ListBuffer

class UserActionEvaluator(val matrix: Matrix, val source : UserData) extends Evaluator[ChangesData] {
  private val changesData : ChangesData = new ChangesData()
  var functionStack : ListBuffer[FunctionType] = ListBuffer[FunctionType]()

  def isInRange(pos : (Int,Int), x1 : (Int,Int), x2: (Int, Int)) : Boolean = {
    pos._1 >= x1._1 && pos._1 <= x2._1 && pos._2 >= x1._2 && pos._2 <= x2._2
  }

  def testFunctionChanges(coord: (Int,Int)) : Unit = {
    for(fun <- matrix.functionStack if isInRange(coord, fun.from,fun.to))
      functionStack.+=:(fun) //FIXME: no duplicates

  }

  def constantEval(v: ConstantType) : Option[Int] = v.value

  def evalFunctionStack() : Unit = {
    //TODO same method as MatrixEvaluator
  }

  override def eval: ChangesData = {
    for((coord,v) <- source.userActionStack){
      testFunctionChanges(coord)
      v match {
        case v : ConstantType => {
          changesData.modificationList.+=:( (coord, constantEval(v)))
        }
        case f : FunctionType => {
          functionStack.+=:(f)
        }
      }
    }
    evalFunctionStack()
    changesData
  }
}
