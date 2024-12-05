import { createHash } from "crypto";
import { readFileSync } from "fs";

const data = readFileSync("./inputs/day4.txt", "utf-8");

function part1() {
  const MD5_RE = /^00000/;
  let i = 1;
  while (i < 1000000) {
    const hash = createHash("md5");
    hash.update(`${data}${i}`);
    if (MD5_RE.test(hash.digest("hex"))) {
      console.log(`Found 00000...: ${i}`);
      return;
    }
    i++;
  }
  console.log("DID NOT FIND i!");
}

function part2() {
  const MD5_RE = /^000000/;
  let i = 1;
  while (i < 10000000) {
    const hash = createHash("md5");
    hash.update(`${data}${i}`);
    if (MD5_RE.test(hash.digest("hex"))) {
      console.log(`Found 000000...: ${i}`);
      return;
    }
    i++;
  }
  console.log("DID NOT FIND i!");
}

console.log("DAY 4");
part1();
part2();
