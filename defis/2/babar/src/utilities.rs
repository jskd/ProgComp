use std;
use std::io::{ BufRead};
use std::fs::File;

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
	while num_bytes!=0 
	{
		has_formula(&buff, &formulas);                   //check if line contains formula, if so, add them to Formulas
		let br = reader.read_until(b'\n',&mut buff).expect("counting bytes in lines");
//  		bytes_by_lines.push(bytes_by_lines.pop().unwrap() + num_bytes);

	}



}
pub fn has_formula(line: Vec<u8>, formulas: &Vec<cell::Formula>)
{
	let lineproper = line.trim()
	let linesplit = lineproper.split(";")
	for_each!(n in linesplit
	{
		if(n[0] == "=")
		{
			println!("Formula !")
			formulas.append(create_formula(n))
		}
		else
		{
			println("Not a formula")
		}
	})
}

pub fn create_formula(form_dec_vec: Vec<&str>) -> formula::Formula
{
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