use std::fs::File;
use std::io::prelude::*;
use std::env;

struct Entry<T>{
    value: T
}
fn read_file(f :&str) -> String
{
    let mut file = File::open(f).expect("Erreur à l'ouverture du fichier");
    let mut data = String::new();
    file.read_to_string(&mut data);
    return data.trim().to_string();
}
/*
fn parsetab(d: &str) -> Vec<Vec<Entry<_>>>
{

}
*/
fn main()
{
    let args: Vec<String> = env::args().collect();
    if args.len() != 5
    {
        panic!("Erreur d'arguments");
    }
    let mut data = read_file(&args[1]);
    println!("{}", data);
}