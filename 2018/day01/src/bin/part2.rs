use std::fs;

use day01::part2;

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();
    println!("{}", part2(&input));
}
