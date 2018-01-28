use std::fs::File;
use std::io::prelude::*;
use std::env;

trait Cellule {
    fn get_value(&self) -> i32;
    fn set_value(&mut self,val:i32);
    fn evaluate(&self, t:&Vec<Vec<Box<Cellule>>>,&mut Vec<(i32,i32,i32,i32)>) -> i32;
    fn copy_cell(&self) -> Box<Cellule>;
    fn print_cell(&self) -> ();
}
struct Number {value:i32}
struct Formule {
    num:i32,
    r1:i32,
    c1:i32,
    r2:i32,
    c2:i32,
    v:i32,
}



impl Cellule for Number
{
    fn evaluate(&self, t:&Vec<Vec<Box<Cellule>>>,current_evaluation:&mut Vec<(i32,i32,i32,i32)>) -> i32
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
    fn copy_cell(&self) ->Box<Cellule>
    {
        let cell = Number{value:self.value,};
        Box::new(cell)
    }
    
    fn print_cell(&self)
    {
    }
}
impl Cellule for Formule
{
    fn evaluate(&self, grill:&Vec<Vec<Box<Cellule>>>,current_evaluation:&mut Vec<(i32,i32,i32,i32)>) -> i32
    {
        if !is_dependency_ok(self,current_evaluation) {
            panic!("Bad dependency");
        }

        calcul_occ(self,grill,current_evaluation)
            
    }
    fn get_value(&self) -> i32
    {
        return self.num;
    }
    
    fn set_value(&mut self, n: i32)
    {
        self.num = n;
    }
    fn copy_cell(&self) -> Box<Cellule>
    {
        let cell = Formule{
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
        println!("b evaluateCell: {} {} {} {} {} {}",self.r1,self.r2,self.c1,self.c2,self.v,self.num);
    }
}

fn is_dependency_ok(cell : &Formule, current_evaluation : &Vec<(i32,i32,i32,i32)>) -> bool
{
    for &(r1,c1,r2,c2) in current_evaluation {
        if cell.r1 >= r1 && cell.r2 <= r2 
          && cell.c1 >= c1 && cell.c2 <= c2 {
            return false;
          }
        }
        true
}

fn calcul_occ(cell:&Formule,grill:&Vec<Vec<Box<Cellule>>>,current_evaluation:&mut Vec<(i32,i32,i32,i32)>) ->i32
{
    let r1 = cell.r1 as usize;
    let r2 = cell.r2 as usize;
    let c1 = cell.c1 as usize;
    let c2 = cell.c2 as usize;
    current_evaluation.push((cell.r1,cell.c1,cell.r2,cell.c2));
    let mut val_num = 0;
    for i in r1..r2+1{
        for j in c1..c2+1{
             let val = grill[i][j].evaluate(grill,current_evaluation);
            if val == cell.v{
                val_num = val_num+1;
            }
                
        }
    } 
    val_num

}

fn evaluate(grill: &Vec<Vec<Box<Cellule>>>) -> Vec<Vec<Box<Cellule>>>
{
    let mut new_grill: Vec<Vec<Box<Cellule>>> =  Vec::new(); 
    for i in 0..grill.len(){
        let mut row: Vec<Box<Cellule>> =  Vec::new(); 
        for j in 0..grill[i].len(){
            let mut cell = grill[i][j].copy_cell();
            let mut current_evaluation : Vec<(i32,i32,i32,i32)> = Vec::new();
            cell.print_cell();
            let val = cell.evaluate(grill,&mut current_evaluation);
            cell.set_value(val);
            cell.print_cell();
            row.push(cell);
            
        }
        new_grill.push(row);
    }
    new_grill
}
fn read_file(f :&str) -> String
{
    let mut file = File::open(f).expect("Error Opening data.csv");
    let mut data = String::new();
    file.read_to_string(&mut data).expect("Error reading file data.csv");
    return data.trim().to_string();
}

fn parsetab(d: &str) -> Vec<Vec<Box<Cellule>>>
{
    let t: Vec<Vec<Box<Cellule>>> =  Vec::new();
    return t; 
}

fn init_formule(form_dec_vec: Vec<&str>) -> Box<Cellule>
{
    let cell = Formule{
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

fn create_cell(str:String) -> Box<Cellule>
{
    if Some('=') == str.chars().next() {
        let form : String = str.trim_matches(|c| c == '(' || c == ')' 
         || c == '=' || c == '#' ).to_string();
         let form_decompose = form.split(",");
         let form_dec_vec: Vec<&str> = form_decompose.collect();
         if form_dec_vec.len() < 5 {
            panic!("Erreur format");
         }
         
        init_formule(form_dec_vec)
    }else {
        let val : i32 = str.trim().parse()
        .expect("Erreur format");
        let cell = Number{value:val,};
        Box::new(cell)
    }
    
}
fn gen_table(data: String) -> Vec<Vec<Box<Cellule>>>
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
fn print_table(t:&Vec<Vec<Box<Cellule>>>)
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
fn write_view0(view0: &str,t:&Vec<Vec<Box<Cellule>>>)
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
            let tmp = b.get_value().to_string();
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
        panic!("Erreur d'arguments");
    }
    let data = read_file(&args[1]);

    let t = gen_table(data);
    
    print_table(&t);
    let grill = evaluate(&t);
    write_view0(&args[3],&grill);
    

}    