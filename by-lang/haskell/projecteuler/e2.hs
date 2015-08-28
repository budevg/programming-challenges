
fibs = 1 : 2 : map (\(a,b) -> a + b) (zip fibs $ tail fibs)

ans = sum $ filter even fibsBelow1M
  where
    fibsBelow1M = takeWhile (<=4000000) fibs
