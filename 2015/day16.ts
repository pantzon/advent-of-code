import { readFileSync } from "fs";

const data = readFileSync("./inputs/day16.txt", "utf-8");

interface Sue {
  children?: number;
  cats?: number;
  samoyeds?: number;
  pomeranians?: number;
  akitas?: number;
  vizslas?: number;
  goldfish?: number;
  trees?: number;
  cars?: number;
  perfumes?: number;
}

const RealSue: Sue = {
  children: 3,
  cats: 7,
  samoyeds: 2,
  pomeranians: 3,
  akitas: 0,
  vizslas: 0,
  goldfish: 5,
  trees: 3,
  cars: 2,
  perfumes: 1,
};

const SUE_RE = /^Sue [0-9]+: (?<fields>.*)$/;

const sues: Sue[] = [];

for (const line of data.split("\n")) {
  const match = SUE_RE.exec(line);
  if (match && match.groups) {
    const currentSue: Sue = {};
    for (const field of match.groups["fields"].split(", ")) {
      const [name, num] = field.split(": ");
      currentSue[name as keyof Sue] = parseInt(num);
    }
    sues.push(currentSue);
  }
}

function doTheyMatch(target: Sue, guess: Sue): boolean {
  for (const f in guess) {
    const k = f as keyof Sue;
    if (target[k] !== guess[k]) {
      return false;
    }
  }
  return true;
}

function part1() {
  for (let i = 0; i < sues.length; i++) {
    if (doTheyMatch(RealSue, sues[i])) {
      console.log("Found Sue! ", i + 1);
      break;
    }
  }
}

function doTheyMatch2(target: Sue, guess: Sue): boolean {
  for (const f in guess) {
    const k = f as keyof Sue;
    if (k === "cats" || k === "trees") {
      if (target[k]! >= guess[k]!) {
        return false;
      }
    } else if (k === "pomeranians" || k === "goldfish") {
      if (target[k]! <= guess[k]!) {
        return false;
      }
    } else if (target[k] !== guess[k]) {
      return false;
    }
  }
  return true;
}

function part2() {
  for (let i = 0; i < sues.length; i++) {
    if (doTheyMatch2(RealSue, sues[i])) {
      console.log("Found Real Sue! ", i + 1);
      break;
    }
  }
}

console.log("DAY 16");
part1();
part2();
