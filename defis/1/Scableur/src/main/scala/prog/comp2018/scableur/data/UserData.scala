package prog.comp2018.scableur.data

import scala.collection.mutable.ListBuffer


class UserData(var userActionStack : ListBuffer[((Int,Int), Value)]) {

  def length(): Int = userActionStack.length

  def get(i: Int): Value = userActionStack(i)._2

  def getCoord(i:Int) : (Int, Int) = userActionStack(i)._1

  def set(coord: (Int, Int), v: Value, i: Int): Unit = userActionStack(i) = (coord, v)
}



class ChangesData(){
  type EvalCell = ((Int,Int), Option[Int])
  type ChangesList = ListBuffer[EvalCell]
  type AfterAction = ( ((Int,Int), Value), ChangesList)
  var modificationList : ListBuffer[AfterAction] = ListBuffer[AfterAction]()


  def length(): Int = modificationList.length

  def get(i: Int, j:Int): Option[Int] = (modificationList(i)._2)(j)._2

  def getCoord(i:Int, j:Int) : (Int, Int) = (modificationList(i)._2)(j)._1

  def set(coord: (Int, Int), v: Option[Int], i: Int, j: Int): Unit = {
    modificationList(i)._2(j) = (coord, v)
  }
}
