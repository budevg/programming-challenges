use std::io::BufRead;
use regex::Regex;
use core::fmt::Display;
use std::fmt;
use std::collections::HashSet;

#[derive(Debug)]
struct Claim {
    left: u32,
    top: u32,
    width: u32,
    height: u32,
    id: u32,
}

impl Claim {
    fn new(s: &str) -> Claim {
        let re = Regex::new(r"#(\d+) @ (\d+),(\d+): (\d+)x(\d+)").unwrap();
        let caps = re.captures(s).unwrap();
        let left = caps.get(2).unwrap().as_str().parse().unwrap();
        let top = caps.get(3).unwrap().as_str().parse().unwrap();
        let width = caps.get(4).unwrap().as_str().parse().unwrap();
        let height = caps.get(5).unwrap().as_str().parse().unwrap();
        let id = caps.get(1).unwrap().as_str().parse().unwrap();

        // println!("{}", s);
        let claim = Claim{left, top, width, height, id};
        // println!("{:?}", claim);
        claim
    }
}

fn intersect(fibre: &mut Vec<Vec<u8>>, n: &Claim, m: &Claim) -> bool {
    let mut intersecting = false;
    for x in n.left .. n.left + n.width {
        for y in n.top .. n.top + n.height {
            if (x >= m.left) &&
                (x < m.left + m.width) &&
                (y >= m.top) &&
                (y < m.top + m.height) {
                    let i = x as usize;
                    let j = y as usize;
                    fibre[i][j] = 1;
                    intersecting = true;
                }
        }
    }
    intersecting
}

fn main() {
    let f = std::fs::File::open("input2.txt").expect("file not found");
    let file = std::io::BufReader::new(&f);
    let mut claims: Vec<Claim> = file
        .lines()
        .map(|l| Claim::new(&(l.unwrap())))
        .collect();
    let (width, height) = claims.iter().fold((0,0), |acc, x| {
        let w = std::cmp::max(acc.0, x.left + x.width);
        let h = std::cmp::max(acc.1, x.top + x.height);
        (w,h)
    });
    println!("width={}, height={}", width, height);
    let mut fibre = vec![vec![0; width as usize]; height as usize];
    let mut ids: HashSet<u32> = claims.iter().map(|x|x.id).collect();
    for i in 0..claims.len() {
        for j in i+1..claims.len() {
            if intersect(&mut fibre, &claims[i], &claims[j]) {
                ids.remove(&claims[i].id);
                ids.remove(&claims[j].id);
            }
        }
    }

    let mut sum = 0;
    for i in 0..fibre.len() {
        for j in 0..fibre[i].len() {
            if fibre[i][j] == 1 {
                sum += 1;
            }
        }
    }
    // println!("{:?}", fibre);
    println!("sum={}", sum);
    println!("ids={:?}", ids);


}
