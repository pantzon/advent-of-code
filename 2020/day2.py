from typing import NamedTuple


data = ""
with open("inputs/day2.txt", "r") as f:
    data = f.read()


class Policy(NamedTuple):
    minimum: int
    maximum: int
    char: str

    def validateSledPolicy(self, pswd: str):
        count = sum(1 if i == self.char else 0 for i in pswd)
        return count >= self.minimum and count <= self.maximum

    def validateTobogganPolicy(self, pswd: str):
        first = self.char == pswd[self.minimum - 1]
        second = self.char == pswd[self.maximum - 1]
        return first != second


pols_and_pswds = []
for line in data.split("\n"):
    if line == "":
        continue
    rge, char, pswd = line.split(" ")
    minimum, maximum = [int(i) for i in rge.split("-")]
    char = char[:-1]
    pols_and_pswds.append((Policy(minimum, maximum, char), pswd))


def Part1():
    print(
        "Valid Sled Count: {}".format(
            sum(pol.validateSledPolicy(pswd) for pol, pswd in pols_and_pswds)
        )
    )


def Part2():
    print(
        "Valid Toboggan Count: {}".format(
            sum(pol.validateTobogganPolicy(pswd) for pol, pswd in pols_and_pswds)
        )
    )


def main():
    print("Day 2")
    Part1()
    Part2()


if __name__ == "__main__":
    main()
