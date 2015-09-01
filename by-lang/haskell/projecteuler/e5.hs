import Data.Numbers.Primes
import Data.List


ans = product factorsTaken
  where
    divRangeEnd = 20
    divRange = [1..divRangeEnd]
    factors = group $ sort $ concatMap  primeFactors divRange
    factorsTaken = map (foldl productBelowRangeEnd 1) factors
    productBelowRangeEnd = \acc x ->
      if acc * x < divRangeEnd then acc * x  else acc

main = do
  putStrLn (show ans)
