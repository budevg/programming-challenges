
ans = squareOfTheSum - sumOfTheSquares
  where
    numbers = take 100 [1..]
    squareOfTheSum = (sum numbers) * (sum numbers)
    sumOfTheSquares = sum $ map (\x -> x*x) numbers


main = putStrLn $ show ans
