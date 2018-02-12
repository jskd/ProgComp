use std::collections::BTreeMap;
use treatment;

pub trait Cell {
    fn get_value(&self) -> i32;
    fn set_value(&mut self,val:i32);
    fn evaluate(&self, t:&Vec<Vec<Box<Cell>>>,ce :&mut Vec<(i32,i32)>,r:i32,c:i32) -> (i32,BTreeMap<(i32,i32),Vec<(i32,i32)>>);
    fn copy_cell(&self) -> Box<Cell>;
    fn print_cell(&self) -> ();
    fn get_string_value(&self) -> String;
    fn get_region(&self) -> (i32,i32,i32,i32);
}
pub struct Number {pub value:i32,}
pub struct Formula {
    pub num:i32,
    pub r1:i32,
    pub c1:i32,
    pub r2:i32,
    pub c2:i32,
    pub v:i32,
 
}

impl Cell for Number
{
    fn evaluate(&self, _t:&Vec<Vec<Box<Cell>>>,_current_evaluation : &mut Vec<(i32,i32)>,_r:i32,_c:i32)
    -> (i32,BTreeMap<(i32,i32),Vec<(i32,i32)>>)
    {
        (self.value,BTreeMap::new())
    }
    fn get_value(&self) -> i32
    {
        return self.value;
    }
    
    fn set_value(&mut self, n: i32)
    {
        self.value = n;
    }
    fn copy_cell(&self) ->Box<Cell>
    {
        let cell = Number{value:self.value};
        Box::new(cell)
    }
    
    fn get_string_value(&self) -> String
    {
        self.value.to_string()
    }
    
    fn print_cell(&self)
    {
    }
    
    fn get_region(&self) -> (i32,i32,i32,i32){
        (-1,-1,-1,-1)
    }
}

impl Cell for Formula
{
    fn evaluate(&self, spreadsheet:&Vec<Vec<Box<Cell>>>,current_evaluation: &mut Vec<(i32,i32)>,posr:i32,posc:i32) 
        -> (i32,BTreeMap<(i32,i32),Vec<(i32,i32)>>)
    {
        current_evaluation.push((posr,posc));
        
        if !treatment::is_dependency_ok(self, current_evaluation) {
            return (-1,BTreeMap::new());
        }

        treatment::calcul_occ(self,spreadsheet,current_evaluation,posr,posc)
    }
    
    fn get_string_value(&self) -> String
    {
        if self.num < 0 {
            return "P".to_string();
        }
        self.num.to_string()
    }
    
    fn get_value(&self) -> i32
    {
        return self.num;
    }
    
    fn set_value(&mut self, n: i32)
    {
        self.num = n;
    }
    fn copy_cell(&self) -> Box<Cell>
    {
        let cell = Formula{
            num:0,
            r1:self.r1,
            c1:self.c1,
            r2:self.r2,
            c2:self.c2,
            v:self.v,
        };
        Box::new(cell)
    }
    fn print_cell(&self)
    {
        println!("b evaluateCell: {} {} {} {} {} {}",self.r1,self.c1,self.r2,self.c2,self.v,self.num);
    }
    
    fn get_region(&self) -> (i32,i32,i32,i32){
        let r1 = self.r1;
        let r2 = self.r2;
        let c1 = self.c1;
        let c2 = self.c2;
        (r1,r2,c1,c2)
    }
}