pub struct Day {
    pub data: String,
}

impl crate::day_util::DaySolution for Day {
    fn part1(&self) -> String {
      let mut jolts = 0;
      for bank in self.data.split('\n') {
        if bank.len() == 0 {
          continue;
        }
        let nums: Vec<u8> = bank.chars().map(|c| c.to_digit(10).expect("should be a digit") as u8).collect();
        jolts += recurse_sub_max(nums, 1);
      }
      format!("Part 1: {jolts}")
    }

    fn part2(&self) -> String {
      let mut jolts = 0;
      for bank in self.data.split('\n') {
        if bank.len() == 0 {
          continue;
        }
        let nums: Vec<u8> = bank.chars().map(|c| c.to_digit(10).expect("should be a digit") as u8).collect();
        jolts += recurse_sub_max(nums, 11)
      }
      format!("Part 2: {jolts}")
    }
}

fn recurse_sub_max(nums: Vec<u8>, depth_left: usize) -> u64{
  let mut sub_max = 0;
  let mut max_i = 0;
  for i in 0..(nums.len() - depth_left) {
    if i == 0 || nums[i] > nums[max_i] {
      max_i = i;
      if depth_left != 0 {
        sub_max = recurse_sub_max(nums[i + 1..].to_vec(), depth_left - 1);
      }
    }
  }
  let res = (nums[max_i] as u64) * 10_u64.pow(depth_left as u32) + sub_max;
  res
}