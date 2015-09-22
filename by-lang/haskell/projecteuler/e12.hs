import Data.Numbers.Primes
import Data.List

triangleNumbers = helper 0 1
  where helper s x = s + x : helper (s+x) (x+1)

factors n = foldl (\acc x -> acc * ((length x) + 1)) 1 groupedPrimeFactors
  where groupedPrimeFactors = groupBy (==) $ primeFactors n

ans = take 1 [x | x <- triangleNumbers, factors x > 500]
main = putStrLn (show ans)
