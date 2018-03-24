package prog.comp2018.scableur.data

trait Value {
  def ==(v : Value) : Boolean
  def to_String(): String
}

//Constant value (Int : 0-255 or P)
case class ConstantType(var value: Option[Int]) extends Value {

  override def ==(v: Value): Boolean = {
    v match {
      case ct : ConstantType => this.value == ct.value
      case _ : FunctionType => false
    }
  }

  override def to_String(): String = {
    var s: String = ""
    value match {
      case Some(i: Int) => s = i.toString
      case None => s = "P"
    }
    return s
  }

}
//Type for formulas
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

  override def to_String(): String = {
    var s: String = "=#("+from._1+","+from._2+","+to._1+","+to._2+","+value.to_String+")"
    return s
  }
}
