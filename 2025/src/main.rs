use std::fs;

use clap::Parser;

use aoc2025::day_util::DaySolution;
use aoc2025::day_1;
use aoc2025::day_2;
use aoc2025::day_3;
use aoc2025::day_4;
use aoc2025::day_5;
use aoc2025::day_6;
use aoc2025::day_7;
use aoc2025::day_8;
use aoc2025::day_9;
use aoc2025::day_10;
use aoc2025::day_11;
use aoc2025::day_12;

/// Search for a pattern in a file and display the lines that contain it.
#[derive(Parser)]
struct AocCli {
    /// Should run on real data?
    #[arg(short='f', long="full", default_value="false")]
    full: bool,

    /// The day to run.
    day: u8,
}

fn main() {
    let args = AocCli::parse();
    let input_file = format!("inputs/day{}{}.txt", args.day, if args.full { "" } else {"example"});
    let text = fs::read_to_string(input_file)
        .expect("Should have been able to read the file");
    
    let d_opt: Option<Box<dyn DaySolution>> = match args.day {
        1 => Some(Box::new(day_1::Day {
            data: text,
        })),
        2 => Some(Box::new(day_2::Day{
            data: text,
        })),
        3 => Some(Box::new(day_3::Day{
            data: text,
        })),
        4 => Some(Box::new(day_4::Day{
            data: text,
        })),
        5 => Some(Box::new(day_5::Day{
            data: text,
        })),
        6 => Some(Box::new(day_6::Day{
            data: text,
        })),
        7 => Some(Box::new(day_7::Day{
            data: text,
        })),
        8 => Some(Box::new(day_8::Day{
            data: text,
        })),
        9 => Some(Box::new(day_9::Day{
            data: text,
        })),
        10 => Some(Box::new(day_10::Day{
            data: text,
        })),
        11 => Some(Box::new(day_11::Day{
            data: text,
        })),
        12 => Some(Box::new(day_12::Day{
            data: text,
        })),
        _ => None,
    };
    let d = d_opt.expect("Expected day between 1 and 12");
    println!("{}", d.part1());
    println!("{}", d.part2());
}