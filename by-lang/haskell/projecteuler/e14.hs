import Data.Array
import Data.List

collatzSeq :: Integer -> [[Integer]]
collatzSeq n = map collatzSeq' [1..n]
  where
    memoLimit = 10000
    memoArray = array (1, memoLimit) [(i, collatzSeq'' i) |
                                      i <- [1..memoLimit]]
    collatzSeq' n
      | n <= memoLimit = memoArray ! n
      | otherwise = collatzSeq'' n
    collatzSeq'' 1 = [1]
    collatzSeq'' n
      | even n = n : collatzSeq' (n `div` 2)
      | otherwise = n : collatzSeq' (3 * n + 1)


ans = foldl1' (\xs ys -> if length xs > length ys
                         then xs
                         else ys) (collatzSeq 1000000)

main = do
  putStrLn $ show (ans !! 0)
