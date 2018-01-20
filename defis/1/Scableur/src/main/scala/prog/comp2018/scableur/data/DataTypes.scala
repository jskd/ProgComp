package prog.comp2018.scableur.data

trait Area[T] {

}

trait History[T] {
  def getModifications : Iterable[T]
}

trait View[T] {
  def getIterable : Iterable[T]
  def getHistory : History[T]
}


trait Value

//Constant value (0-255)
case class Constant() extends Value
//Function
case class Function() extends Value

// Function types
case class IterationsOfVInArea(a: Area[Value], value: Value) extends Function {
  
}





