package scableur

import scableur.data._

package object cycles {

  type FormulaList = scala.collection.mutable.Map[Point,PCountFormula]

  @volatile private var cycle_functions : List[Point] = List()
  @volatile private var formulaList: FormulaList = 
    scala.collection.mutable.Map()

  def setFormulaList(formulaListe: FormulaList){
    this.formulaList = formulaList
  }

  /*find_cycle: Calculates the list of the functions of the cycle
   arguments:
    f: is the calling function 
    g: is the called function
    dep: dependency list between function on formulaList for each 
   call to the function find_cycle we add the the calling function
   to the dependency list*/
  def find_cycle(f:PCountFormula, dep: List[Point]): Unit = {

    for ((p,g) <- this.formulaList){
     
      if(isPointInArea(p,f.defArea)){

        cycle_functions = (cycle_functions.:::(dep.+:(f.p))).distinct
        return
      }else{
        find_cycle(g, (dep.+:(f.p)))
      }
    }
  }

  def cycle_removal() : FormulaList = {

     for( (p,f) <- formulaList){
           find_cycle(f, List())
     }

     for(p <-  cycle_functions) {
       //set the function at point p as invalid
       this.formulaList(p).setInvalid()
     }

    return this.formulaList
  }
}

