package prog.comp2018.scableur.data

trait Value {

  def ==(v: Value): Boolean = {
    v match{
      case v: ConstantType => true //vrai pour l'instant
      case v: FunctionType => true //vrai pour l'instant
      case _ => false
    }
  }
}

//Constant value (0-255)
case class ConstantType(private var value: Option[Int]) extends Value {

  override def ==(v: Value): Boolean = this.getClass == v.getClass && compareValue(v)

  def compareValue(value: Value): Boolean = {

    false
  }
}
//Function
abstract case class FunctionType(var coord: Array[Int], v: Value) extends Value {
}