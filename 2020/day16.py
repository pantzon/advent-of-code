import re
from typing import NamedTuple

data = ""
with open("inputs/day16.txt", "r") as f:
    data = f.read()

FIELD_RE = re.compile(
    r"^(?P<name>[a-z ]+): (?P<rangeALow>[0-9]+)-(?P<rangeAHi>[0-9]+) or (?P<rangeBLow>[0-9]+)-(?P<rangeBHi>[0-9]+)$"
)

FIELDS: dict[str, tuple[tuple[int, int], tuple[int, int]]] = {}
TICKETS: list[list[int]] = []

for line in data.split("\n"):
    if line == "":
        continue
    match = FIELD_RE.match(line)
    if match:
        FIELDS[match.group("name")] = (
            (int(match.group("rangeALow")), int(match.group("rangeAHi"))),
            (int(match.group("rangeBLow")), int(match.group("rangeBHi"))),
        )
    pieces = line.split(",")
    if len(pieces) > 1:
        TICKETS.append([int(i) for i in pieces])


def GetInvalidValues(ticket: list[int]) -> list[int]:
    invalid_vals: list[int] = []
    for i in ticket:
        valid = False
        for ranges in FIELDS.values():
            for r in ranges:
                if r[0] <= i and i <= r[1]:
                    valid = True
                    break
            if valid:
                break
        if not valid:
            invalid_vals.append(i)
    return invalid_vals


def Part1() -> None:
    total = sum([sum(GetInvalidValues(ticket)) for ticket in TICKETS[1:]])
    print("Ticket Scanning Error Rate ", total)


def GetFieldIndexes(tickets: list[list[int]]) -> dict[int, str]:
    ticketSize = len(tickets[0])
    possibleFieldIndexes: dict[str, set[int]] = dict(
        (name, set()) for name in FIELDS.keys()
    )
    for name, (aRange, bRange) in FIELDS.items():
        for i in range(0, ticketSize):
            match = True
            for t in tickets:
                if (
                    t[i] < aRange[0]
                    or (aRange[1] < t[i] and bRange[0] > t[i])
                    or t[i] > bRange[1]
                ):
                    match = False
                    break
            if match:
                possibleFieldIndexes[name].add(i)
    fieldIndexes: dict[int, str] = {}
    lastCheck = len(possibleFieldIndexes)
    while len(possibleFieldIndexes) > 0:
        singleItemFields = [
            i for i in filter(lambda x: len(x[1]) == 1, possibleFieldIndexes.items())
        ]
        for name, fields in singleItemFields:
            i = fields.pop()
            fieldIndexes[i] = name
            possibleFieldIndexes.pop(name)
            for v in possibleFieldIndexes.values():
                v.discard(i)
        if len(possibleFieldIndexes) == lastCheck:
            print("ERROR! Can't filter further!")
            print(possibleFieldIndexes)
            break
        lastCheck = len(possibleFieldIndexes)
    return fieldIndexes


def Part2() -> None:
    cleaned = [ticket for ticket in TICKETS if len(GetInvalidValues(ticket)) == 0]
    fieldIndexes = GetFieldIndexes(cleaned)
    total = 1
    for i, name in fieldIndexes.items():
        if name.startswith("departure "):
            total *= TICKETS[0][i]
    print("Departure multiple: ", total)


def main() -> None:
    print("Day 16")
    Part1()
    Part2()


if __name__ == "__main__":
    main()
