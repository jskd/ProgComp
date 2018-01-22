package prog.comp2018.scableur.eval

import java.util

import prog.comp2018.scableur.data.Value


class Evaluator(private var matrix :util.ArrayList[util.ArrayList[Value]]) {
  private var result : util.ArrayList[util.ArrayList[Option[Int]]] = new util.ArrayList[util.ArrayList[Option[Int]]]()


  def eval : util.ArrayList[util.ArrayList[Option[Int]]] = {
    //TODO evaluate lines of Array
    result
  }
}
