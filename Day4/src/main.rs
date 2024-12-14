//PART 1
//use std::fs;
//use std::io;
//
//fn read_grid_from_file(filename: &str) -> Result<Vec<Vec<char>>, io::Error> {
//    let content = fs::read_to_string(filename)?;
//    let grid: Vec<Vec<char>> = content
//        .lines()
//        .map(|line| line.chars().collect())
//        .collect();
//    Ok(grid)
//}
//
//fn count_word(grid: &[Vec<char>], word: &str) -> usize{
//    let word_chars: Vec<char> = word.chars().collect();
//    let word_length = word_chars.len();
//    let rows = grid.len();
//    let cols = grid[0].len();
//
//    let directions = [
//        (0, 1),
//        (0, -1),
//        (1, 0),
//        (-1, 0),
//        (1, 1),
//        (1, -1),
//        (-1, 1),
//        (-1, -1)
//    ];
//
//    let mut count = 0;
//
//    for r in 0..rows {
//        for c in 0..cols {
//            for &(dr, dc) in &directions {
//                let mut match_found = true;
//                for i in 0..word_length {
//                    let nr = r as isize + dr * i as isize;
//                    let nc = c as isize + dc * i as isize;
//                    if nr < 0 || nr >= rows as isize || nc < 0 || nc >= cols as isize {
//                        match_found = false;
//                        break;
//                    }
//                    if grid[nr as usize][nc as usize] != word_chars[i] {
//                        match_found = false;
//                        break;
//                    }
//                }
//                if match_found {
//                    count += 1
//                }
//            }
//        }
//    }
//    count
//}
//
//fn main() -> Result<(), io::Error> {
//    let grid = read_grid_from_file('input.txt')?;
//    
//    let word = 'XMAS';
//    let result = count_word(&grid, word);
//    println!('Total occurrences of '{}': {}', word, result);
//
//    Ok(())
//}


//PART 2
use std::fs;

fn main() {
    let input = fs::read_to_string("input.txt").expect("Failed to read input.txt");
    let grid: Vec<Vec<char>> = input.lines().map(|line| line.chars().collect()).collect();
    let rows = grid.len();
    let cols = grid[0].len();
    
    let mut count = 0; 

    for r in 1..rows - 1 {
        for c in 1..cols - 1 {
            if grid[r][c] != 'A' {
                continue;
            }

            if (grid[r - 1][c - 1] == 'M' && grid[r + 1][c + 1] == 'S'
                || grid[r - 1][c - 1] == 'S' && grid[r + 1][c + 1] == 'M')
            && (grid[r - 1][c + 1] == 'M' && grid[r + 1][c - 1] == 'S'
                || grid[r - 1][c + 1] == 'S' && grid[r + 1][c - 1] == 'M') {
                count += 1;
            }
        }
    }

    println!(count);
}

