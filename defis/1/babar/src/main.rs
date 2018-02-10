use std::fs::File;
use std::io::prelude::*;
use std::env;
use std::io::{BufReader,BufRead};
use std::collections::BTreeMap;

trait Cell {
    fn get_value(&self) -> i32;
    fn set_value(&mut self,val:i32);
    fn evaluate(&self, t:&Vec<Vec<Box<Cell>>>,ce :&mut Vec<(i32,i32)>,r:i32,c:i32) -> (i32,BTreeMap<(i32,i32),Vec<(i32,i32)>>);
    fn copy_cell(&self) -> Box<Cell>;
    fn print_cell(&self) -> ();
    fn get_string_value(&self) -> String;
    fn get_region(&self) -> (i32,i32,i32,i32);
}
struct Number {value:i32,}
struct Formula {
    num:i32,
    r1:i32,
    c1:i32,
    r2:i32,
    c2:i32,
    v:i32,
 
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
    fn evaluate(&self, grid:&Vec<Vec<Box<Cell>>>,current_evaluation: &mut Vec<(i32,i32)>,posr:i32,posc:i32) 
        -> (i32,BTreeMap<(i32,i32),Vec<(i32,i32)>>)
    {
        current_evaluation.push((posr,posc));
        
        if !is_dependency_ok(self, current_evaluation) {
            return (-1,BTreeMap::new());
        }

        calcul_occ(self,grid,current_evaluation,posr,posc)
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

fn is_dependency_ok(cell : &Formula, current_evaluation : &Vec<(i32,i32)>) -> bool
{

        for &(row,col) in current_evaluation {
        if row >= cell.r1  && row <= cell.r2
          && col >= cell.c1 && col <= cell.c2 {
            return false;
          }
        }
        true
}

fn calcul_occ(cell:&Formula,grid:&Vec<Vec<Box<Cell>>>,current_evaluation:&mut Vec<(i32,i32)>,
    posr:i32,posc:i32) ->(i32,BTreeMap<(i32,i32),Vec<(i32,i32)>>)
{
    let r1 = cell.r1 as usize;
    let r2 = cell.r2 as usize;
    let c1 = cell.c1 as usize;
    let c2 = cell.c2 as usize;
    let mut dependences = BTreeMap::new();
    let mut val_num = 0;
    for i in r1..r2+1{
        for j in c1..c2+1{
            let (val,_) = grid[i][j].evaluate(grid,current_evaluation,i as i32,j as i32);
            if val < 0{
                return (val,BTreeMap::new());
            }
            dependences.insert((i as i32,j as i32),Vec::new());
            
            match dependences.get_mut(&(i as i32,j as i32)) {
                Some(dependency) => dependency.push((posr,posc)),
                _ => println!("err"),
            }
            if val == cell.v{
                val_num = val_num+1;
            }
        }
    }
    (val_num,dependences)

}

fn evaluate(grid: &Vec<Vec<Box<Cell>>>) -> (Vec<Vec<Box<Cell>>>,BTreeMap<(i32,i32),Vec<(i32,i32)>>)
{
    let mut new_grid: Vec<Vec<Box<Cell>>> =  Vec::new();
    let mut dependences : BTreeMap<(i32,i32),Vec<(i32,i32)>> =  BTreeMap::new();
    for i in 0..grid.len(){
        let mut row: Vec<Box<Cell>> =  Vec::new(); 
        for j in 0..grid[i].len(){
            let mut cell = grid[i][j].copy_cell();
            let mut current_evaluation : Vec<(i32,i32)> = Vec::new();
            let (val,d) = cell.evaluate(grid,&mut current_evaluation,i as i32,j as i32);
            cell.set_value(val);
            row.push(cell);
            for (&key, val) in d.iter(){
                if !dependences.contains_key(&key) {
                    dependences.insert(key,Vec::new());
                }
                match dependences.get_mut(&key) {
                    Some(dependency) => dependency.extend(val),
                    _ => println!("err"),
                }
                
            }
        }
        new_grid.push(row);
    }
    (new_grid,dependences)
}

fn read_file(f :&str) -> String
{
    let mut file = File::open(f).expect("Error Opening data.csv");
    let mut data = String::new();
    file.read_to_string(&mut data).expect("Error reading file data.csv");
    return data.trim().to_string();
}

fn init_formula(form_dec_vec: Vec<&str>) -> Box<Cell>
{
    let cell = Formula{
            num:0,
            r1:form_dec_vec[0].trim().parse()
            .expect("Erreur format"),
            c1:form_dec_vec[1].trim().parse()
            .expect("Erreur format"),
            r2:form_dec_vec[2].trim().parse()
            .expect("Erreur format"),
            c2:form_dec_vec[3].trim().parse()
            .expect("Erreur format"),
            v:form_dec_vec[4].trim().parse()
            .expect("Erreur format"),
    };
    Box::new(cell)
         
}

fn create_cell(str:String) -> Box<Cell>
{
    if Some('=') == str.chars().next() {
        let form : String = str.trim_matches(|c| c == '(' || c == ')' 
         || c == '=' || c == '#' ).to_string();
         let form_decompose = form.split(",");
         let form_dec_vec: Vec<&str> = form_decompose.collect();
         if form_dec_vec.len() < 5 {
            panic!("Erreur format");
         }
         
        init_formula(form_dec_vec)
    }else {
        let val : i32 = str.trim().parse()
        .expect("Erreur format");
        let cell = Number{
            value:val,
            };
        Box::new(cell)
    }
    
}
fn gen_table(data: String) -> Vec<Vec<Box<Cell>>>
{
    let mut t = Vec::new();
    let a =  data.split("\n");
    let vec = a.collect::<Vec<&str>>();
    for c in &vec{
        let mut temp = c.split(";");
        let  vec2 = temp.collect::<Vec<&str>>();
        let mut row = Vec::new();
        for d in &vec2{
            let cell = create_cell(d.to_string());
            row.push(cell);
            
        }
        t.push(row);
    }
    return t;
}
fn print_table(t:&Vec<Vec<Box<Cell>>>)
{
    let mut i:i32 = 0;
    for k in t{
        for b in k{
            if i!=0
            {
                print!(";");
            }
            i+=1;
            print!("{}", b.get_value());
        }
        i=0;
        println!("");
    }
}
fn write_view0(view0: &str,t:&Vec<Vec<Box<Cell>>>)
{
    let mut file = File::create(view0).expect("Error writing file");
    let mut mystring = String::new();
    let mut i:i32 = 0;
    for k in t{
        for b in k{
            if i!=0
            {
                mystring += ";";
            }
            i+=1;
            let tmp = b.get_string_value();
            mystring.push_str(&tmp);
        }
        i=0;
        mystring +="\n";
    }
    write!(file, "{}", mystring).expect("Error Writing into the view0");
}

fn write_change(user: &str,change:&str,spreadsheet:&mut Vec<Vec<Box<Cell>>>,dependencies:&mut BTreeMap<(i32,i32),Vec<(i32,i32)>>){
    let mut file_change = File::create(change).expect("Error at file creation");
    let file_user = File::open(user).expect("Error at file opening");
    for line in BufReader::new(file_user).lines() {
        let mut line_iter = line.expect("Error write_change");
        let mut ite = line_iter.trim().split_whitespace();
        let r = match ite.next(){
            Some(x) => x,
            None => continue,
        };
        let c = match ite.next(){
            Some(x) => x,
            None => continue,
        };
        let d = match ite.next(){
            Some(x) => x,  
            None => continue,
        };
        let r :i32 = r.trim().parse().expect("bad format");
        let c :i32 = c.trim().parse().expect("bad format");
        let mut current_evaluation : Vec<(i32,i32)> = Vec::new();
        let mut cell = create_cell(d.to_string());
        let (r1,r2,c1,c2) = spreadsheet[r as usize][c as usize].get_region();
        if (r1,r2,c1,c2) != (-1,-1,-1,-1) {
            for i in r1..r2 {
                for j in c1..c2 {
                    match dependencies.get_mut(&(i as i32,j as i32)) {
                        Some(dependence) => dependence.retain(|&x| x != (r,c)),
                        _ => println!("err"),
                    }
                }
            }
        }
        let (val,de) = cell.evaluate(spreadsheet,&mut current_evaluation,r,c);
        cell.set_value(val);
        for (&key, val) in de.iter(){
            if !dependencies.contains_key(&key) {
                dependencies.insert(key,Vec::new());
            }
            match dependencies.get_mut(&key) {
                Some(dependency) => dependency.extend(val),
                _ => (),
            }
                
        }
        spreadsheet[r as usize][c as usize] = cell;
        write!(file_change, "after  \"{} {} {}\":\n", r,c,d).expect("Error Writing into the change");
        match dependencies.get(&(r,c)) {
            Some(d) => 
                for &(x,y) in d {
                    let mut cell = spreadsheet[x as usize][y as usize].copy_cell();
                    let (val,_) = cell.evaluate(spreadsheet,&mut current_evaluation,x,y);
                    if val != spreadsheet[x as usize][y as usize].get_value(){
                        cell.set_value(val);
                        spreadsheet[x as usize][y as usize] = cell;
                        write!(file_change, "{} {} {}\n", x,y,val).expect("Error Writing into the change");
                    }
                },
            _ => (),
        }
    }
}

fn main()
{
    let args: Vec<String> = env::args().collect();
    if args.len() != 5
    {
        panic!("Erreur d'arguments, nombre : {}", args.len());
    }
    let data = read_file(&args[1]);

    let t = gen_table(data);
    
    print_table(&t);
    let (mut grid,mut dependences) = evaluate(&t);
    write_view0(&args[3],&grid);
    write_change(&args[2],&args[4],&mut grid,&mut dependences)
}    