data = ""
with open("inputs/day15.txt", "r") as f:
    data = f.read()

MEMORY: dict[int, list[int]] = {}
SPOKEN: list[int] = []
for num in data.split(","):
    SPOKEN.append(int(num))
    MEMORY[int(num)] = [len(SPOKEN)]


def Run(ticks: int) -> None:
    for i in range(0, ticks):
        next = 0
        if len(MEMORY[SPOKEN[-1]]) != 1:
            next = MEMORY[SPOKEN[-1]][-1] - MEMORY[SPOKEN[-1]][-2]
        SPOKEN.append(next)
        if next not in MEMORY:
            MEMORY[next] = []
        MEMORY[next].append(len(SPOKEN))


def Part1() -> None:
    Run(2020)
    print("2020th: ", SPOKEN[2020 - 1])


def Part2() -> None:
    Run(30000000)
    print("30000000th: ", SPOKEN[30000000 - 1])


def main() -> None:
    print("Day 15")
    Part1()
    Part2()


if __name__ == "__main__":
    main()
