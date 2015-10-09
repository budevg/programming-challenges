import Data.Array

lattice h w = table ! (h,w)
  where
    table = array ((1,1), (h,w))
            [((n, m), cell n m) | n <- [1..h], m <- [1..w]]
    cell 1 m = m + 1
    cell n 1 = n + 1
    cell n m = (table ! (n - 1, m)) + (table ! (n, m - 1))

ans = lattice 20 20
main = do
  putStrLn $ show ans
