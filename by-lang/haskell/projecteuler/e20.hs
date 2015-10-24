
fac 0 = 1
fac n = n * fac (n-1)

sumDigits n
  | n < 9 = n
  | otherwise = m + sumDigits d
  where
    m = n `mod` 10
    d = n `div` 10


ans = sumDigits $ fac 100
main = do
  putStrLn $ show ans
