import Data.Numbers.Primes

target = 600851475143
maxFactor 1 = 1
maxFactor n = max firstFactor (maxFactor (n `div` firstFactor))
  where
    [firstFactor] = take 1 [x | x <- primes, n `mod` x == 0]

ans = maxFactor target
