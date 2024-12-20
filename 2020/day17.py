from typing import NamedTuple


data = ""
with open("inputs/day17.txt", "r") as f:
    data = f.read()


class Point3(NamedTuple):
    x: int
    y: int
    z: int


def GetInitialActive3D() -> set[Point3]:
    active: set[Point3] = set()
    for y, line in enumerate(data.split("\n")):
        for x, c in enumerate(line):
            if c == "#":
                active.add(Point3(x, y, 0))
    return active


def Neighbors3D(loc: Point3) -> set[Point3]:
    n: set[Point3] = set()
    for x in range(-1, 2):
        for y in range(-1, 2):
            for z in range(-1, 2):
                if x != 0 or y != 0 or z != 0:
                    n.add(Point3(loc.x + x, loc.y + y, loc.z + z))
    return n


def LocConway3D(active: set[Point3], loc: Point3) -> bool:
    active_neighbors = sum(1 if n in active else 0 for n in Neighbors3D(loc))
    return (loc in active and active_neighbors in (2, 3)) or (
        loc not in active and active_neighbors == 3
    )


def Conway3D(active: set[Point3]) -> set[Point3]:
    newActives: set[Point3] = set()
    for loc in active.union(*[Neighbors3D(n) for n in active]):
        if LocConway3D(active, loc):
            newActives.add(loc)
    return newActives


def Part1() -> None:
    actives = GetInitialActive3D()
    for _ in range(0, 6):
        actives = Conway3D(actives)
    print("3D Actives after 6 ", len(actives))


class Point4(NamedTuple):
    x: int
    y: int
    z: int
    w: int


def GetInitialActive4D() -> set[Point4]:
    active: set[Point4] = set()
    for y, line in enumerate(data.split("\n")):
        for x, c in enumerate(line):
            if c == "#":
                active.add(Point4(x, y, 0, 0))
    return active


def Neighbors4D(loc: Point4) -> set[Point4]:
    n: set[Point4] = set()
    for x in range(-1, 2):
        for y in range(-1, 2):
            for z in range(-1, 2):
                for w in range(-1, 2):
                    if x != 0 or y != 0 or z != 0 or w != 0:
                        n.add(Point4(loc.x + x, loc.y + y, loc.z + z, loc.w + w))
    return n


def LocConway4D(active: set[Point4], loc: Point4) -> bool:
    active_neighbors = sum(1 if n in active else 0 for n in Neighbors4D(loc))
    return (loc in active and active_neighbors in (2, 3)) or (
        loc not in active and active_neighbors == 3
    )


def Conway4D(active: set[Point4]) -> set[Point4]:
    newActives: set[Point4] = set()
    for loc in active.union(*[Neighbors4D(n) for n in active]):
        if LocConway4D(active, loc):
            newActives.add(loc)
    return newActives


def Part2() -> None:
    actives = GetInitialActive4D()
    for _ in range(0, 6):
        actives = Conway4D(actives)
    print("4D Actives after 6 ", len(actives))


def main() -> None:
    print("Day 17")
    Part1()
    Part2()


if __name__ == "__main__":
    main()
