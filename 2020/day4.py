import math
import re

data = ""
with open("inputs/day4.txt", "r") as f:
    data = f.read()

IDS = {
    "byr": 1,
    "iyr": 1 << 1,
    "eyr": 1 << 2,
    "hgt": 1 << 3,
    "hcl": 1 << 4,
    "ecl": 1 << 5,
    "pid": 1 << 6,
}
VALID = math.pow(2, 7) - 1


def Part1():
    total = 0
    for passport in data.split("\n\n"):
        passport_mask = 0
        for line in passport.split("\n"):
            for field_and_val in line.split(" "):
                field = field_and_val.split(":")[0]
                if field in IDS:
                    passport_mask |= IDS[field]
        if passport_mask == VALID:
            total += 1
    print("Valid passports: {}".format(total))


HCL_RE = re.compile(r"#[0-9a-f]{6}")
ECL_RE = re.compile(r"amb|blu|brn|gry|grn|hzl|oth")
PID_RE = re.compile(r"[0-9]{9}")


def CheckField(field: str, val: str):
    if field == "byr":
        return len(val) == 4 and 1920 <= int(val) and int(val) <= 2002
    elif field == "iyr":
        return len(val) == 4 and 2010 <= int(val) and int(val) <= 2020
    elif field == "eyr":
        return len(val) == 4 and 2020 <= int(val) and int(val) <= 2030
    elif field == "hgt":
        if len(val) > 2:
            unit = val[-2:]
            num = int(val[:-2])
            if unit == "cm":
                return 150 <= num and num <= 193
            elif unit == "in":
                return 59 <= num and num <= 76
    elif field == "hcl":
        return HCL_RE.fullmatch(val) is not None
    elif field == "ecl":
        return ECL_RE.fullmatch(val) is not None
    elif field == "pid":
        return PID_RE.fullmatch(val) is not None
    return False


def Part2():
    total = 0
    for passport in data.split("\n\n"):
        passport_mask = 0
        for line in passport.split("\n"):
            for field_and_val in line.split(" "):
                if field_and_val != "":
                    field, val = field_and_val.split(":")
                    if CheckField(field, val):
                        passport_mask |= IDS[field]
        if passport_mask == VALID:
            total += 1
    print("Valid passports: {}".format(total))


def main():
    print("Day 4")
    Part1()
    Part2()


if __name__ == "__main__":
    main()
