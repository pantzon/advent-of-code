import { readFileSync } from "fs";

const data = readFileSync("./inputs/day1.txt", "utf-8");

function part1() {
  let count = 0;
  for (const i of data) {
    if (i === "(") {
      count++;
    } else if (i === ")") {
      count--;
    }
  }

  console.log(`Floor ${count}`);
}

function part2() {
  let count = 0;
  let position = 1;
  for (const i of data) {
    if (i === "(") {
      count++;
    } else if (i === ")") {
      count--;
    }
    if (count < 0) {
      break;
    }
    position++;
  }

  console.log(`Position ${position}`);
}

console.log("DAY 1");
part1();
part2();
