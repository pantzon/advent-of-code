use std::collections::VecDeque;

pub struct Day {
    pub data: String,
}

impl crate::day_util::DaySolution for Day {
    fn part1(&self) -> String {
      let mut lines: Vec<_> = self.data.split("\n").filter_map(|line| if line == "" { None } else { Some(line.split(" ").filter(|n| *n != "")) }).collect();
      let ops: Vec<&str> = lines.pop().unwrap().collect();
      let mut vals: VecDeque<u64> = VecDeque::new();
      for line in lines {
        let mut ops_iter = ops.iter();
        for num in line.filter_map(|s| s.parse::<u64>().ok()) {
          let op = (*ops_iter.next().unwrap()).chars().nth(0).unwrap();
          if vals.len() != ops.len() {
            vals.push_back(num);
          } else {
            let new_val = agg_value(vals.pop_front().unwrap(), num, op);
            vals.push_back(new_val);
          }
        }
      }
      let sum: u64 = vals.into_iter().reduce(|acc, v| acc + v).unwrap();
      format!("Part 1: {sum}")
    }

    fn part2(&self) -> String {
      let mut lines: Vec<_> = self.data.split("\n").filter_map(|line| if line == "" { None } else { Some(line.as_bytes()) }).collect();
      let ops = lines.pop().unwrap();
      let mut sum = 0_u64;
      let mut curr_vals: Vec<u64> = Vec::new();
      for i in (0..ops.len()).rev() {
        let mut curr = 0_u64;
        let mut has_val = false;
        for l in lines.iter() {
          let chr = (*l)[i] as char;
          if chr != ' ' {
            has_val = true;
            curr = curr * 10 + (chr.to_digit(10).unwrap() as u64);
          }
        }
        if has_val {
          curr_vals.push(curr);
        }
        let op = ops[i] as char;
        if op != ' ' {
          let mut result = if op == '+' { 0_u64 } else { 1_u64 };
          for v in curr_vals.iter() {
            result = agg_value(result, *v, op)
          }
          sum += result;
          curr_vals.clear();
        }
      }
      format!("Part 2: {sum}")
    }
}

fn agg_value(v1: u64, v2: u64, op: char) -> u64 {
  if op == '+' {
    v1 + v2
  } else if op == '*' {
    v1 * v2
  } else {
    panic!("Expected + or *, got {op}")
  }
}