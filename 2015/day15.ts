import { readFileSync } from "fs";

const data = readFileSync("./inputs/day15.txt", "utf-8");

const INGREDIENT_RE =
  /^(?<name>[A-Za-z]+): capacity (?<cap>-?[0-9]+), durability (?<dur>-?[0-9]+), flavor (?<flavor>-?[0-9]+), texture (?<texture>-?[0-9]+), calories (?<cal>[0-9]+)/;

interface Ingredient {
  name: string;
  capacity: number;
  durability: number;
  flavor: number;
  texture: number;
  calories: number;
}

const INGREDIENTS: Map<string, Ingredient> = new Map();
for (const line of data.split("\n")) {
  const match = INGREDIENT_RE.exec(line);
  if (match && match.groups) {
    const ingredient: Ingredient = {
      name: match.groups["name"],
      capacity: parseInt(match.groups["cap"]),
      durability: parseInt(match.groups["dur"]),
      flavor: parseInt(match.groups["flavor"]),
      texture: parseInt(match.groups["texture"]),
      calories: parseInt(match.groups["cal"]),
    };
    INGREDIENTS.set(ingredient.name, ingredient);
  }
}

function score(choices: Map<string, number>): number {
  let capacity = 0;
  let durability = 0;
  let flavor = 0;
  let texture = 0;
  for (const [name, count] of choices) {
    const ingredient = INGREDIENTS.get(name)!;
    capacity += ingredient.capacity * count;
    durability += ingredient.durability * count;
    flavor += ingredient.flavor * count;
    texture += ingredient.texture * count;
  }
  if (capacity <= 0 || durability <= 0 || flavor <= 0 || texture <= 0) {
    return 0;
  }
  return capacity * durability * flavor * texture;
}

function calories(choices: Map<string, number>): number {
  let calories = 0;
  for (const [name, count] of choices) {
    const ingredient = INGREDIENTS.get(name)!;
    calories += ingredient.calories * count;
  }
  return calories;
}

const TOTAL_TSP = 100;

interface Recipe {
  choices?: Map<string, number>;
  score: number;
  calories: number;
}

function findBest(choices: Map<string, number>, calorieMatch?: number): Recipe {
  const total = [...choices.values()].reduce((acc, v) => acc + v, 0);
  const nextChoice = [...INGREDIENTS.keys()].find((v) => !choices.has(v))!;
  if (choices.size === INGREDIENTS.size - 1) {
    choices.set(nextChoice, TOTAL_TSP - total);
    return {
      choices: choices,
      score: score(choices),
      calories: calories(choices),
    };
  }
  let maximum: Recipe = { score: 0, calories: 0 };
  for (let i = 0; i < TOTAL_TSP - total; i++) {
    const newChoices = new Map(choices);
    newChoices.set(nextChoice, i);
    const recipe = findBest(newChoices, calorieMatch);
    if (!calorieMatch || recipe.calories === calorieMatch) {
      if (recipe.score > maximum.score) {
        maximum = recipe;
      }
    }
  }
  return maximum;
}

function part1() {
  console.log(findBest(new Map()));
}

function part2() {
  console.log(findBest(new Map(), 500));
}

console.log("DAY 15");
part1();
part2();
