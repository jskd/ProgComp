use std::fs::File;
use std::io::prelude::*;
use std::env;

trait Cellule {
    fn getValue(&self) -> i32;
}
struct Number {value:i32}
struct Formule {form:String}

impl Number
{
    fn setValue(&mut self, n: i32)
    {
        self.value = n;
    }
}
impl Cellule for Number
{
    
    fn getValue(&self) -> i32
    {
        return self.value;
    }
}
impl Cellule for Formule
{
    // TODO
    fn getValue(&self) -> i32
    {
        return 0;
    }
}
fn read_file(f :&str) -> String
{
    let mut file = File::open(f).expect("Erreur à l'ouverture du fichier");
    let mut data = String::new();
    file.read_to_string(&mut data);
    return data.trim().to_string();
}

fn parsetab(d: &str) -> Vec<Vec<Box<Cellule>>>
{
    let t: Vec<Vec<Box<Cellule>>> =  Vec::new();

    return t; 
}

fn main()
{
    let args: Vec<String> = env::args().collect();
    if args.len() != 5
    {
        panic!("Erreur d'arguments");
    }
    let mut data = read_file(&args[1]);
    let mut Num: Number =  Number{value: 3};
    Num.setValue(4);
    let mut a =  data.split("\n");
    let vec = a.collect::<Vec<&str>>();
    for c in &vec{
        println!("-> {}", c.trim());
        let mut temp = c.split(";");
        let  vec2 = temp.collect::<Vec<&str>>();
        for d in &vec2{
            println!("L:{}", d);
        }
    }
}