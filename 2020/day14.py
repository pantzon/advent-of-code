import math
import re


data = ""
with open("inputs/day14.txt", "r") as f:
    data = f.read()

MASK_RE = re.compile(r"^mask = (?P<mask>[X01]+)$")
MEMORY_RE = re.compile(r"^mem\[(?P<loc>[0-9]+)\] = (?P<value>[0-9]+)$")


def LoadData() -> dict[int, int]:
    memory: dict[int, int] = {}
    andMask = 0
    orMask = 0
    for line in data.split("\n"):
        m = MASK_RE.match(line)
        if m is not None:
            andMask = 0
            orMask = 0
            for c in m.group("mask"):
                andMask = andMask << 1
                orMask = orMask << 1
                if c == "X":
                    andMask += 1
                elif c == "1":
                    andMask += 1
                    orMask += 1
        else:
            m = MEMORY_RE.match(line)
            if m is not None:
                value = int(m.group("value"))
                memory[int(m.group("loc"))] = value & andMask | orMask
    return memory


def Part1() -> None:
    mem = LoadData()
    print("Sum of memory: ", sum(mem.values()))


def LoadData2() -> dict[int, int]:
    memory: dict[int, int] = {}
    masks: list[tuple[int, int]] = []
    orMask = 0
    for line in data.split("\n"):
        m = MASK_RE.match(line)
        if m is not None:
            masks = [(0, 0)]
            orMask = 0
            for c in m.group("mask"):
                masks = [((aMask << 1) + 1, oMask << 1) for aMask, oMask in masks]
                orMask = orMask << 1
                if c == "1":
                    orMask += 1
                if c == "X":
                    newMasks = []
                    for aMask, oMask in masks:
                        newMasks.append((aMask - 1, oMask))
                        newMasks.append((aMask, oMask + 1))
                    masks = newMasks
            # print("Or {:b}".format(orMask))
            # for aMask, oMask in masks:
            #     print(">> and {:b}, or {:b}".format(aMask, oMask))
        else:
            m = MEMORY_RE.match(line)
            if m is not None:
                value = int(m.group("value"))
                loc = int(m.group("loc")) | orMask
                for aMask, oMask in masks:
                    memory[loc & aMask | oMask] = value
    return memory


def Part2() -> None:
    mem = LoadData2()
    print("Sum of memory: ", sum(mem.values()))


def main() -> None:
    print("Day 14")
    Part1()
    Part2()


if __name__ == "__main__":
    main()
