import { readFileSync } from "fs";

const data = readFileSync("./inputs/day6.txt", "utf-8");

const instructionRe =
  /^(?<cmd>turn on|turn off|toggle) (?<sX>[0-9]+),(?<sY>[0-9]+) through (?<eX>[0-9]+),(?<eY>[0-9]+)$/;

function process(data: string): [boolean[], number[]] {
  const p1Lights = new Array<boolean>(1_000_000).fill(false);
  const p2Lights = new Array<number>(1_000_000).fill(0);

  for (const line of data.split("\n")) {
    const match = instructionRe.exec(line);
    if (match && match.groups) {
      for (
        let x = parseInt(match.groups["sX"]);
        x <= parseInt(match.groups["eX"]);
        x++
      ) {
        for (
          let y = parseInt(match.groups["sY"]);
          y <= parseInt(match.groups["eY"]);
          y++
        ) {
          const i = x * 1000 + y;
          switch (match.groups["cmd"]) {
            case "turn on":
              p1Lights[i] = true;
              p2Lights[i] += 1;
              break;
            case "turn off":
              p1Lights[i] = false;
              if (p2Lights[i] > 0) {
                p2Lights[i] -= 1;
              }
              break;
            case "toggle":
              p1Lights[i] = !p1Lights[i];
              p2Lights[i] += 2;
              break;
            default:
              console.log(`Unknown Command! ${line}`);
          }
        }
      }
    } else {
      console.log("Bad line: " + line);
    }
  }
  return [p1Lights, p2Lights];
}

const [p1, p2] = process(data);

function part1() {
  const lit = p1.reduce((acc, v) => (v ? acc + 1 : acc), 0);
  console.log(`Lights Lit: ${lit}`);
}

function part2() {
  const brightness = p2.reduce((acc, v) => acc + v, 0);
  console.log(`Lights Brightness: ${brightness}`);
}

console.log("DAY 6");
part1();
part2();
