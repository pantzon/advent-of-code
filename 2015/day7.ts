import { readFileSync } from "fs";

const data = readFileSync("./inputs/day7.txt", "utf-8");

const instructionRe = /^(?<input>.+) -> (?<output>[a-z]+)$/;
const twoInputRe =
  /^(?<input1>[a-z0-9]+) (?<cmd>AND|OR|LSHIFT|RSHIFT) (?<input2>[a-z0-9]+)$/;
const oneInputRe = /^(?<cmd>NOT) (?<input1>[a-z0-9]+)$/;
enum Cmd {
  SET = "SET",
  AND = "AND",
  OR = "OR",
  NOT = "NOT",
  LSHIFT = "LSHIFT",
  RSHIFT = "RSHIFT",
}
interface Instruction {
  command: Cmd;
  inputs: string[];
}

const instructions: Record<string, Instruction> = {};
for (const line of data.split("\n")) {
  const match = instructionRe.exec(line);
  if (match && match.groups) {
    const inputGroups = twoInputRe.exec(match.groups["input"])?.groups ||
      oneInputRe.exec(match.groups["input"])?.groups || {
        input1: match.groups["input"],
      };
    instructions[match.groups["output"]] = {
      command: inputGroups["cmd"] ? (inputGroups["cmd"] as Cmd) : Cmd.SET,
      inputs: inputGroups["input2"]
        ? [inputGroups["input1"], inputGroups["input2"]]
        : [inputGroups["input1"]],
    };
  }
}

function findValue(wire: string, cache: Record<string, number>): number {
  const inst = instructions[wire];
  if (!inst) {
    console.log(`Unknown wire: ${wire}`);
    return -1;
  }
  if (cache[wire] < 0) {
    console.log(`Circular dependency! ${wire}`);
    return -1;
  }
  if (cache[wire]) {
    return cache[wire];
  }
  cache[wire] = -1;
  let retVal = 0;
  switch (inst.command) {
    case Cmd.SET:
      retVal = processInput(inst.inputs[0], cache);
      break;
    case Cmd.NOT:
      retVal = ~processInput(inst.inputs[0], cache);
      break;
    case Cmd.AND:
      retVal =
        processInput(inst.inputs[0], cache) &
        processInput(inst.inputs[1], cache);
      break;
    case Cmd.OR:
      retVal =
        processInput(inst.inputs[0], cache) |
        processInput(inst.inputs[1], cache);
      break;
    case Cmd.LSHIFT:
      retVal =
        processInput(inst.inputs[0], cache) <<
        processInput(inst.inputs[1], cache);
      break;
    case Cmd.RSHIFT:
      retVal =
        processInput(inst.inputs[0], cache) >>
        processInput(inst.inputs[1], cache);
      break;
    default:
      console.log(`Unknown Instruction: ${inst}`);
      return -1;
  }
  cache[wire] = uint16(retVal);
  return cache[wire];
}

function uint16(v: number) {
  return v & 0xffff;
}

function processInput(input: string, cache: Record<string, number>): number {
  if (input in instructions) {
    return findValue(input, cache);
  }
  return uint16(parseInt(input));
}

function part1() {
  console.log(`Original A: ${findValue("a", {})}`);
}

function part2() {
  instructions["b"] = { command: Cmd.SET, inputs: [`${findValue("a", {})}`] };
  console.log(`Recalc'd A: ${findValue("a", {})}`);
}

console.log("DAY 7");
part1();
part2();
