use std;
use std::io::{ BufRead};
use std::fs::File;
use std::str;
// mod parser;
use cell;
// mod treatment;


pub const BUFF_SIZE: usize = 500;

///count bytes by lines, & get all formula in formula tab
pub fn read_first_time(path: &str, bytes_by_lines: &mut Vec<usize>, formulas: &mut Vec<cell::Formula>)
{
	let file = File::open(path).expect("fail to open");
	let mut buff = Vec::with_capacity(BUFF_SIZE);
	let mut reader = std::io::BufReader::new(file);
	let mut num_bytes = reader.read_until(b'\n',&mut buff).expect("counting bytes in lines");
	bytes_by_lines.push(num_bytes);
	let mut tmp:usize;
	while num_bytes!=0 //Buffer not empty
	{
		let mut index:usize = 0;
		while has_formula(&buff){ //If buffer still got a formula
			let mut formula: Vec<u8> = Vec::new();
			let flag:bool;
			 match buff.get(index..){ //handling buffer's bounds
				Some(_) => //if buffer has formula
				{	
					flag = true;
				},
			None => flag = false,	//if buffer has no more formula
			}
			if flag // has formula do
			{
				buff = (buff[index..]).to_vec();
				index = buff.iter().position(|&eq| eq == b'=').unwrap();
				
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
				index+=formula.len();
				println!("{}", String::from_utf8(formula).unwrap())
				//formulas.push(create_formula(String::from_utf8(formula).unwrap()));
			}
			else  //has no more formula do 
			{
				break;
			}
		}
		num_bytes = reader.read_until(b'\n',&mut buff).expect("counting bytes in lines");
		tmp = bytes_by_lines.pop().unwrap();
		bytes_by_lines.push(tmp);  		
		bytes_by_lines.push(tmp + num_bytes);
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