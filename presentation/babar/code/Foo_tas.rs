struct Foo {
    pub v: Vec<u32>,
}

impl Foo {
    fn new () -> Foo {
        Foo {v:Vec::new()}
    }
}


fn main() {
    let mut foo = Foo::new();
    foo.v.push(24);
    println!("{}", foo.v[0])
}
