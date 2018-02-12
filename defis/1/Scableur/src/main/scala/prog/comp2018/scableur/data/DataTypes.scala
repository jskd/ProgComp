package prog.comp2018.scableur.data

trait Value {
  def ==(v : Value) : Boolean
}

//Constant value (Int : 0-255 or P)
case class ConstantType(var value: Option[Int]) extends Value {

  override def ==(v: Value): Boolean = {
    v match {
      case ct : ConstantType => this.value == ct.value
      case _ : FunctionType => false
    }
  }

}
//Type for functions
abstract case class FunctionType(var coordinates  : (Int, Int),
                                 var from : (Int,Int),
                                 var to: (Int, Int),
                                 var value: Value) extends Value {

  def ==(v : Value) : Boolean = {
    v match{
      case ft: FunctionType => 
        (this.coordinates == ft.coordinates &&
         this.from == ft.from &&
         this.to == ft.to &&
         this.value == ft.value)
      case _ => false
    }
  }
}
