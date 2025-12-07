use fancy_regex::Regex;

pub struct Day {
    pub data: String,
}

impl crate::day_util::DaySolution for Day {
    fn part1(&self) -> String {
      let mut result: u64 = 0;
      for range in self.data.split(',') {
        let b_and_e: Vec<u64> = range.split('-').map(|i| i.parse::<u64>().unwrap()).collect();
        if b_and_e.len() != 2 {
          panic!("got data with too many pieces: {range}")
        }
        for id in b_and_e[0]..b_and_e[1] {
          if self.p1_value_checker(id) {
            result += id;
          }
        }

      }
      format!("Part 1: {result}")
    }

    fn part2(&self) -> String {
      let mut result: u64 = 0;
      for range in self.data.split(',') {
        let b_and_e: Vec<u64> = range.split('-').map(|i| i.parse::<u64>().unwrap()).collect();
        if b_and_e.len() != 2 {
          panic!("got data with too many pieces: {range}")
        }
        for id in b_and_e[0]..b_and_e[1] {
          if self.p2_value_checker(id) { 
            result += id;
          }
        }
      }
      format!("Part 2: {result}")
    }
}

impl Day {
    fn p1_value_checker(&self, v: u64) -> bool {
      let splitter = 10_u64.pow((v.ilog10() + 1) / 2);
      v / splitter == v % splitter
    } 

    fn p2_value_checker(&self, v: u64) -> bool {
      let matcher = Regex::new(r"^(.*)\1+$").expect("regex compiles");
      matcher.is_match(&format!("{v}")).expect("matcher failed")
    } 
}