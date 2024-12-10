import { readFileSync } from "fs";

const data = readFileSync("./inputs/day9.txt", "utf-8");

type Graph = Map<string, Node>;
function newGraph(toClone?: Graph): Graph {
  return new Map<string, Node>(toClone);
}
const GRAPH: Graph = newGraph();

interface Node {
  id: string;
  edges: Map<Node, number>;
}

const EDGE_RE = /^(?<a>[A-Za-z]+) to (?<b>[A-Za-z]+) = (?<weight>[0-9]+)$/;

function getOrMakeNode(id: string): Node {
  if (!GRAPH.has(id)) {
    GRAPH.set(id, { id: id, edges: new Map() });
  }
  return GRAPH.get(id)!;
}

for (const line of data.split("\n")) {
  const m = EDGE_RE.exec(line);
  if (m && m.groups) {
    const a = getOrMakeNode(m.groups["a"]);
    const b = getOrMakeNode(m.groups["b"]);
    const weight = parseInt(m.groups["weight"]);
    a.edges.set(b, weight);
    b.edges.set(a, weight);
  }
}

function findBestPath(nodes: Graph): { weight: number; path: Node[] } {
  const vals = [...nodes.values()];
  if (vals.length === 2) {
    return {
      weight: vals[0].edges.get(vals[1])!,
      path: vals,
    };
  }
  return vals
    .map((v) => {
      const newNodes = newGraph(nodes);
      newNodes.delete(v.id);
      const best = findBestPath(newNodes);
      const first = best.path[0];
      const last = best.path[best.path.length - 1];
      if (v.edges.get(first)! < v.edges.get(last)!) {
        best.weight += v.edges.get(first)!;
        best.path.unshift(v);
      } else {
        best.weight += v.edges.get(last)!;
        best.path.push(v);
      }
      return best;
    })
    .sort((a, b) => a.weight - b.weight)[0];
}

function findWorstPath(nodes: Graph): { weight: number; path: Node[] } {
  const vals = [...nodes.values()];
  if (vals.length === 2) {
    return {
      weight: vals[0].edges.get(vals[1])!,
      path: vals,
    };
  }
  return vals
    .map((v) => {
      const newNodes = newGraph(nodes);
      newNodes.delete(v.id);
      const worst = findWorstPath(newNodes);
      const first = worst.path[0];
      const last = worst.path[worst.path.length - 1];
      if (v.edges.get(first)! > v.edges.get(last)!) {
        worst.weight += v.edges.get(first)!;
        worst.path.unshift(v);
      } else {
        worst.weight += v.edges.get(last)!;
        worst.path.push(v);
      }
      return worst;
    })
    .sort((a, b) => b.weight - a.weight)[0];
}

function part1() {
  const best = findBestPath(GRAPH);
  console.log(`Shortest distance: ${best.weight}`);
}

function part2() {
  const worst = findWorstPath(GRAPH);
  console.log(`Longest distance: ${worst.weight}`);
}

console.log("DAY 9");
part1();
part2();
