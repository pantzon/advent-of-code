import { readFileSync } from "fs";

const data = readFileSync("./inputs/day12.txt", "utf-8");

type Data = number | string | Data[] | { [key: string]: Data };

const dataObj: Data = JSON.parse(data);

function sum(d: Data): number {
  if (typeof d === "number") {
    return d;
  }
  if (typeof d === "string") {
    return 0;
  }
  if (Array.isArray(d)) {
    return d.reduce((acc: number, subData) => acc + sum(subData), 0);
  }
  return Object.values(d).reduce(
    (acc: number, subData) => acc + sum(subData),
    0
  );
}

function part1() {
  console.log(`Sum: ${sum(dataObj)}`);
}

function noRedSum(d: Data): number {
  if (typeof d === "number") {
    return d;
  }
  if (typeof d === "string") {
    return 0;
  }
  if (Array.isArray(d)) {
    return d.reduce((acc: number, subData) => acc + noRedSum(subData), 0);
  }
  const vals = Object.values(d);
  if (vals.find((v) => v === "red")) {
    return 0;
  }
  return vals.reduce((acc: number, subData) => acc + noRedSum(subData), 0);
}

function part2() {
  console.log(`Sum: ${noRedSum(dataObj)}`);
}

console.log("DAY 12");
part1();
part2();
