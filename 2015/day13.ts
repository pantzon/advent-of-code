import { readFileSync } from "fs";

const data = readFileSync("./inputs/day13.txt", "utf-8");

const LINE_RE =
  /^(?<A>[A-Za-z]+) would (?<dir>gain|lose) (?<score>[0-9]+) happiness units by sitting next to (?<B>[A-Za-z]+)\.$/;

const pairs: Map<string, Map<string, number>> = new Map();

for (const line of data.split("\n")) {
  const m = LINE_RE.exec(line);
  if (m && m.groups) {
    const a = m.groups["A"];
    const b = m.groups["B"];
    const val =
      (m.groups["dir"] === "gain" ? 1 : -1) * parseInt(m.groups["score"]);
    if (!pairs.has(a)) {
      pairs.set(a, new Map());
    }
    const aMap = pairs.get(a)!;
    aMap.set(b, (aMap.get(b) ?? 0) + val);
    if (!pairs.has(b)) {
      pairs.set(b, new Map());
    }
    const bMap = pairs.get(b)!;
    bMap.set(a, (bMap.get(a) ?? 0) + val);
  }
}

function findBest(start: string, last: string, seated: Set<string>) {
  if (seated.size === pairs.size) {
    return pairs.get(start)!.get(last)!;
  }
  let maximum = Number.MIN_SAFE_INTEGER;
  for (const a of [...pairs.keys()]) {
    if (!seated.has(a)) {
      const moreSeated = new Set(seated);
      moreSeated.add(a);
      const weight =
        start && last
          ? pairs.get(last)!.get(a)! + findBest(start, a, moreSeated)
          : findBest(a, a, moreSeated);
      if (weight > maximum) {
        maximum = weight;
      }
    }
  }
  return maximum;
}

function part1() {
  console.log(`Most Happiness: ${findBest("", "", new Set())}`);
}

function part2() {
  for (const m of pairs.values()) {
    m.set("Me", 0);
  }
  pairs.set("Me", new Map([...pairs.keys()].map((k) => [k, 0])));
  console.log(`Most Happiness With Me: ${findBest("", "", new Set())}`);
}

console.log("DAY 13");
part1();
part2();
