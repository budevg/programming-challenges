use std::io::BufRead;
use std::collections::HashMap;

fn count_letters(line: &str) -> HashMap<char, u32> {
    let mut count_map: HashMap<char, u32> = HashMap::new();
    for c in line.chars() {
        let counter = count_map.entry(c).or_insert(0);
        *counter += 1;
    }
    count_map

}

fn label_diff(a: &str, b: &str) -> String {
    let mut out = String::new();
    for (a,b) in a.chars().zip(b.chars()) {
        if a == b {
            out.push(a)
        }
    }
    out

}
fn stage2() {
    let f  = std::fs::File::open("input.txt").expect("file not found");
    let file = std::io::BufReader::new(&f);
    let labels: Vec<String> = file.lines().map(|l|l.unwrap()).collect();
    'exit: for i in 0..labels.len() {
        for j in i + 1 .. labels.len() {
            let diff = label_diff(&labels[i], &labels[j]);
            if diff.len() == labels[i].len() - 1 {
                println!("label: {}", diff);
                break 'exit;
            }

        }
    }
    println!("done!");
}

fn stage1() {
    let f  = std::fs::File::open("input.txt").expect("file not found");
    let file = std::io::BufReader::new(&f);
    let mut contains_two = 0;
    let mut contains_three = 0;
    for line in file.lines().map(|l|l.unwrap()) {
        let count_map = count_letters(&line);
        for v in count_map.values() {
            if *v == 2 {
                contains_two += 1;
                break;
            }
        }
        for v in count_map.values() {
            if *v == 3 {
                contains_three += 1;
                break;
            }
        }
    }
    println!("checksum={}", contains_two * contains_three);
}

fn main() {
    stage1();
    stage2();
}
