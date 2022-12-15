use std::collections::BTreeMap;
use std::ops::RangeInclusive;

use itertools::Itertools;
use nom::bytes::complete::tag;
use nom::character::complete;
use nom::multi::separated_list1;
use nom::sequence::{preceded, separated_pair};
use nom::IResult;

pub fn part1(input: &str, y: i64) -> i64 {
    point_pairs(input)
        .unwrap()
        .1
        .into_iter()
        .map(|(sensor, beacon)| (sensor, sensor.distance(&beacon)))
        .filter(|(sensor, distance)| ((sensor.y - distance)..(sensor.y + distance)).contains(&y))
        .flat_map(|(sensor, distance)| {
            let d = distance - (sensor.y - y).abs();
            (sensor.x - d)..(sensor.x + d)
        })
        .unique()
        .count() as i64
}

pub fn part2(input: &str, max_xy: i64) -> i64 {
    let mut all_ranges: BTreeMap<i64, Vec<RangeInclusive<_>>> = BTreeMap::new();
    point_pairs(input)
        .unwrap()
        .1
        .into_iter()
        .map(|(sensor, beacon)| (sensor, sensor.distance(&beacon)))
        .flat_map(|(sensor, max_distance)| {
            ((sensor.y - max_distance).max(0)..(sensor.y + max_distance).min(max_xy)).map(
                move |y| {
                    let d = max_distance - (sensor.y - y).abs();
                    (y, (sensor.x - d).max(0)..=(sensor.x + d).min(max_xy))
                },
            )
        })
        .filter(|(y, _)| &0 <= y && y <= &max_xy)
        .for_each(|(y, range)| {
            all_ranges
                .entry(y)
                .and_modify(|entry| {
                    entry.push(range.clone());
                })
                .or_insert_with(|| vec![range]);
        });
    let (beacon_x, beacon_y) = all_ranges
        .into_iter()
        .find_map(|(y, mut ranges)| {
            ranges.sort_by(|a, b| a.start().cmp(b.start()));
            let result: (RangeInclusive<i64>, Option<i64>) =
                ranges.iter().fold((0..=0, None), |mut acc, range| {
                    if acc.1.is_some() {
                        return acc;
                    }
                    if acc.0.end() >= range.start() {
                        acc.0 = *acc.0.start()..=(*acc.0.end().max(range.end()));
                    } else {
                        acc.1 = Some(acc.0.end() + 1);
                    }
                    acc
                });
            result.1.map(|x| (x, y))
        })
        .unwrap();
    beacon_x * 4_000_000 + beacon_y
}

fn point_pairs(input: &str) -> IResult<&str, Vec<(Point, Point)>> {
    separated_list1(complete::newline, point_pair)(input)
}

fn point_pair(input: &str) -> IResult<&str, (Point, Point)> {
    separated_pair(sensor, tag(": "), beacon)(input)
}

fn sensor(input: &str) -> IResult<&str, Point> {
    preceded(tag("Sensor at "), point)(input)
}

fn beacon(input: &str) -> IResult<&str, Point> {
    preceded(tag("closest beacon is at "), point)(input)
}

fn point(input: &str) -> IResult<&str, Point> {
    let (input, coordinates) = separated_pair(
        preceded(tag("x="), complete::i64),
        tag(", "),
        preceded(tag("y="), complete::i64),
    )(input)?;
    Ok((
        input,
        Point {
            x: coordinates.0,
            y: coordinates.1,
        },
    ))
}

#[derive(Debug, Copy, Clone, Eq, PartialEq, Ord, PartialOrd, Hash)]
struct Point {
    x: i64,
    y: i64,
}

impl Point {
    fn distance(&self, other: &Point) -> i64 {
        (self.x - other.x).abs() + (self.y - other.y).abs()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3";

    #[test]
    fn part1_works() {
        let y = 10;
        assert_eq!(part1(INPUT, y), 26);
    }

    #[test]
    fn part2_works() {
        assert_eq!(part2(INPUT, 20), 56000011);
    }
}
