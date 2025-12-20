use std::collections::BTreeSet;

pub struct Day {
    pub data: String,
}

impl crate::day_util::DaySolution for Day {
    fn part1(&self) -> String {
      let mut fresh= 0;
      let mut ranges: Vec<Range> = Vec::new();
      let (range_descs, ingredients) = self.data.split_once("\n\n").expect("should have 2 pieces");
      for desc in range_descs.split('\n') {
        if desc == "" {
          continue;
        }
        ranges.push(Range::new_from_desc(desc));
      }
      for ingredient in ingredients.split('\n') {
        if ingredient == "" {
          continue;
        }
        let id= ingredient.parse().unwrap();
        for range in &ranges {
          if range.contains(id) {
            fresh += 1;
            break;
          }
        }
      }
      format!("Part 1: {fresh}")
    }

    fn part2(&self) -> String {
      let mut ranges: BTreeSet<Range> = BTreeSet::new();
      'parser: for desc in self.data.split('\n') {
        if !desc.contains("-"){
          continue 'parser;
        }
        ranges.insert(Range::new_from_desc(desc));
      }
      'merge: loop {
        let mut new_ranges: BTreeSet<Range> = BTreeSet::new();
        let mut range_iter = ranges.clone().into_iter();
        let mut curr = range_iter.next().unwrap();
        new_ranges.insert(curr);
        for next in range_iter {
          let combined = curr.combine(&next);
          if combined.is_some() {
            new_ranges.remove(&curr);
            new_ranges.insert(combined.unwrap());
            curr = combined.unwrap();
          } else {
            new_ranges.insert(next);
            curr = next;
          }
        }
        if new_ranges == ranges {
          break 'merge;
        }
        ranges = new_ranges;
      }
      let mut fresh= 0;
      for range in &ranges {
        let change = range.end - range.start + 1;
        fresh += change;
      }
      format!("Part 2: {fresh}")
    }
}

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord, Copy, Clone)]
struct Range {
  start: u64,
  end: u64,
}

impl Range {
  pub fn new_from_desc(desc: &str) -> Range {
    let (start_str, end_str) = desc.split_once('-').expect("2 pieces");
    Range::new(start_str.parse().unwrap(), end_str.parse().unwrap())
  }

  pub fn new(start: u64, end: u64) -> Range {
    Range {
      start,
      end,
    }
  }

  pub fn contains(&self, val: u64) -> bool {
    self.start <= val && val <= self.end
  }

  pub fn combine(&self, range: &Range) -> Option<Range> {
    if self != range && (self.contains(range.start) || self.contains(range.end)) {
      Some(Range::new(self.start.min(range.start), self.end.max(range.end)))
    } else {
      None
    }
  }
}