use std::io::prelude::*;
use std::fs::File;

use cell;

///Opens the file f and returns its content as a String.
pub fn read_file(f: &str) -> String
{
    let mut file = File::open(f).expect("Error Opening data.csv");
    let mut data = String::new();
    file.read_to_string(&mut data).expect("Error reading file data.csv");
    return data.trim().to_string();
}

///Creates a new cell containing the formula given in parameter.
pub fn init_formula(form_dec_vec: Vec<&str>) -> Box<cell::Cell>
{
    let cell = cell::Formula{
            num: 0,
            r1: form_dec_vec[0].trim().parse()
            .expect("Erreur format"),

            c1: form_dec_vec[1].trim().parse()
            .expect("Erreur format"),

            r2: form_dec_vec[2].trim().parse()
            .expect("Erreur format"),

            c2: form_dec_vec[3].trim().parse()
            .expect("Erreur format"),

            val: form_dec_vec[4].trim().parse()
            .expect("Erreur format"),
    };
    Box::new(cell)
}

pub fn create_cell(str:String) -> Box<cell::Cell>
{
    if Some('=') == str.chars().next() {
        //TODO : what happens if trim_matches does not return a string ?
        let form : String = str.trim_matches(|c| c == '(' || c == ')' || c == '=' || c == '#' ).to_string();
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

///Generates a cell matrix from the data contained in the string passed in parameter.
pub fn gen_table(data: String) -> Vec<Vec<Box<cell::Cell>>>
{
    let mut table = Vec::new();
    let data_str =  data.split("\n");
    let data_collection = data_str.collect::<Vec<&str>>();
    for row_str in &data_collection{
        let mut temp = row_str.split(";");
        let row_collection = temp.collect::<Vec<&str>>();
        let mut row = Vec::new();
        for elem in &row_collection{
            let cell = create_cell(elem.to_string());
            row.push(cell);
        }
        table.push(row);
    }
    return table;
}

pub fn print_table(table: &Vec<Vec<Box<cell::Cell>>>)
{
    let mut i:i32 = 0;
    for row in table{
        for elem in row{
            if i!=0
            {
                print!(";");
            }
            i+=1;
            print!("{}", elem.get_value());
        }
        i=0;
        println!("");
    }
}
