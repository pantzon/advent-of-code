pub struct Day {
    pub data: String,
}

impl crate::day_util::DaySolution for Day {
    fn part1(&self) -> String {
      let mut result: u16 = 0;
      let mut pos = 50;
      let instructions: Vec<&str> = self.data.split('\n').collect();
      for inst in instructions {
        if inst.len() == 0 {
          continue;
        }
        let dir = inst.chars().nth(0).expect("has an initial character");
        let count = inst[1..].parse::<i32>().unwrap();
        pos += count * if dir == 'L' { -1 } else { 1 };
        pos %= 100;
        if pos == 0 {
          result += 1;
        }
      }
      format!("Part 1: {result}")
    }

    fn part2(&self) -> String {
      let mut result = 0;
      let mut pos = 50;
      let instructions: Vec<&str> = self.data.split('\n').collect();
      for inst in instructions {
        if inst.len() == 0 {
          continue;
        }
        let dir = inst.chars().nth(0).expect("has an initial character");
        let count = inst[1..].parse::<i32>().unwrap();
        let orig_pos = pos;
        result += count / 100;
        pos += (count % 100) * if dir == 'L' { -1 } else { 1 };
        if orig_pos != 0 && pos <= 0 || pos >= 100 {
          result += 1;
        }
        pos = (100 + pos) % 100;
      }
      format!("Part 2: {result}")
    }
}