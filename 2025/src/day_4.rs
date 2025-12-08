pub struct Day {
    pub data: String,
}

impl crate::day_util::DaySolution for Day {
    fn part1(&self) -> String {
      let mut movable_rolls= 0;
      let mut grid: Vec<Vec<char>> = Vec::new();
      for line in self.data.split('\n') {
        grid.push(line.chars().collect());
      }
      let height = grid.len();
      let width = grid[0].len();
      for y in 0..height {
        for x in 0..width {
          if grid[y][x] == '@' && count_neighbors(&grid, width, height, x, y) < 4 {
            grid[y][x] = 'x';
            movable_rolls += 1;
          }
        }
      }
      format!("Part 1: {movable_rolls}")
    }

    fn part2(&self) -> String {
      let mut grid: Vec<Vec<char>> = Vec::new();
      for line in self.data.split('\n') {
        grid.push(line.chars().collect());
      }
      let height = grid.len();
      let width = grid[0].len();
      let mut total_movable_rolls = 0;
      loop {
        let mut movable_rolls= 0;
        for y in 0..height {
          for x in 0..width {
            if grid[y][x] == '@' && count_neighbors(&grid, width, height, x, y) < 4 {
              grid[y][x] = 'x';
              movable_rolls += 1;
            }
          }
        }
        if movable_rolls == 0 {
          break;
        }
        total_movable_rolls += movable_rolls;
        for y in 0..height {
          for x in 0..width {
            if grid[y][x] == 'x' {
              grid[y][x] = '.';
            }
          }
        }
      }
      format!("Part 2: {total_movable_rolls}")
    }
}

fn count_neighbors(grid: &Vec<Vec<char>>, width: usize, height: usize, x: usize, y: usize) -> i32 {
    let mut neighbors = 0;
    for i in (if y > 0 { y - 1 } else {0})..(if y < height - 1 { y + 2 } else { height }) {
      for j in (if x > 0 { x - 1 } else {0})..(if x < width - 1 { x + 2 } else { width }) {
        if grid[i][j] == '@' || grid[i][j] == 'x' {
          neighbors += 1;
        }
      }
    }
    neighbors - 1  // Remove the center, i.e. itself.
}