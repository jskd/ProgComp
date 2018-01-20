package prog.comp2018.scableur.utils

object Debug {
  def d(msg : String) : Unit = println("[DEBUG]: " + msg)
  def e(msg : String) : Unit = System.err.println("[ERROR]: " + msg)
  def i(msg : String) : Unit = println("[INFO]: " + msg)
}
