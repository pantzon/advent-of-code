import re
from typing import Generator


data = ""
with open("inputs/day9.txt", "r") as f:
    data = f.read()

NUM_RE = re.compile("[0-9]+\n")


def DataSequence() -> Generator[int]:
    for match in NUM_RE.finditer(data):
        yield int(match.group().rstrip())


def FindError(preamble: int) -> int:
    window: list[int] = []
    i = 0
    for v in DataSequence():
        if len(window) < preamble:
            window.append(v)
        else:
            f = False
            for index, x in enumerate(window):
                for y in window[index:]:
                    if (x + y) == v:
                        f = True
                        break
                if f:
                    break
            if not f:
                return v
            window[i] = v
            i += 1
            i %= preamble
    raise Exception("Sequence exhausted!")


def Part1() -> None:
    bad = FindError(25)
    print("First bad value: {}".format(bad))


def FindBadRange(v: int) -> tuple[int, int]:
    vals: list[int] = []
    found = None
    for i in DataSequence():
        total = i
        mx = i
        mn = i
        for next in reversed(vals):
            total += next
            if next > mx:
                mx = next
            if next < mn:
                mn = next
            if total == v:
                return (mx, mn)
        vals.append(i)
    return (-1, -1)


def Part2() -> None:
    bad = FindError(25)
    mxAndMn = FindBadRange(bad)
    print("Range diff: {}".format(sum(mxAndMn)))


def main() -> None:
    print("Day 9")
    Part1()
    Part2()


if __name__ == "__main__":
    main()
