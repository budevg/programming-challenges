import Data.Numbers.Primes

ans = last $ take 10001 primes
main = putStrLn (show ans)
