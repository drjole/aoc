pub fn part1(input: &str) -> i32 {
    let trees = parse_trees(input);
    let n_rows = trees.len();
    let n_cols = trees[0].len();
    let inside_trees = trees_with_coords(&trees)
        .filter(|(x, y, tree)| {
            let neighbours = prepare_neighbours(&trees, *x, *y);
            let mut blocked = vec![false, false, false, false];
            for (i, other_trees) in neighbours.iter().enumerate() {
                for &other_tree in other_trees {
                    if (other_tree as u32) >= *tree {
                        blocked[i] = true;
                        break;
                    }
                }
            }
            blocked.iter().any(|&x| !x)
        })
        .count();
    let border_trees = 2 * (n_rows + n_cols - 2);
    border_trees as i32 + inside_trees as i32
}

pub fn part2(input: &str) -> i32 {
    let trees = parse_trees(input);
    trees_with_coords(&trees)
        .map(|(x, y, tree)| {
            let neighbours = prepare_neighbours(&trees, x, y);
            let mut scenic_score = 1;
            for other_trees in neighbours.iter() {
                let mut distance = 0;
                for &other_tree in other_trees {
                    distance += 1;
                    if (other_tree as u32) >= tree {
                        break;
                    }
                }
                scenic_score *= distance;
            }
            scenic_score
        })
        .max()
        .unwrap()
}

fn parse_trees(input: &str) -> Vec<Vec<u32>> {
    input
        .lines()
        .map(|line| {
            line.chars()
                .map(|c| c.to_digit(10).unwrap())
                .collect::<Vec<_>>()
        })
        .collect::<Vec<_>>()
}

fn trees_with_coords(trees: &Vec<Vec<u32>>) -> impl Iterator<Item = (usize, usize, u32)> + '_ {
    trees
        .iter()
        .enumerate()
        .skip(1)
        .take(trees.len() - 2)
        .flat_map(|(y, row)| {
            row.iter()
                .enumerate()
                .skip(1)
                .take(row.len() - 2)
                .map(move |(x, &tree)| (x, y, tree))
        })
}

fn prepare_neighbours(trees: &Vec<Vec<u32>>, x: usize, y: usize) -> Vec<Vec<u32>> {
    let n_rows = trees.len();
    let n_cols = trees[0].len();
    vec![
        (0..x)
            .rev()
            .into_iter()
            .map(|x2| trees[y][x2])
            .collect::<Vec<_>>(),
        (x + 1..n_cols)
            .into_iter()
            .map(|x2| trees[y][x2])
            .collect::<Vec<_>>(),
        (0..y)
            .rev()
            .into_iter()
            .map(|y2| trees[y2][x])
            .collect::<Vec<_>>(),
        (y + 1..n_rows)
            .into_iter()
            .map(|y2| trees[y2][x])
            .collect::<Vec<_>>(),
    ]
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "30373
25512
65332
33549
35390";

    #[test]
    fn part1_works() {
        assert_eq!(part1(INPUT), 21);
    }

    #[test]
    fn part2_works() {
        assert_eq!(part2(INPUT), 8);
    }
}
