import re

data = ""
with open("inputs/day7.txt", "r") as f:
    data = f.read()

BAG_RE = re.compile(r"^(?P<count>[0-9]+) (?P<color>[a-z ]+) bags?$")
RULE_RE = re.compile(r"^(?P<color>[a-z ]+) bags contain (?P<bags>[0-9a-z, ]+)\.$")

bag_groupings: dict[str, dict[str, int]] = {}
bag_holdings: dict[str, dict[str, int]] = {}
for line in data.split("\n"):
    rule = RULE_RE.match(line)
    if rule is None:
        print("> ERROR: Bad rule! {}".format(line))
    else:
        outer_bag = rule.group("color")
        inner_bags = rule.group("bags")
        if inner_bags != "no other bags":
            bag_holdings[outer_bag] = {}
            for b in inner_bags.split(", "):
                bag_match = BAG_RE.match(b)
                if bag_match is None:
                    print("> ERROR: Bad bag! {} ({})".format(b, line))
                else:
                    count = int(bag_match.group("count"))
                    inner_bag = bag_match.group("color")
                    if inner_bag not in bag_groupings:
                        bag_groupings[inner_bag] = {}
                    bag_groupings[inner_bag][outer_bag] = count
                    bag_holdings[outer_bag][inner_bag] = count


def GetBagHolders(color: str, holders: set[str]):
    if color not in bag_groupings:
        return
    for k in bag_groupings[color].keys():
        holders.add(k)
        GetBagHolders(k, holders)


def Part1() -> None:
    all_shiny_holders = set[str]()
    if "shiny gold" not in bag_holdings:
        all_shiny_holders.add("shiny gold")
    GetBagHolders("shiny gold", all_shiny_holders)
    print("Bags for Shiny Gold: {}".format(len(all_shiny_holders)))


def GetBagsHeld(color: str) -> int:
    total = 0
    for k, count in bag_holdings.get(color, {}).items():
        total += count + count * GetBagsHeld(k)
    return total


def Part2() -> None:
    print("Bags in Shiny Gold: {}".format(GetBagsHeld("shiny gold")))


def main() -> None:
    print("Day 7")
    Part1()
    Part2()


if __name__ == "__main__":
    main()
