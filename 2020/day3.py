data = ""
with open("inputs/day3.txt", "r") as f:
    data = f.read()

rows = data.split("\n")


def Traverse(y: int, x: int):
    index = 0
    trees = 0
    for row in rows[::y]:
        index %= len(row)
        if row[index] == "#":
            trees += 1
        index += x
    return trees


def Part1() -> None:
    print("Trees Encountered: {}".format(Traverse(1, 3)))


def Part2() -> None:
    print(
        "Multiplied Trees: {}".format(
            Traverse(1, 1)
            * Traverse(1, 3)
            * Traverse(1, 5)
            * Traverse(1, 7)
            * Traverse(2, 1)
        )
    )


def main() -> None:
    print("Day 3")
    Part1()
    Part2()


if __name__ == "__main__":
    main()
