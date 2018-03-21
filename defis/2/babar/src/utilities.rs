use std;
use std::io::{BufRead,BufReader,Read};
use std::fs::File;
use std::fs;
use std::io::prelude::*;
use std::str;
use std::thread;
// mod parser;
use cell;
use searchTree;
// mod treatment;


pub const BUFF_SIZE: usize = 16384;

///count bytes by lines, & get all formula in formula tab rmove counting
pub fn read_first_time(path: &str, formulas: &mut Vec<cell::Formula>)
{
 	let file = File::open(path).expect("fail to open");
	let mut buff = Vec::with_capacity(BUFF_SIZE);
	let mut reader = std::io::BufReader::new(file);
	let mut num_bytes = reader.read_until(b'=',&mut buff).expect("read until formula");
	buff.clear();
	buff.clear();
	num_bytes=reader.read_until(b')',&mut buff).expect("read formula");
	while num_bytes!=0 //Buffer not empty0
	{
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
		/*Ca marche mais c'est pas bon*/
		/*let thread = thread::spawn(move ||
		{create_formula(String::from_utf8(formula).unwrap())});
		let res=thread.join();
		match res {
			Ok(f) => {
				formulas.push(f);
			},
			Err(e) => panic!("thread child return None")
			// add code here
		}*/
		num_bytes = reader.read_until(b'=',&mut buff).expect("read until formula or end file");
		buff.clear();
		num_bytes = reader.read_until(b')',&mut buff).expect("read file");
		
	}
	

// 	let file = File::open(path).expect("fail to open");
// 	let mut buff = BufReader::with_capacity(BUFF_SIZE,file);
//         loop {
//         let length  = {
//             let mut buffer = buff.fill_buf().expect("err read_first_time");
//             let mut line = vec![];
//             let mut num_byte = 0;
//             loop{
//                 let mut n = buffer.read_until(b'=',&mut line).expect("err");
//                 num_byte = num_byte + n;
//                 let mut formula = vec![];
//                 formula.push(b'=');
//                 n = buffer.read_until(b')',&mut formula).expect("err");
//                 match formula.last(){
//                     Some(&b')') => {
//                     num_byte = num_byte + n;
//                     },
//                     _ => { break},
//                 }
//                  let s  = String::from_utf8(formula).expect("err converting [u8] to string");
//                 formulas.push(create_formula(s));
//               
//                 
//             }
//             num_byte
//         };
//         
//         
//         if length == 0 { break; }
//         buff.consume(length);
//    }
    println!("len: {}",formulas.len());
        
}

pub fn write_view(path:&str,formulas: &mut Vec<String>){
    let mut count : i32 = 0;
    let mut view = File::create("view0.csv").expect("Error creating file");
    let file = File::open(path).expect("fail to open");
    let mut buff = BufReader::with_capacity(BUFF_SIZE,file);
    let mut f = match  formulas.pop(){
        Some(x) =>x,
        None => return,
    };
    loop {
    let length  = {
        let mut buffer = buff.fill_buf().expect("err read_first_time");
        let mut line = String::new();
        let num_byte = buffer.read_line(&mut line).expect("err");
        while line.contains(&f){
            line = line.replace(&f,"-1");
            f = match formulas.pop(){
                Some(x) => x,
                None => break,
            };
        }
        write!(view,"{}",line).expect("Error Writing into the view0");

            
        num_byte
    };
        count = count + 1;
    if length == 0 { break; }
    buff.consume(length);
}

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

pub fn get_area(formula: cell::Formula, path: &str ,buff_target: &mut Vec<u8>){
	let file = File::open(path).expect("fail to open");
	let mut reader = std::io::BufReader::new(file);
	let mut jump = String::new();
	for i in 0..formula.r1-1{
		reader.read_line(&mut jump).expect("jumping to r1");
	}
	for i in 0..formula.c1-1{
		reader.read_until(b';', buff_target);
	}
	buff_target.clear();
	if formula.r2!=formula.r1{
		for i in 0..formula.r2-formula.r1{
			reader.read_until(b'\n', buff_target);
		}
		for i in 0..formula.c2-1{
			reader.read_until(b';', buff_target);	
		}
	}
	else{
		for i in 0..formula.c2-formula.c1{
			reader.read_until(b';', buff_target);
		}
	}
}