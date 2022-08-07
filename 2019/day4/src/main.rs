use std::io;
use std::io::BufRead;

fn main() {
    let stdin = io::stdin();
    let range = stdin.lock().lines().next().unwrap().unwrap();
    let split: Vec<&str> = range.split("-").collect();
    let from = split[0].parse::<i32>().unwrap();
    let to = split[1].parse::<i32>().unwrap();
    let mut valid_passwords = 0;
    let mut valid_passwords_without_large_groups = 0;
    for password in from..=to {
        let password_string = password.to_string();
        let password_bytes = password_string.as_bytes();
        let mut monotonic = true;
        let mut contains_double = false;
        let mut consecutive = 1;
        let mut exactly_two = false;
        for i in 0..5 {
            if password_bytes[i] > password_bytes[i + 1] {
                monotonic = false;
                break;
            }
            if password_bytes[i] == password_bytes[i + 1] {
                contains_double = true;
                consecutive += 1;
            } else {
                if consecutive == 2 {
                    exactly_two = true;
                }
                consecutive = 1;
            }
        }
        if monotonic {
            if contains_double {
                valid_passwords += 1;
            } else if exactly_two || consecutive == 2 {
                valid_passwords_without_large_groups += 1;
            }
        }
    }
    println!("{}", valid_passwords);
    println!("{}", valid_passwords_without_large_groups);
}
