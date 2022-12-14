use std::cmp::{max, min};

use itertools::Itertools;
use nom::bytes::complete::tag;
use nom::character::complete;
use nom::multi::separated_list1;
use nom::sequence::separated_pair;
use nom::IResult;

pub fn part1(input: &str) -> i32 {
    let lines = parse_input(input);
    let max_y = lines
        .iter()
        .flat_map(|line| match line {
            Line::Segment(a, b) => vec![a.y, b.y],
            Line::Horizontal(_) => panic!("should never happen"),
        })
        .max()
        .unwrap();
    let sand_spots = sand_pouring(&lines, max_y);
    sand_spots.len() as i32
}

pub fn part2(input: &str) -> i32 {
    let mut lines = parse_input(input);
    let mut max_y = lines
        .iter()
        .flat_map(|line| match line {
            Line::Segment(a, b) => vec![a.y, b.y],
            Line::Horizontal(_) => panic!("should never happen"),
        })
        .max()
        .unwrap();
    max_y += 2;
    lines.push(Line::Horizontal(max_y));
    let sand_spots = sand_pouring(&lines, max_y);
    sand_spots.len() as i32
}

fn sand_pouring(lines: &[Line], max_y: i32) -> Vec<Point> {
    let mut sand_spots: Vec<Point> = vec![];
    loop {
        let mut sand = Point { x: 500, y: 0 };
        while let Some(spot) = next_spot(lines, &sand_spots, &sand) {
            sand = spot;
            if sand.y >= max_y {
                break;
            }
        }
        if sand.y >= max_y {
            break;
        }
        sand_spots.push(sand);
        if sand == (Point { x: 500, y: 0 }) {
            break;
        }
    }
    sand_spots
}

fn next_spot(lines: &[Line], sand_spots: &[Point], sand: &Point) -> Option<Point> {
    let spots = [
        Point {
            x: sand.x,
            y: sand.y + 1,
        },
        Point {
            x: sand.x - 1,
            y: sand.y + 1,
        },
        Point {
            x: sand.x + 1,
            y: sand.y + 1,
        },
    ];
    spots
        .into_iter()
        .find(|spot| !(lines.iter().any(|line| line.contains(spot)) || sand_spots.contains(spot)))
}

#[derive(Debug, Copy, Clone, PartialEq)]
struct Point {
    x: i32,
    y: i32,
}

#[derive(Debug)]
enum Line {
    Horizontal(i32),
    Segment(Point, Point),
}

impl Line {
    fn contains(&self, point: &Point) -> bool {
        match self {
            Line::Horizontal(y) => point.y == *y,
            Line::Segment(a, b) => (a.x..=b.x).contains(&point.x) && (a.y..=b.y).contains(&point.y),
        }
    }
}

fn parse_input(input: &str) -> Vec<Line> {
    lines_of_rock(input)
        .unwrap()
        .1
        .iter()
        .flat_map(|line_of_rock| {
            line_of_rock.iter().tuple_windows().map(|(a, b)| {
                Line::Segment(
                    Point {
                        x: min(a.0, b.0),
                        y: min(a.1, b.1),
                    },
                    Point {
                        x: max(a.0, b.0),
                        y: max(a.1, b.1),
                    },
                )
            })
        })
        .collect::<Vec<_>>()
}

fn lines_of_rock(input: &str) -> IResult<&str, Vec<Vec<(i32, i32)>>> {
    separated_list1(complete::newline, line_of_rock)(input)
}

fn line_of_rock(input: &str) -> IResult<&str, Vec<(i32, i32)>> {
    separated_list1(
        tag(" -> "),
        separated_pair(complete::i32, tag(","), complete::i32),
    )(input)
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9";

    #[test]
    fn part1_works() {
        assert_eq!(part1(INPUT), 24);
    }

    #[test]
    fn part2_works() {
        assert_eq!(part2(INPUT), 93);
    }
}
