package prog.comp2018.scableur.data

trait Value {
  def eval(view: View) : Option[Int]
}

//Constant value (0-255)
case class ConstantType(private var value: Option[Int]) extends Value {
  def eval(view: View) : Option[Int] = value
}
//Function
abstract case class FunctionType(var args: Array[String]) extends Value {
  def eval(view: View) : Option[Int]
}