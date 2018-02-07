package prog.comp2018.scableur.eval

import prog.comp2018.scableur.data.functions.NbrIteration
import prog.comp2018.scableur.data.{EvaluatedMatrix, FunctionType, Matrix, Value, ConstantType}
import prog.comp2018.scableur.utils.Conf

/**
  * Evaluates a FunctionType to Option[Int]
  */
trait FunctionEval[T] {
  def eval(f: FunctionType, m: Matrix): Option[Int]
}


object NbrIterationEval extends FunctionEval[Matrix] {
 
  def eval(f: NbrIteration, m: Matrix, dependencies: List[Int] = List()): Option[Int] =  {
    var counter = 0
    for{
      i <- Range(f.from._1,(f.to._1 + 1))
      j <- Range(f.from._2, (f.to._2 + 1))
    }{

      (m.get(i,j)) match{

        case x:ConstantType => if(x == f.value){counter += 1}

        case g:NbrIteration =>
          try{//TEST IF CYCLE EXIST RETURN None ELSE EVALUATE g 
            if(dependencies.exists(i => i==(m.functionStack.indexOf(g)))){
              m.set(ConstantType(None), f.coordinates._1, f.coordinates._2)
              return None
            }else{
             
              val reslt = eval(g,m,( dependencies.+: (m.functionStack.indexOf(f))))
              if(ConstantType(reslt) == f.value)
                counter += 1
            }
          }catch{
            case _ : Exception =>
              return None
            }
         case _ => 

       }
       if(counter > Conf.MAX_VALUE) return None
     }
     Some(counter)
   }

  override def eval(f: FunctionType, m: Matrix): Option[Int] = None
}



