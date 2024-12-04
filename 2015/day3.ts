import { readFileSync } from "fs";

const data = readFileSync("./inputs/day3.txt", "utf-8");

class Point {
  constructor(readonly x: number, readonly y: number) {}

  getNext(val: string) {
    switch (val) {
      case "^":
        return new Point(this.x, this.y + 1);
      case "v":
        return new Point(this.x, this.y - 1);
      case ">":
        return new Point(this.x + 1, this.y);
      case "<":
        return new Point(this.x - 1, this.y);
      default:
        throw Error(`Unknown direction: ${val}`);
    }
  }

  toString() {
    return `[${this.x},${this.y}]`;
  }
}

function part1() {
  const houseDeliveries: Map<string, number> = data.split("").reduce(
    (acc, val) => {
      acc.last = acc.last.getNext(val);
      const key = acc.last.toString();
      if (!acc.visits.has(key)) {
        acc.visits.set(key, 0);
      }
      acc.visits.set(key, acc.visits.get(key)! + 1);
      return acc;
    },
    {
      visits: new Map([["[0,0]", 1]]),
      last: new Point(0, 0),
    }
  ).visits;
  console.log(`Houses receiving presents: ${houseDeliveries.size}`);
}

function part2() {
  const houseDeliveries: Map<string, number> = data.split("").reduce(
    (acc, val) => {
      let key = "";
      if (acc.santaTurn) {
        acc.santaLast = acc.santaLast.getNext(val);
        key = acc.santaLast.toString();
      } else {
        acc.roboLast = acc.roboLast.getNext(val);
        key = acc.roboLast.toString();
      }
      acc.santaTurn = !acc.santaTurn;
      if (!acc.visits.has(key)) {
        acc.visits.set(key, 0);
      }
      acc.visits.set(key, acc.visits.get(key)! + 1);
      return acc;
    },
    {
      visits: new Map([["[0,0]", 2]]),
      santaLast: new Point(0, 0),
      roboLast: new Point(0, 0),
      santaTurn: true,
    }
  ).visits;
  console.log(`Houses receiving presents: ${houseDeliveries.size}`);
}

console.log("DAY 3");
part1();
part2();
