import { readFileSync } from "fs";

const data = readFileSync("./inputs/day5.txt", "utf-8");

const p1rule1Regex = /[aeiou]/g;
function p1rule1(val: string): boolean {
  const result = val.match(p1rule1Regex);
  return !!result && result.length > 2;
}

function p1rule2(val: string): boolean {
  return !!val.split("").find((v, i, arr) => v == arr[i + 1]);
}

const p1rule3Regex = /ab|cd|pq|xy/;
function p1rule3(val: string): boolean {
  return !p1rule3Regex.test(val);
}

function part1() {
  const nice = data
    .split("\n")
    .filter((v) => p1rule1(v) && p1rule2(v) && p1rule3(v));
  console.log(`P1 Nice: ${nice.length}`);
}

function p2rule1(val: string): boolean {
  return !!val
    .split("")
    .find((_, i) => val.indexOf(val.slice(i, i + 2), i + 2) > 0);
}

function p2rule2(val: string): boolean {
  return !!val.split("").find((v, i, arr) => v == arr[i + 2]);
}

function part2() {
  const nice = data.split("\n").filter((v) => p2rule1(v) && p2rule2(v));
  console.log(`P1 Nice: ${nice.length}`);
}

console.log("DAY 5");
part1();
part2();
