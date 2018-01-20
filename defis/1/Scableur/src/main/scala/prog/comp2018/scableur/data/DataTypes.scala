package prog.comp2018.scableur.data

trait Value {
  def eval : Option[Int]
}

//Constant value (0-255)
case class ConstantType(private var value: Option[Int]) extends Value {
  def eval : Option[Int] = value
}
//Function
abstract case class FunctionType() extends Value {
  def eval : Option[Int]
}