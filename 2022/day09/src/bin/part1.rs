use std::fs;

use day09::part1;

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();
    println!("{}", part1(&input));
}
