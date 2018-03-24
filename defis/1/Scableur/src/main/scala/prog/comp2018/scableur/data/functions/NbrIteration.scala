package prog.comp2018.scableur.data.functions

import prog.comp2018.scableur.data._

// Count formula
class NbrIteration(coordinates  : (Int, Int),
                   from : (Int,Int),
                   to: (Int, Int),
                   value: Value)  extends FunctionType(coordinates, from, to, value) {

  override def ==(v: Value): Boolean = false
}