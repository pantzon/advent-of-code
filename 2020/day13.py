data = ""
with open("inputs/day13.txt", "r") as f:
    data = f.read()

START_TIME = 0
BUSES: list[int] = []

for line in data.split("\n"):
    if line != "":
        pieces = line.split(",")
        if len(pieces) == 1:
            START_TIME = int(pieces[0])
        else:
            BUSES = [int(num) if num != "x" else -1 for num in pieces]


def Part1() -> None:
    first_bus = 0
    min_wait = max(BUSES) + 1
    for i in [b for b in BUSES if b > 0]:
        wait = i - (START_TIME % i)
        if wait < min_wait:
            first_bus = i
            min_wait = wait
    print(
        "First Bus is {} after {} mins: {}".format(
            first_bus, min_wait, first_bus * min_wait
        )
    )


def Part2() -> None:
    buses = [(b, i % b) for i, b in enumerate(BUSES) if b > 0]
    t = buses[0][0]
    c = buses[0][1]
    for b in buses[1:]:
        m = b[0] - b[1]
        for i in range(0, b[0]):
            # Fun modulo math, have to find the new offset
            # with each added term.
            if m == (t * i + c) % b[0]:
                break
        c = t * i + c
        t = b[0] * t
    print("In time order: {}".format(c))


def main() -> None:
    print("Day 13")
    Part1()
    Part2()


if __name__ == "__main__":
    main()
