struct Foo {
    pub n: u32,
}

impl Foo {
    fn new () -> Foo {
        Foo {n: 24}
    }
}


fn main() {
    let foo = Foo::new();
    println!("{}", foo.n)
}
