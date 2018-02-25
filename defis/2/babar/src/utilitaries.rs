use std::io;
use std::fs::File;
use std::io::BufReader;

use cell;


pub const BUFF_SIZE: usize = 500;

///count bytes by lines, & get all formula in formula tab
pub fn read_first_time(path: &str, bytes_by_lines: &Vec<usize>, formulas: &Vec<cell::Formula>)
{
	let file = File::open("foo.txt").expect("fail to open");
	let mut buff = Vec::with_capacity(BUFF_SIZE);
	let mut reader = BufReader::(file);
	let mut num_bytes = reader.read_until(b'\n',&mut buff).expect("counting bytes in lines");
	bytes_by_lines.push(num_bytes);
	while num_bytes!=0 
	{
		has_formula(buff, formulas);                   //check if line contains formula, if so, add them to Formulas
		reader.read_until(b'\n',&mut buff).expect("counting bytes in lines");
		bytes_by_lines.push(bytes_by_lines.pop().unwrap() + num_bytes);

	}



}

pub fn has_formula(line: Vec<u8>, formulas: &Vec<cell::Formula>)
{

}

