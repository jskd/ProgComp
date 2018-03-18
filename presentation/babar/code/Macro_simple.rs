macro_rules! sumIntArr {
    (
        $(
            $x:expr; [ $( $y:expr ),* ]
        );*
    ) => {
        &[ $($( $x + $y ),*),* ]
    }
}

fn main() {
    let a: &[i32]
        = sumIntArr!(10; [1, 2, 3];
               20; [4, 5, 6]);
    println!("a[0]: {}, a[1]: {}, a[2]: {}, a[3]: {}, a[4]: {}, a[5]: {}",
                    a[0], a[1], a[2], a[3], a[4], a[5]);
}
