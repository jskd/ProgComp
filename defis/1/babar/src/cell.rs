//!This is documentaition for the 'cell' crate.
//!The cell crate is meant to be used as a way to manipulate raw data from the files.

use std::collections::BTreeMap;
use treatment;

///A Cell contains data, either as a Number or as a Formula.
///Defines a set of operations on its content.
pub trait Cell
{
    fn evaluate(&self,
                        spreadsheet: &Vec<Vec<Box<Cell>>>,
                        current_evaluation: &mut Vec<(i32, i32)>,
                        row: i32,
                        col: i32)
    -> (i32, BTreeMap<(i32, i32),Vec<(i32, i32)>>);

    fn copy_cell(&self) -> Box<Cell>;
    fn get_region(&self) -> (i32, i32, i32, i32);
    fn get_string_value(&self) -> String;
    fn get_value(&self) -> i32;
    fn set_value(&mut self, val:i32);
    fn print_cell(&self) -> ();
}

///A Number represents an integer in the data.
pub struct Number
{
    pub value: i32,
}

///A Formula represents a formula in the data.
pub struct Formula
{
    pub num: i32, //TODO
    pub r1: i32,
    pub c1: i32,
    pub r2: i32,
    pub c2: i32,
    pub val: i32,
}

impl Cell for Number
{
    ///Returns the value of the Cell since it is only a number.
    fn evaluate(&self,
                        _spreadsheet: &Vec<Vec<Box<Cell>>>,
                        _current_evaluation: &mut Vec<(i32, i32)>,
                        _row: i32,
                        _col: i32)
    -> (i32, BTreeMap<(i32, i32),Vec<(i32, i32)>>)
    {
        (self.value, BTreeMap::new())
    }

    ///Returns a Box containing a copy of the current Cell.
    fn copy_cell(&self) -> Box<Cell>
    {
        let cell = Number{value:self.value};
        Box::new(cell)
    }

    ///Had to be implemented but is never called.
    fn get_region(&self) -> (i32 ,i32, i32, i32){
        (-1,-1,-1,-1)
    }

    fn get_string_value(&self) -> String
    {
        self.value.to_string()
    }

    fn get_value(&self) -> i32
    {
        return self.value;
    }

    fn set_value(&mut self, n: i32)
    {
        self.value = n;
    }

    ///Had to be implemented but is never called.
    fn print_cell(&self)
    {
    }
}

impl Cell for Formula
{
    ///Evaluates a Formula :
    ///- if it is properly formatted and no cyclical dependancy is found, returns the result of the Formula
    ///- if there is any issue, return -1.
    fn evaluate(&self,
                        spreadsheet: &Vec<Vec<Box<Cell>>>,
                        current_evaluation: &mut Vec<(i32, i32)>,
                        row: i32,
                        col: i32)
    -> (i32, BTreeMap<(i32, i32),Vec<(i32, i32)>>)
    {
        current_evaluation.push((row, col));
        if !treatment::is_dependency_ok(self, current_evaluation) {
            return (-1, BTreeMap::new());
        }
        treatment::calcul_occ(self, spreadsheet, current_evaluation, row, col)
    }

    ///Returns a Box containing a copy of the current Formula.
     fn copy_cell(&self) -> Box<Cell>
    {
        let cell = Formula{
            num: 0,
            r1: self.r1,
            c1: self.c1,
            r2: self.r2,
            c2: self.c2,
            val: self.val,
        };
        Box::new(cell)
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

    fn print_cell(&self)
    {
        println!("b evaluateCell: {} {} {} {} {} {}", self.r1, self.c1, self.r2, self.c2, self.val, self.num);
    }

    ///Returns four integers corresponding to the edges of the field the current Formula applies to.
    fn get_region(&self) -> (i32,i32,i32,i32){
        (self.r1, self.r2, self.c1, self.c2)
    }
}
