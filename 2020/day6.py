data = ""
with open("inputs/day6.txt", "r") as f:
    data = f.read()

any_groups = []
current = set()
for l in data.split("\n"):
    if l != "":
        current |= set(c for c in l)
    else:
        any_groups.append(current)
        current = set()


def Part1():
    print("Any Yes: {}".format(sum(len(s) for s in any_groups)))


every_groups = []
current = []
for l in data.split("\n"):
    if l != "":
        current.append(set(c for c in l))
    else:
        every_groups.append(set.intersection(*current))
        current = []


def Part2():
    print("Every Yes: {}".format(sum(len(s) for s in every_groups)))


def main():
    print("Day 6")
    Part1()
    Part2()


if __name__ == "__main__":
    main()
