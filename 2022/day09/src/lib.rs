use std::collections::HashSet;

pub fn part1(input: &str) -> i32 {
    process(input, 2)
}

pub fn part2(input: &str) -> i32 {
    process(input, 10)
}

fn process(input: &str, n: usize) -> i32 {
    let mut knots = vec![Vector { x: 0, y: 0 }; n];
    let mut tail_positions: HashSet<Vector> = HashSet::from([Vector { x: 0, y: 0 }]);
    get_directions(input).iter().for_each(|vector| {
        knots[0].move_by(vector);
        for i in 1..knots.len() {
            if !knots[i].is_neighbour_of(&knots[i - 1]) {
                let (dx, dy) = (
                    (knots[i - 1].x - knots[i].x).signum(),
                    (knots[i - 1].y - knots[i].y).signum(),
                );
                knots[i].move_by(&Vector { x: dx, y: dy });
                if i == knots.len() - 1 {
                    tail_positions.insert(*knots.last().unwrap());
                }
            }
        }
    });
    tail_positions.len() as i32
}

fn get_directions(input: &str) -> Vec<Vector> {
    input
        .lines()
        .flat_map(|line| {
            let (direction, n) = line.split_once(' ').unwrap();
            let n = n.parse::<i32>().unwrap();
            (0..n)
                .into_iter()
                .map(|_| match direction {
                    "R" => Vector { x: 1, y: 0 },
                    "L" => Vector { x: -1, y: 0 },
                    "U" => Vector { x: 0, y: 1 },
                    "D" => Vector { x: 0, y: -1 },
                    other => panic!("Invalid direction: {}", other),
                })
                .collect::<Vec<Vector>>()
        })
        .collect()
}

#[derive(Debug, Eq, Hash, PartialEq, Copy, Clone)]
struct Vector {
    x: i32,
    y: i32,
}

impl Vector {
    fn move_by(&mut self, vector: &Vector) {
        self.x += vector.x;
        self.y += vector.y;
    }

    fn is_neighbour_of(&self, other: &Vector) -> bool {
        (self.x - other.x).abs() <= 1 && (self.y - other.y).abs() <= 1
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2";

    const INPUT2: &str = "R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20";

    #[test]
    fn part1_works() {
        assert_eq!(part1(INPUT), 13);
    }

    #[test]
    fn part2_works() {
        assert_eq!(part2(INPUT), 1);
        assert_eq!(part2(INPUT2), 36);
    }
}
