
ans = sum [x | x <- natBelow1000, x `mod` 3 == 0 || x `mod` 5 == 0]
  where
    natBelow1000 = takeWhile (<1000) [1..]
