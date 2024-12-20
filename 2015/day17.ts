import { readFileSync } from "fs";

const data = readFileSync("./inputs/day17.txt", "utf-8");

interface Data {
  target: number;
  containers: number[];
}

function processData(): Data {
  let target = -1;
  const containers: number[] = [];
  for (const line of data.split("\n")) {
    if (line != "") {
      const num = parseInt(line);
      if (target < 0) {
        target = num;
      } else {
        containers.push(num);
      }
    }
  }
  return { target: target, containers: containers.toSorted((a, b) => b - a) };
}

function combinations(d: Data): number[][] {
  const ways: number[][] = [];
  for (let i = 0; i < d.containers.length; i++) {
    const b = d.containers[i];
    if (b === d.target) {
      ways.push([b]);
    } else if (b < d.target) {
      ways.push(
        ...combinations({
          target: d.target - b,
          containers: d.containers.slice(i + 1),
        }).map((v) => [b, ...v])
      );
    }
  }
  return ways;
}

function part1() {
  console.log("Combos: ", combinations(processData()).length);
}

function part2() {
  let minSize = -1;
  let minCount = 0;
  for (const combo of combinations(processData())) {
    if (minSize < 0 || combo.length < minSize) {
      minSize = combo.length;
      minCount = 0;
    }
    if (minSize === combo.length) {
      minCount++;
    }
  }
  console.log("Minimum Combos: ", minSize, minCount);
}

console.log("DAY 17");
part1();
part2();
