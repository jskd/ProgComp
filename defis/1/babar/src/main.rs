use std::fs::File;
use std::io::prelude::*;
use std::env;

trait Cell {
    fn get_value(&self) -> i32;
    fn set_value(&mut self,val:i32);
    fn evaluate(&self, t:&Vec<Vec<Box<Cell>>>,ce :&mut Vec<(i32,i32)>,r:i32,c:i32) -> i32;
    fn copy_cell(&self) -> Box<Cell>;
    fn print_cell(&self) -> ();
    fn get_string_value(&self) -> String;
}
struct Number {value:i32}
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
    fn evaluate(&self, t:&Vec<Vec<Box<Cell>>>,current_evaluation : &mut Vec<(i32,i32)>,r:i32,c:i32)-> i32
    {
        self.value
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
        let cell = Number{value:self.value,};
        Box::new(cell)
    }
    
    fn get_string_value(&self) -> String
    {
        self.value.to_string()
    }
    
    fn print_cell(&self)
    {
    }
}
impl Cell for Formula
{
    fn evaluate(&self, grid:&Vec<Vec<Box<Cell>>>,current_evaluation: &mut Vec<(i32,i32)>,posr:i32,posc:i32) -> i32
    {
        current_evaluation.push((posr,posc));
        
        if !is_dependency_ok(self, current_evaluation) {
            return -1;
        }
        

        calcul_occ(self,grid,current_evaluation)
            
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
}

fn is_dependency_ok(cell : &Formula, current_evaluation : &Vec<(i32,i32)>) -> bool
{

        for &(row,col) in current_evaluation {
        if row >= cell.r1  && row <= cell.r2
          && col >= cell.c1 && col <= cell.c2 {
             println!("r1:{} row:{} r2:{} c1:{} col: {} c2:{}",cell.r1,row,cell.r2,cell.c1,col,cell.c2);
            return false;
          }
        }
        true
}

fn calcul_occ(cell:&Formula,grid:&Vec<Vec<Box<Cell>>>,current_evaluation:&mut Vec<(i32,i32)>) ->i32
{
    let r1 = cell.r1 as usize;
    let r2 = cell.r2 as usize;
    let c1 = cell.c1 as usize;
    let c2 = cell.c2 as usize;
    let mut val_num = 0;
    for i in r1..r2+1{
        for j in c1..c2+1{
            let val = grid[i][j].evaluate(grid,current_evaluation,i as i32,j as i32);
            if val < 0{
                return val;
            }
            if val == cell.v{
                val_num = val_num+1;
            }
        }
    } 
    val_num

}

fn evaluate(grid: &Vec<Vec<Box<Cell>>>) -> Vec<Vec<Box<Cell>>>
{
    let mut new_grid: Vec<Vec<Box<Cell>>> =  Vec::new(); 
    for i in 0..grid.len(){
        let mut row: Vec<Box<Cell>> =  Vec::new(); 
        for j in 0..grid[i].len(){
            let mut cell = grid[i][j].copy_cell();
            let mut current_evaluation : Vec<(i32,i32)> = Vec::new();
            cell.print_cell();
            let val = cell.evaluate(grid,&mut current_evaluation,i as i32,j as i32);
            cell.set_value(val);
            cell.print_cell();
            row.push(cell);
        }
        new_grid.push(row);
    }
    new_grid
}
fn read_file(f :&str) -> String
{
    let mut file = File::open(f).expect("Error Opening data.csv");
    let mut data = String::new();
    file.read_to_string(&mut data).expect("Error reading file data.csv");
    return data.trim().to_string();
}

fn parsetab(d: &str) -> Vec<Vec<Box<Cell>>>
{
    let t: Vec<Vec<Box<Cell>>> =  Vec::new();
    return t; 
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
    println!("cell: {} {} {} {} {}",cell.r1,cell.r2,cell.c1,cell.c2,cell.v);
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
        let cell = Number{value:val,};
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
            let s = d.to_string();
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
    let grid = evaluate(&t);
    write_view0(&args[3],&grid);
    

}    