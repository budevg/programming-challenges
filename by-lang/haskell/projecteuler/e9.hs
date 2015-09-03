
ans = let [(a,b,c)] = take 1 triplets in a * b * c
  where triplets = [(a,b,c) |
                    a <- [1..1000],
                    b <- [1..1000],
                    let c = 1000 - a - b,
                    a < b, b < c,
                    a * a + b * b == c * c]

main = putStrLn (show ans)
