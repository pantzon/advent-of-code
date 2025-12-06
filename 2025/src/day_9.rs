pub struct Day {
    pub data: String,
}

impl crate::day_util::DaySolution for Day {
    fn part1(&self) -> String {
      let len = self.data.len();
      format!("Part 1: {len}")
    }

    fn part2(&self) -> String {
      let len = self.data.len();
      format!("Part 2: {len}")
    }
}