use std::cmp::Ordering;

use nom::branch::alt;
use nom::bytes::complete::tag;
use nom::character::complete;
use nom::character::complete::{char, newline};
use nom::combinator::map;
use nom::multi::{count, separated_list0, separated_list1};
use nom::sequence::{delimited, separated_pair};
use nom::IResult;

pub fn part1(input: &str) -> i32 {
    packet_pairs(input)
        .unwrap()
        .1
        .iter()
        .enumerate()
        .filter_map(|(i, (left, right))| match left.cmp(right) {
            Ordering::Less => Some(i as i32 + 1),
            Ordering::Greater => None,
            Ordering::Equal => panic!("should never happen"),
        })
        .sum::<i32>()
}

pub fn part2(input: &str) -> i32 {
    let mut packets: Vec<Packet> = packet_pairs(input)
        .unwrap()
        .1
        .into_iter()
        .flat_map(|(l, r)| vec![l, r])
        .collect::<Vec<_>>();
    let divider1 = Packet::List(vec![Packet::List(vec![Packet::Number(6)])]);
    let divider2 = Packet::List(vec![Packet::List(vec![Packet::Number(2)])]);
    packets.push(divider1.clone());
    packets.push(divider2.clone());
    packets.sort();
    packets
        .iter()
        .enumerate()
        .filter_map(|(i, packet)| {
            if packet == &divider1 || packet == &divider2 {
                Some(i as i32 + 1)
            } else {
                None
            }
        })
        .product()
}

#[derive(Debug, Clone, Eq)]
enum Packet {
    List(Vec<Packet>),
    Number(i32),
}

impl PartialEq for Packet {
    fn eq(&self, other: &Self) -> bool {
        match (self, other) {
            (Self::List(l), Self::List(r)) => l == r,
            (Self::Number(l), Self::Number(r)) => l == r,
            (Self::List(l), Self::Number(r)) => l == &vec![Self::Number(*r)],
            (Self::Number(l), Self::List(r)) => &vec![Self::Number(*l)] == r,
        }
    }
}

impl Ord for Packet {
    fn cmp(&self, other: &Self) -> Ordering {
        match (self, other) {
            (Self::List(l), Self::List(r)) => l.cmp(r),
            (Self::Number(l), Self::Number(r)) => l.cmp(r),
            (Self::List(l), Self::Number(r)) => l.cmp(&vec![Packet::Number(*r)]),
            (Self::Number(l), Self::List(r)) => vec![Packet::Number(*l)].cmp(r),
        }
    }
}

impl PartialOrd for Packet {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

fn packet(input: &str) -> IResult<&str, Packet> {
    let (input, packets) = delimited(
        tag("["),
        separated_list0(char(','), alt((map(complete::i32, Packet::Number), packet))),
        tag("]"),
    )(input)?;
    Ok((input, Packet::List(packets)))
}

fn packet_pairs(input: &str) -> IResult<&str, Vec<(Packet, Packet)>> {
    separated_list1(count(newline, 2), separated_pair(packet, newline, packet))(input)
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]";

    #[test]
    fn part1_works() {
        assert_eq!(part1(INPUT), 13);
    }

    #[test]
    fn part2_works() {
        assert_eq!(part2(INPUT), 140);
    }
}
