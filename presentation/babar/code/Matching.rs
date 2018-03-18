fn main() {
    let n = 7;

    match n {
        1 => println!("Un"),
        2 | 4 => println!("Pair <5"),
        5...7 => println!(">4"),
        _ =>println!("3 ou >7 ,ou <1"),
    }
}
