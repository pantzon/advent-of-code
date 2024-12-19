import { readFileSync } from "fs";

const data = readFileSync("./inputs/day14.txt", "utf-8");

const REINDEER_RE =
  /^(?<name>[A-Za-z]+) can fly (?<speed>[0-9]+) km\/s for (?<flytime>[0-9]+) seconds, but then must rest for (?<resttime>[0-9]+) seconds.$/;

interface Race {
  ticks: number;
  reindeer: Reindeer[];
}

interface Reindeer {
  name: string;
  speed: number;
  flyTime: number;
  restTime: number;
}

const race: Race = {
  ticks: 0,
  reindeer: [],
};

for (const line of data.split("\n")) {
  if (line != "") {
    const match = REINDEER_RE.exec(line);
    if (match && match.groups) {
      const deer: Reindeer = {
        name: match.groups["name"],
        speed: parseInt(match.groups["speed"]),
        flyTime: parseInt(match.groups["flytime"]),
        restTime: parseInt(match.groups["resttime"]),
      };
      race.reindeer.push(deer);
    } else {
      race.ticks = parseInt(line);
    }
  }
}

function getDistance(deer: Reindeer, ticks: number): number {
  const cycles = Math.floor(ticks / (deer.flyTime + deer.restTime));
  const leftover = ticks % (deer.flyTime + deer.restTime);
  return (
    (cycles * deer.flyTime + Math.min(leftover, deer.flyTime)) * deer.speed
  );
}

function part1() {
  console.log(
    race.reindeer
      .map((deer) => {
        return [deer.name, getDistance(deer, race.ticks)] as [string, number];
      })
      .toSorted((a, b) => b[1] - a[1])
  );
}

function part2() {
  const scores = new Map(
    race.reindeer.map((v) => [v.name, 0] as [string, number])
  );
  for (let i = 1; i < race.ticks; i++) {
    const winners = race.reindeer
      .map((deer) => {
        return [deer.name, getDistance(deer, i)] as [string, number];
      })
      .toSorted((a, b) => b[1] - a[1]);
    const topDistance = winners[0][1];
    for (const [winner, dist] of winners) {
      if (dist === topDistance) {
        scores.set(winner, scores.get(winner)! + 1);
      }
    }
  }
  console.log([...scores.entries()].toSorted((a, b) => b[1] - a[1]));
}

console.log("DAY 14");
part1();
part2();
