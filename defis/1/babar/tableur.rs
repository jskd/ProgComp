use std::fs::File;
use std::io::prelude::*;
use std::env;

trait Cellule {
    fn get_value(&self) -> i32;
    fn evaluate(&self, t:&Vec<Vec<Box<Cellule>>>) -> ();
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

impl Number
{
    fn set_value(&mut self, n: i32)
    {
        self.value = n;
    }
}
impl Cellule for Number
{
    fn evaluate(&self, t:&Vec<Vec<Box<Cellule>>>) -> ()
    {

    }
    fn get_value(&self) -> i32
    {
        return self.value;
    }
}
impl Cellule for Formule
{
    fn evaluate(&self, t:&Vec<Vec<Box<Cellule>>>) -> ()
    {

    }
    fn get_value(&self) -> i32
    {
        return 0;
    }
}
fn evaluate(t: &Vec<Vec<Box<Cellule>>>)
{
    for a in t{
        for b in a{
            b.evaluate(t);
        }
    }
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
fn write_view0(t:&Vec<Vec<Box<Cellule>>>)
{
    let mut file = File::create("view0.csv").expect("Error writing file");
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
    write!(file, "{}", mystring).expect("Error Writing into view0.csv");
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
    write_view0(&t);

}    