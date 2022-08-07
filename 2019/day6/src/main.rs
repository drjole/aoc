use std::collections::HashMap;
use std::io;
use std::io::BufRead;

fn main() {
    let stdin = io::stdin();
    let mut orbits = HashMap::new();
    for line in stdin.lock().lines() {
        let l = line.unwrap();
        let objects: Vec<&str> = l.split(")").collect();
        orbits.insert(objects[1].to_string(), objects[0].to_string());
    }
    let mut num_orbits = 0;
    for object in orbits.keys() {
        num_orbits += parents(&orbits, object).len();
    }
    println!("{}", num_orbits);
    println!("{}", shortest_path(&orbits, "YOU", "SAN"));
}

fn shortest_path(orbits: &HashMap<String, String>, from: &str, to: &str) -> i32 {
    let from_parents = parents(&orbits, orbits.get(from).unwrap());
    let to_parents = parents(&orbits, orbits.get(to).unwrap());
    let mut c = 0;
    for (a, b) in from_parents.iter().zip(to_parents.iter()) {
        if a == b {
            c += 1;
            continue;
        }
    }
    (from_parents.len() as i32 - c) + (to_parents.len() as i32 - c)
}

fn parents(orbits: &HashMap<String, String>, object: &str) -> Vec<String> {
    let mut result = vec![];
    if object != "COM" {
        let parent = orbits.get(object).unwrap();
        result.extend(parents(&orbits, parent));
    }
    result.push(object.to_string());
    result
}
