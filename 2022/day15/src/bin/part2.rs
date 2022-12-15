use std::fs;

use day15::part2;

fn main() {
    let input = fs::read_to_string("input.txt").unwrap();
    let max_xy = 4_000_000;
    println!("{}", part2(&input, max_xy));
}
