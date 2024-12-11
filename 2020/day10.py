data = ""
with open("inputs/day10.txt", "r") as f:
    data = f.read()

adapters = [0] + sorted([int(s) for s in data.split("\n") if len(s) > 0])
adapters.append(adapters[-1] + 3)

diffs: dict[int, int] = {}
for i, a in enumerate(adapters[:-1]):
    diff = adapters[i + 1] - a
    if diff not in diffs:
        diffs[diff] = 0
    diffs[diff] += 1


def Part1() -> None:
    print("1s*3s: {}".format(diffs[1] * diffs[3]))


def Part2() -> None:
    vals: dict[int, int] = {adapters[-1]: 1}
    for a in reversed(adapters[:-1]):
        vals[a] = vals.get(a + 1, 0) + vals.get(a + 2, 0) + vals.get(a + 3, 0)
    print("Distinct Ways: {}".format(vals.get(0)))


def main() -> None:
    print("Day 10")
    Part1()
    Part2()


if __name__ == "__main__":
    main()
