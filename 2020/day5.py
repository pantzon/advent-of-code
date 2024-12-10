data = ""
with open("inputs/day5.txt", "r") as f:
    data = f.read()


data = set(
    int("".join(["0" if c == "F" or c == "L" else "1" for c in s]), 2)
    for s in data.split("\n")
)


def Part1() -> None:
    print("Max Seat: {}".format(max(data)))


def Part2() -> None:
    for i in range(min(data) + 1, max(data)):
        if i not in data:
            print("My Seat: {}".format(i))
            break


def main() -> None:
    print("Day 5")
    Part1()
    Part2()


if __name__ == "__main__":
    main()
