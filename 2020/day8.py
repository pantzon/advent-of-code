from dataclasses import dataclass
from typing import Literal, NamedTuple, cast


data = ""
with open("inputs/day8.txt", "r") as f:
    data = f.read()


@dataclass
class Memory(object):
    acc = 0
    terminated = False
    visited: set[int]


type Cmd = Literal["acc"] | Literal["jmp"] | Literal["nop"]


class Instruction(NamedTuple):
    cmd: Cmd
    val: int


program: list[Instruction] = []
for line in data.split("\n"):
    cmd, val = line.split(" ")
    program.append(
        Instruction(
            cast(Cmd, cmd),
            (1 if val[0] == "+" else -1) * int(val[1:]),
        )
    )


def Run(pgrm: list[Instruction]) -> Memory:
    m = Memory(visited=set())
    curr = 0
    while curr >= 0 and curr < len(pgrm) and curr not in m.visited:
        m.visited.add(curr)
        inst = pgrm[curr]
        if inst.cmd == "jmp":
            curr += inst.val
            continue
        if inst.cmd == "acc":
            m.acc += inst.val
        curr += 1
    if curr == len(pgrm):
        m.terminated = True
    return m


def Part1() -> None:
    print("ACC at infinite loop: {}".format(Run(program).acc))


def Part2() -> None:
    m = Memory(visited=set())
    for index in range(0, len(program)):
        orig = program[index]
        if orig.cmd == "nop":
            program[index] = Instruction("jmp", orig.val)
            m = Run(program)
        elif orig.cmd == "jmp":
            program[index] = Instruction("nop", orig.val)
            m = Run(program)
        program[index] = orig
        if m.terminated:
            print("Fix, flip line {}".format(index))
            break
    print("ACC for Fixed: {}".format(m.acc))


def main() -> None:
    print("Day 8")
    Part1()
    Part2()


if __name__ == "__main__":
    main()
