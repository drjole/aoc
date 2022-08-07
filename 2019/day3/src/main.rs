use std::collections::HashSet;
use std::io;
use std::io::BufRead;

fn main() {
    let stdin = io::stdin();
    let mut wires: Vec<Vec<String>> = vec![];
    for line in stdin.lock().lines() {
        wires.push(line.unwrap()
            .split(",")
            .map(|s| s.to_string())
            .collect());
    }
    let mut points: Vec<HashSet<(i32, i32)>> = vec![];
    for wire in &wires {
        points.push(HashSet::from_iter(walk(&wire)));
    }
    let intersections = points[0].intersection(&points[1]);
    let mut closest = i32::MAX;
    let mut fewest_steps = i32::MAX;
    for intersection in intersections {
        let d = dist(intersection.0, intersection.1, 0, 0);
        if d < closest {
            closest = d;
        }
        let mut steps = 0;
        for wire in &wires {
            for point in walk(&wire) {
                steps += 1;
                if steps >= fewest_steps || point == *intersection {
                    break;
                }
            }
        }
        if steps < fewest_steps {
            fewest_steps = steps;
        }
    }
    println!("{}", closest);
    println!("{}", fewest_steps);
}

fn walk(wire: &Vec<String>) -> Vec<(i32, i32)> {
    let mut x = 0;
    let mut y = 0;
    let mut result: Vec<(i32, i32)> = vec![];
    for instruction in wire {
        let direction = &instruction[..1];
        let steps = (&instruction[1..]).parse::<i32>().unwrap();
        for _ in 0..steps {
            match direction {
                "U" => y += 1,
                "R" => x += 1,
                "D" => y -= 1,
                "L" => x -= 1,
                _ => (),
            }
            result.push((x, y));
        }
    }
    result
}

fn dist(x1: i32, y1: i32, x2: i32, y2: i32) -> i32 {
    (y2 - y1).abs() + (x2 - x1).abs()
}
