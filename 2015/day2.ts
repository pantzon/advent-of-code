import { readFileSync } from "fs";

const data = readFileSync("./inputs/day2.txt", "utf-8");
const gifts = data.split("\n").filter((v) => v !== "");

function part1() {
  const amount = gifts
    .map((v) => {
      const [l, w, h] = v.split("x").map((v) => parseInt(v));
      const panels = [l * w, w * h, h * l];
      return Math.min(...panels) + panels.reduce((acc, v) => 2 * v + acc, 0);
    })
    .reduce((acc, v) => acc + v, 0);
  console.log(`Sq ft: ${amount}`);
}

function part2() {
  const amount = gifts
    .map((v) => {
      const [l, w, h] = v.split("x").map((v) => parseInt(v));
      const panels = [l + w, w + h, h + l];
      return 2 * Math.min(...panels) + l * w * h;
    })
    .reduce((acc, v) => acc + v, 0);
  console.log(`Ribbon length: ${amount}`);
}

console.log("DAY 2");
part1();
part2();
