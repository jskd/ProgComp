use std;
use std::io::{ BufRead};
use std::fs::File;
use std::str;
// mod parser;
use cell;
// mod treatment;


pub const BUFF_SIZE: usize = 16000;

///count bytes by lines, & get all formula in formula tab rmove counting
pub fn read_first_time(path: &str, formulas: &mut Vec<cell::Formula>)
{
	let file = File::open(path).expect("fail to open");
	let mut buff = Vec::with_capacity(BUFF_SIZE);
	let mut reader = std::io::BufReader::new(file);
	let mut num_bytes = reader.read_until(b'=',&mut buff).expect("read until formula");
	while num_bytes!=0 //Buffer not empty0
	{

		let mut index:usize = 0;
		while has_formula(&buff){ //If buffer still got a formula
			let mut formula: Vec<u8> = Vec::new();
			formula.push(b'=');
				for byte in &buff{
					if byte == (&b';')
					{
						break;
					}
					else
					{
						formula.push(*byte);
					}
				}
			formulas.push(create_formula(String::from_utf8(formula).unwrap()));
		}
		num_bytes = reader.read_until(b'=',&mut buff).expect("read until formula or end file");
	}
}


pub fn check_bounds(buff: &Vec<u8>, index: usize) -> bool
{
	match buff.get(index..){
		Some(buff_ok) => {
			match buff.iter().position(|&eq| eq == b'='){
				Some(iter_ok) => true,
				None => false,
			}
		},
		None => false,
	}
}
pub fn has_formula(line: &Vec<u8>) -> bool
{
 line.contains(&b'=')
}

pub fn create_formula(form_string: String) -> cell::Formula
{
	let form : String = form_string.trim_matches(|c| c == '(' || c == ')' || c == '=' || c == '#' ).to_string();
    let form_decompose = form.split(",");
    let form_dec_vec: Vec<&str> = form_decompose.collect();
    if form_dec_vec.len() < 5 {
        panic!("Erreur format");
    }
         
    let formula = cell::Formula{
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
    return formula
}