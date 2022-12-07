use std::collections::HashMap;
use std::path::PathBuf;

pub fn part1(input: &str) -> i32 {
    let fs = parse_file_system(input);
    fs.directories
        .keys()
        .map(|directory| fs.size(directory))
        .filter(|&size| size <= 100000)
        .sum::<i32>()
}

pub fn part2(input: &str) -> i32 {
    let fs = parse_file_system(input);
    let available_space = 70000000;
    let required_space = 30000000;
    let used_space = fs.size("/");
    let minimum_directory_size = required_space - (available_space - used_space);
    fs.directories
        .keys()
        .map(|directory| fs.size(directory))
        .filter(|&size| size >= minimum_directory_size)
        .min()
        .unwrap()
}

fn parse_file_system(input: &str) -> FileSystem {
    let lines = input.lines().collect::<Vec<_>>();
    let mut files = HashMap::new();
    let mut directories = HashMap::new();
    let mut current_directory = PathBuf::new();
    input
        .lines()
        .enumerate()
        .filter(|(_i, line)| line.starts_with('$'))
        .for_each(|(i, line)| {
            match &line[2..4] {
                "cd" => {
                    current_directory = match &line[5..] {
                        "/" => PathBuf::from("/"),
                        ".." => current_directory.parent().unwrap().to_path_buf(),
                        other => current_directory.join(other),
                    };
                }
                "ls" => {
                    let children = lines
                        .iter()
                        .skip(i + 1)
                        .take_while(|line| !line.starts_with('$'))
                        .map(|f| {
                            let (dir_or_size, name) = f.split_once(' ').unwrap();
                            let full_name =
                                current_directory.join(name).to_str().unwrap().to_string();
                            if dir_or_size != "dir" {
                                files
                                    .insert(full_name.clone(), dir_or_size.parse::<i32>().unwrap());
                            }
                            full_name
                        })
                        .collect::<Vec<_>>();
                    directories.insert(current_directory.to_str().unwrap().to_string(), children);
                }
                other => panic!("Invalid command: {}", other),
            };
        });
    FileSystem { directories, files }
}

struct FileSystem {
    directories: HashMap<String, Vec<String>>,
    files: HashMap<String, i32>,
}

impl FileSystem {
    fn size(&self, path: &str) -> i32 {
        if self.directories.contains_key(path) {
            self.directories
                .get(path)
                .unwrap()
                .iter()
                .map(|child| self.size(child))
                .sum::<i32>()
        } else {
            *self
                .files
                .get(path)
                .unwrap_or_else(|| panic!("Could not find file: {}", path))
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k";

    #[test]
    fn part1_works() {
        assert_eq!(part1(INPUT), 95437);
    }

    #[test]
    fn part2_works() {
        assert_eq!(part2(INPUT), 24933642);
    }
}
