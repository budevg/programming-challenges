
threeDigitNumber = [x | x <- [100..999]]
candidates = [x*y |
              x <- threeDigitNumber,
              y <- threeDigitNumber,
              x /= y]

isPolindrom n  = isPolindromStr (show n)
  where
    isPolindromStr [] = True
    isPolindromStr [x] = True
    isPolindromStr [x,y] = x == y
    isPolindromStr (x:xs) = x == (last xs) &&
                            isPolindromStr (init xs)

ans = maximum $ [x | x <- candidates, isPolindrom x]
