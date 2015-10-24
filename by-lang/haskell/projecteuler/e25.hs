
fibs = 1 : 1 : zipWith (+) fibs (tail fibs)

ans = 1 + (length $ takeWhile (\n -> (length (show n)) < 1000) fibs)

main = do
  putStrLn $ show ans
