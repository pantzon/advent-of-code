import { readFileSync } from "fs";

const data = readFileSync("./inputs/day11.txt", "utf-8");

const A = "a".charCodeAt(0);
const I = "i".charCodeAt(0);
const L = "i".charCodeAt(0);
const O = "i".charCodeAt(0);

function ToInts(d: string) {
  return d
    .split("")
    .reverse()
    .map((c) => c.charCodeAt(0) - A);
}

function ToString(pwd: number[]) {
  return pwd
    .reverse()
    .map((c) => String.fromCharCode(c + A))
    .join("");
}

function Increment(pwd: number[]) {
  const bad = new Set([I, L, O]);
  for (let i = 0; i < pwd.length; i++) {
    const c = pwd[i];
    // Part 1 Rule 2
    const inc = bad.has(c + 1 + A) ? 2 : 1;
    pwd[i] = (c + inc) % 26;
    if (pwd[i] != 0) {
      break;
    }
  }
}

function CheckP1(pwd: number[]) {
  return CheckP1R1(pwd) && CheckP1R3(pwd);
}

function CheckP1R1(pwd: number[]) {
  for (let i = 0; i < pwd.length - 2; i++) {
    const c = pwd[i];
    if (c == pwd[i + 1] + 1 && c == pwd[i + 2] + 2) {
      return true;
    }
  }
  return false;
}

function CheckP1R3(pwd: number[]) {
  let last = -1;
  for (let i = 0; i < pwd.length - 1; i++) {
    const c = pwd[i];
    if (c == pwd[i + 1]) {
      if (last == -1) {
        last = i;
      } else if (last != i - 1) {
        return true;
      }
    }
  }
  return false;
}

function part1() {
  const pwd = ToInts(data);
  do {
    Increment(pwd);
  } while (!CheckP1(pwd));
  console.log(`Next valid: ${ToString(pwd)}`);
}

function part2() {
  const pwd = ToInts(data);
  do {
    Increment(pwd);
  } while (!CheckP1(pwd));
  do {
    Increment(pwd);
  } while (!CheckP1(pwd));
  console.log(`Next Next valid: ${ToString(pwd)}`);
}

console.log("DAY 11");
part1();
part2();
