use std ::thread;
use std::time::Duration;

fn main () {
    for i in 1..5 {
        thread::spawn(move || {
            println!("thread: {}", i);
            thread::sleep(Duration::from_millis(3));
        });
    }

    for i in 1..5 {
        println!("main: {}", i);
        thread::sleep(Duration::from_millis(3));
    }
}
