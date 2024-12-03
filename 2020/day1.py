data = ""
with open("inputs/day1.txt", "r") as f:
    data = f.read()

expenses = [int(l) for l in data.split("\n") if l != ""]


def Part1():
    for i, x in enumerate(expenses):
        for y in expenses[i:]:
            if x + y == 2020:
                print("2 Multiple: {}".format(x * y))
                return
    print("None found!")


def Part2():
    for i, x in enumerate(expenses):
        for j, y in enumerate(expenses[i:]):
            for z in expenses[i:][j:]:
                if x + y + z == 2020:
                    print("3 Multiple: {}".format(x * y * z))
                    return
    print("None found!")


def main():
    print("Day 1")
    Part1()
    Part2()


if __name__ == "__main__":
    main()
