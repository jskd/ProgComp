package prog.comp2018.scableur.data


trait ViewBuilder {
  // Build a Partial View from data
  def build(i: Int, j: Int, width: Int, height: Int ) : View[Value]
}


