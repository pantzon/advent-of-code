import { readFileSync } from "fs";

const data = readFileSync("./inputs/day8.txt", "utf-8");

const hexRe = /\\x[0-9a-f]{2}/;

function part1() {
  var strCount = 0;
  var memCount = 0;
  for (var line of data.split("\n")) {
    strCount += line.length;
    for (var i = 1; i < line.length - 1; i++) {
      switch (line.at(i)) {
        case "\\":
          switch (line.at(i + 1)) {
            case "\\":
            case '"':
              i++;
              break;
            default:
              if (hexRe.test(line.substring(i, i + 4))) {
                i += 3;
              }
          }
          break;
        default:
          break;
      }
      memCount += 1;
    }
  }
  console.log(`Extra space: ${strCount - memCount}`);
}

function part2() {
  var strCount = 0;
  var encCount = 0;
  for (var line of data.split("\n")) {
    strCount += line.length;
    encCount += 2 + line.length + [...line.matchAll(/\\|"/g)].length;
  }
  console.log(`Extra space: ${encCount - strCount}`);
}

console.log("DAY 8");
part1();
part2();
console.log("----------");
console.log("");
