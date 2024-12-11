import { readFileSync } from "fs";

const data = readFileSync("./inputs/day10.txt", "utf-8");

function lookAndSay(input: string): string {
  let count = 1;
  let last = input.at(0);
  let output = "";
  for (const c of input.substring(1)) {
    if (c == last) {
      count++;
      continue;
    }
    output = `${output}${count}${last}`;
    last = c;
    count = 1;
  }
  output = `${output}${count}${last}`;
  return output;
}

function part1() {
  let last = data;
  for (let i = 0; i < 40; i++) {
    last = lookAndSay(last);
  }
  console.log(`40*LookAndSay: ${last.length}`);
}

function part2() {
  let last = data;
  for (let i = 0; i < 50; i++) {
    last = lookAndSay(last);
  }
  console.log(`50*LookAndSay: ${last.length}`);
}

console.log("DAY 10");
part1();
part2();
