data = ""
with open("inputs/day11.txt", "r") as f:
    data = f.read()

seatmap: list[list[str]] = [
    [c for c in line] for line in data.split("\n") if line != ""
]

SEAT_DIFFS = [(-1, -1), (-1, 0), (-1, 1), (0, 1), (1, 1), (1, 0), (1, -1), (0, -1)]


def IterateP1(seatmap: list[list[str]]) -> tuple[list[list[str]], bool, int]:
    occupied = 0
    changed = False
    newMap: list[list[str]] = []
    for y, r in enumerate(seatmap):
        newMap.append([])
        for x, seat in enumerate(r):
            newSeat = seat
            if seat != ".":
                occupiedNeighbors = 0
                for dx, dy in SEAT_DIFFS:
                    if IsTaken(seatmap, x + dx, y + dy):
                        occupiedNeighbors += 1
                if seat == "L" and occupiedNeighbors == 0:
                    changed = True
                    newSeat = "#"
                    occupied += 1
                elif seat == "#" and occupiedNeighbors >= 4:
                    changed = True
                    newSeat = "L"
                elif seat == "#":
                    occupied += 1
            newMap[y].append(newSeat)
    return (newMap, changed, occupied)


def Part1() -> None:
    seats = seatmap
    occupied = 0
    changed = True
    while changed:
        seats, changed, occupied = IterateP1(seats)
    print("Number of taken seats: {}".format(occupied))


def IsTaken(seatmap: list[list[str]], x: int, y: int) -> bool | None:
    if 0 <= y and y < len(seatmap) and 0 <= x and x < len(seatmap[y]):
        if seatmap[y][x] == ".":
            return None
        elif seatmap[y][x] == "L":
            return False
        else:
            return True
    else:
        # Leaving the boundaries counts as an empty seat.
        return False


def IterateP2(seatmap: list[list[str]]) -> tuple[list[list[str]], bool, int]:
    occupied = 0
    changed = False
    newMap: list[list[str]] = []
    for y, r in enumerate(seatmap):
        newMap.append([])
        for x, seat in enumerate(r):
            newSeat = seat
            if seat != ".":
                occupiedNeighbors = 0
                for dx, dy in SEAT_DIFFS:
                    isTaken = None
                    mul = 0
                    while isTaken is None:
                        mul += 1
                        isTaken = IsTaken(seatmap, x + (dx * mul), y + (dy * mul))
                    if isTaken:
                        occupiedNeighbors += 1
                if seat == "L" and occupiedNeighbors == 0:
                    changed = True
                    newSeat = "#"
                    occupied += 1
                elif seat == "#" and occupiedNeighbors >= 5:
                    changed = True
                    newSeat = "L"
                elif seat == "#":
                    occupied += 1
            newMap[y].append(newSeat)
    return (newMap, changed, occupied)


def Part2() -> None:
    seats = seatmap
    occupied = 0
    changed = True
    while changed:
        seats, changed, occupied = IterateP2(seats)
    print("Number of taken seats: {}".format(occupied))


def main() -> None:
    print("Day 11")
    Part1()
    Part2()


if __name__ == "__main__":
    main()
