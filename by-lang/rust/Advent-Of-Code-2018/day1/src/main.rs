use std::collections::HashSet;
use std::fs::File;
use std::io::{BufRead, BufReader};

fn stage1() {
    let f = File::open("input.txt").expect("file not found");
    let file = BufReader::new(f);
    let mut frequency: i32 = 0;
    for line in file.lines().map(|line| line.unwrap()) {
        let change: i32 = line.parse().expect("failed to parse line");
        frequency += change;
    }
    println!("stage1, frequency={}", frequency);
}

fn stage2() {
    let mut frequency: i32 = 0;
    let mut frequency_map: HashSet<i32> = HashSet::new();
    'exit: loop {
        let f = File::open("input2.txt").expect("file not found");
        let file = BufReader::new(f);
        for line in file.lines().map(|line| line.unwrap()) {
            let change: i32 = line.parse().expect("failed to parse line");
            frequency += change;
            if frequency_map.contains(&frequency) {
                println!("stage2, frequency={}", frequency);
                break 'exit;
            }
            frequency_map.insert(frequency);
        }
    }
}

fn main() {
    stage1();
    stage2();
}
