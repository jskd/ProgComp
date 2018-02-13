mod parser;
mod treatment;
mod cell;
use std::env;

fn main()
{
    let args: Vec<String> = env::args().collect();
    if args.len() != 5
    {
        panic!("Erreur d'arguments, nombre : {}", args.len());
    }

    let file_content_str = parser::read_file(&args[1]);
    let structured_data = parser::gen_table(file_content_str);
    parser::print_table(&structured_data);

    let (mut grid,mut dependences) = treatment::evaluate(&structured_data);
    treatment::write_view0(&args[3],&grid);
    treatment::write_change(&args[2],&args[4],&mut grid,&mut dependences)
}
