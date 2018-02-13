use std::io::prelude::*;
use std::fs::File;

use cell;


pub fn read_file(f :&str) -> String
{
    let mut file = File::open(f).expect("Error Opening data.csv");
    let mut data = String::new();
    file.read_to_string(&mut data).expect("Error reading file data.csv");
    return data.trim().to_string();
}

pub fn init_formula(form_dec_vec: Vec<&str>) -> Box<cell::Cell>
{
    let cell = cell::Formula{
            num:0,
            r1:form_dec_vec[0].trim().parse()
            .expect("Erreur format"),
            c1:form_dec_vec[1].trim().parse()
            .expect("Erreur format"),
            r2:form_dec_vec[2].trim().parse()
            .expect("Erreur format"),
            c2:form_dec_vec[3].trim().parse()
            .expect("Erreur format"),
            val:form_dec_vec[4].trim().parse()
            .expect("Erreur format"),
    };
    Box::new(cell)
         
}

pub fn create_cell(str:String) -> Box<cell::Cell>
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
        let cell = cell::Number{
            value:val,
            };
        Box::new(cell)
    }
    
}
pub fn gen_table(data: String) -> Vec<Vec<Box<cell::Cell>>>
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
pub fn print_table(t:&Vec<Vec<Box<cell::Cell>>>)
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
