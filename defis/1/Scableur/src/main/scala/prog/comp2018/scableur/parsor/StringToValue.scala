package prog.comp2018.scableur.parsor
import prog.comp2018.scableur.data.functions.NbrIteration
import prog.comp2018.scableur.data.{ConstantType, Matrix, Value}

object StringToValue {


  def stv(str: String, i : Int, j: Int) : Value = {


    if (str(0) == '=') {

      val pattern = """[#=(](\d+),(\d+),(\d+),(\d+),(\d+)[)]""".r

      if(str(1)=='#'){
        val parts=pattern.findAllIn(str)
        val i1=parts.group(1)
        val j1=parts.group(2)
        val i2=parts.group(3)
        val j2=parts.group(4)
        val value=parts.group(5)
        try{
          new NbrIteration((i,j),(i1.toInt,j1.toInt),(i2.toInt,j2.toInt),
                                               ConstantType(Some(value.toInt)))
        }catch {
          case _ : Exception => ConstantType(None)
        }

      }else{
        ConstantType(None)
      }

    }
    else{
      ConstantType(Some(str.toInt))
    }

  }

}
