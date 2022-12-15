use std::fs;

use day15::part1;

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();
    let y = 2_000_000;
    println!("{}", part1(&input, y));
}
