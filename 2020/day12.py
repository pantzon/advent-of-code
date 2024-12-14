from enum import Enum


data = ""
with open("inputs/day12.txt", "r") as f:
    data = f.read()

INSTRUCTIONS: list[tuple[str, int]] = [
    (l[0], int(l[1:])) for l in data.split("\n") if len(l) > 0
]


class Dir(Enum):
    NORTH = 0
    EAST = 1
    SOUTH = 2
    WEST = 3


def Move(direction: Dir, x: int, y: int, val: int) -> tuple[int, int]:
    if direction == Dir.NORTH:
        y += val
    elif direction == Dir.EAST:
        x += val
    elif direction == Dir.SOUTH:
        y -= val
    elif direction == Dir.WEST:
        x -= val
    return x, y


def P1FollowInstructions(dir: Dir, x: int, y: int) -> tuple[Dir, int, int]:
    for r, val in INSTRUCTIONS:
        if r == "N":
            x, y = Move(Dir.NORTH, x, y, val)
        elif r == "E":
            x, y = Move(Dir.EAST, x, y, val)
        elif r == "S":
            x, y = Move(Dir.SOUTH, x, y, val)
        elif r == "W":
            x, y = Move(Dir.WEST, x, y, val)
        elif r == "L":
            dir = Dir((dir.value - (val / 90)) % 4)
        elif r == "R":
            dir = Dir((dir.value + (val / 90)) % 4)
        elif r == "F":
            x, y = Move(dir, x, y, val)
        else:
            print("UNKNOWN INSTRUCTION {} {}".format(r, val))
    return dir, x, y


def Part1() -> None:
    _, x, y = P1FollowInstructions(Dir.EAST, 0, 0)
    print("Mahattan Distance: {}".format(abs(x) + abs(y)))


def P2FollowInstructions(x: int, y: int, wx: int, wy: int) -> tuple[int, int]:
    for r, val in INSTRUCTIONS:
        if r == "N":
            wx, wy = Move(Dir.NORTH, wx, wy, val)
        elif r == "E":
            wx, wy = Move(Dir.EAST, wx, wy, val)
        elif r == "S":
            wx, wy = Move(Dir.SOUTH, wx, wy, val)
        elif r == "W":
            wx, wy = Move(Dir.WEST, wx, wy, val)
        elif r == "L":
            val = int(val / 90)
            for _ in range(0, val):
                wx, wy = (-wy, wx)
        elif r == "R":
            val = int(val / 90)
            for _ in range(0, val):
                wx, wy = (wy, -wx)
        elif r == "F":
            for _ in range(0, val):
                x, y = Move(Dir.NORTH, x, y, wy)
                x, y = Move(Dir.EAST, x, y, wx)
        else:
            print("UNKNOWN INSTRUCTION {} {}".format(r, val))
    return x, y


def Part2() -> None:
    x, y = P2FollowInstructions(0, 0, 10, 1)
    print("Mahattan Distance: {}".format(abs(x) + abs(y)))


def main() -> None:
    print("Day 12")
    Part1()
    Part2()


if __name__ == "__main__":
    main()
