use std::io;
use std::io::BufRead;

fn main() {
    let stdin = io::stdin();
    let mut first = 0;
    let mut second = 0;
    for line in stdin.lock().lines() {
        let mut fuel = line.unwrap().parse::<i32>().unwrap() / 3 - 2;
        first += fuel;
        while fuel > 0 {
            second += fuel;
            fuel = fuel / 3 - 2;
        }
    }
    println!("{}", first);
    println!("{}", second);
}
