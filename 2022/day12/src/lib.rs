use nom::{
    character::complete::{newline, one_of},
    multi::{many1, separated_list1},
    IResult,
};
use petgraph::algo::dijkstra;
use petgraph::prelude::*;

pub fn part1(input: &str) -> i32 {
    let grid = parse_grid(input).unwrap().1;
    let num_nodes = grid.len() * grid[0].len();
    let start_node = grid
        .iter()
        .flatten()
        .enumerate()
        .find(|&(_i, &c)| c == 'S')
        .unwrap()
        .0;
    let end_node = grid
        .iter()
        .flatten()
        .enumerate()
        .find(|(_i, &c)| c == 'E')
        .unwrap()
        .0;
    let grid: Vec<Vec<i32>> = grid
        .iter()
        .map(|row| {
            row.iter()
                .map(|&c| match c {
                    'S' => 'a' as i32,
                    'E' => 'z' as i32,
                    other => other as i32,
                })
                .map(|i| i - 'a' as i32 + 1)
                .collect()
        })
        .collect();
    let mut graph: Graph<i32, i32> = Graph::new();
    for _ in 0..num_nodes {
        graph.add_node(0);
    }
    for y in 0..grid.len() {
        for x in 0..grid[0].len() {
            let from_node_value = grid[y][x];
            let from_node = y * grid[0].len() + x;
            for (dy, dx) in [(-1, 0), (0, -1), (1, 0), (0, 1)] {
                let ty = y as i32 + dy;
                let tx = x as i32 + dx;
                if ty < 0 || tx < 0 || ty >= grid.len() as i32 || tx >= grid[0].len() as i32 {
                    continue;
                }
                let to_node_value = grid[ty as usize][tx as usize] as i32;
                let to_node = ty as usize * grid[0].len() + tx as usize;
                let weight = to_node_value - from_node_value as i32;
                if weight <= 1 {
                    graph.add_edge(NodeIndex::new(from_node), NodeIndex::new(to_node), 1);
                }
            }
        }
    }
    *dijkstra(
        &graph,
        NodeIndex::new(start_node),
        Some(NodeIndex::new(end_node)),
        |_| 1,
    )
    .get(&NodeIndex::new(end_node))
    .unwrap()
}

pub fn part2(input: &str) -> i32 {
    let grid = parse_grid(input).unwrap().1;
    let num_nodes = grid.len() * grid[0].len();
    let end_node = grid
        .iter()
        .flatten()
        .enumerate()
        .find(|(_i, &c)| c == 'E')
        .unwrap()
        .0;
    let grid: Vec<Vec<i32>> = grid
        .iter()
        .map(|row| {
            row.iter()
                .map(|&c| match c {
                    'S' => 'a' as i32,
                    'E' => 'z' as i32,
                    other => other as i32,
                })
                .map(|i| i - 'a' as i32 + 1)
                .collect()
        })
        .collect();
    let mut graph: Graph<i32, i32> = Graph::new();
    for _ in 0..num_nodes {
        graph.add_node(0);
    }
    for y in 0..grid.len() {
        for x in 0..grid[0].len() {
            let from_node_value = grid[y][x];
            let from_node = y * grid[0].len() + x;
            for (dy, dx) in [(-1, 0), (0, -1), (1, 0), (0, 1)] {
                let ty = y as i32 + dy;
                let tx = x as i32 + dx;
                if ty < 0 || tx < 0 || ty >= grid.len() as i32 || tx >= grid[0].len() as i32 {
                    continue;
                }
                let to_node_value = grid[ty as usize][tx as usize] as i32;
                let to_node = ty as usize * grid[0].len() + tx as usize;
                let weight = to_node_value - from_node_value as i32;
                if weight <= 1 {
                    graph.add_edge(NodeIndex::new(from_node), NodeIndex::new(to_node), 1);
                }
            }
        }
    }
    let mut minimum = i32::MAX;
    for y in 0..grid.len() {
        for x in 0..grid[0].len() {
            let from_node_value = grid[y][x];
            if from_node_value != 1 {
                continue;
            }
            let from_node = y * grid[0].len() + x;
            if from_node == end_node {
                continue;
            }
            let m = *dijkstra(
                &graph,
                NodeIndex::new(from_node),
                Some(NodeIndex::new(end_node)),
                |_| 1,
            )
            .get(&NodeIndex::new(end_node))
            .unwrap_or(&i32::MAX);
            if m < minimum {
                minimum = m;
            }
        }
    }
    minimum
}

fn parse_grid(input: &str) -> IResult<&str, Vec<Vec<char>>> {
    separated_list1(newline, many1(one_of("abcdefghijklmnopqrstuvwxyzES")))(input)
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi";

    #[test]
    fn part1_works() {
        assert_eq!(part1(INPUT), 31);
    }

    #[test]
    fn part2_works() {
        assert_eq!(part2(INPUT), 29);
    }
}
