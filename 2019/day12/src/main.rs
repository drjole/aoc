use std::io;
use std::io::Read;

#[derive(Debug, Copy, Clone)]
struct Moon {
    x: i64,
    y: i64,
    z: i64,
    dx: i64,
    dy: i64,
    dz: i64,
}

fn main() {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input).expect("Failed to read input");

    let mut moons: Vec<Moon> = Vec::with_capacity(input.lines().count());
    for line in input.lines() {
        let positions = line.replace("<", "").replace(">", "").replace(", ", "");
        let x_pos = positions.find("x=").expect("Did not find x=");
        let y_pos = positions.find("y=").expect("Did not find y=");
        let z_pos = positions.find("z=").expect("Did not find z=");
        let x = positions[(x_pos + 2)..y_pos].parse::<i64>().expect("Could not parse x");
        let y = positions[(y_pos + 2)..z_pos].parse::<i64>().expect("Could not parse y");
        let z = positions[(z_pos + 2)..].parse::<i64>().expect("Could not parse z");
        let moon = Moon { x, y, z, dx: 0, dy: 0, dz: 0 };
        moons.push(moon);
    }
    let initial_moons = moons.clone();

    let num_moons = moons.len();
    let mut time_step = 0;
    let mut x_freq: i64 = -1;
    let mut y_freq: i64 = -1;
    let mut z_freq: i64 = -1;
    loop {
        if time_step == 1000 {
            let mut total = 0;
            for i in 0..num_moons {
                let pot = moons[i].x.abs() + moons[i].y.abs() + moons[i].z.abs();
                let kin = moons[i].dx.abs() + moons[i].dy.abs() + moons[i].dz.abs();
                total += pot * kin;
            }
            println!("{}", total);
        }
        if x_freq != -1 && y_freq != -1 && z_freq != -1 {
            let l = lcm(lcm(x_freq, y_freq), z_freq);
            println!("{}", l);
            break;
        }
        if time_step > 0 {
            if x_freq == -1 {
                let mut fail = false;
                for i in 0..num_moons {
                    if moons[i].x != initial_moons[i].x || moons[i].dx != initial_moons[i].dx {
                        fail = true;
                        break;
                    }
                }
                if !fail {
                    x_freq = time_step;
                }
            }
            if y_freq == -1 {
                let mut fail = false;
                for i in 0..num_moons {
                    if moons[i].y != initial_moons[i].y || moons[i].dy != initial_moons[i].dy {
                        fail = true;
                        break;
                    }
                }
                if !fail {
                    y_freq = time_step;
                }
            }
            if z_freq == -1 {
                let mut fail = false;
                for i in 0..num_moons {
                    if moons[i].z != initial_moons[i].z || moons[i].dz != initial_moons[i].dz {
                        fail = true;
                        break;
                    }
                }
                if !fail {
                    z_freq = time_step;
                }
            }
        }
        for i in 0..num_moons {
            for j in (i + 1)..num_moons {
                if moons[i].x > moons[j].x {
                    moons[i].dx -= 1;
                    moons[j].dx += 1;
                } else if moons[i].x < moons[j].x {
                    moons[i].dx += 1;
                    moons[j].dx -= 1;
                }

                if moons[i].y > moons[j].y {
                    moons[i].dy -= 1;
                    moons[j].dy += 1;
                } else if moons[i].y < moons[j].y {
                    moons[i].dy += 1;
                    moons[j].dy -= 1;
                }

                if moons[i].z > moons[j].z {
                    moons[i].dz -= 1;
                    moons[j].dz += 1;
                } else if moons[i].z < moons[j].z {
                    moons[i].dz += 1;
                    moons[j].dz -= 1;
                }
            }
        }
        for i in 0..num_moons {
            moons[i].x += moons[i].dx;
            moons[i].y += moons[i].dy;
            moons[i].z += moons[i].dz;
        }
        time_step += 1;
    }
}

fn lcm(first: i64, second: i64) -> i64 {
    first * second / gcd(first, second)
}

fn gcd(first: i64, second: i64) -> i64 {
    let mut max = first;
    let mut min = second;
    if min > max {
        let val = max;
        max = min;
        min = val;
    }

    loop {
        let res = max % min;
        if res == 0 {
            return min;
        }

        max = min;
        min = res;
    }
}
