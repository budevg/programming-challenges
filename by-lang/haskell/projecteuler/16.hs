import Data.Char

ans = sum $ map (\c -> (ord c - ord '0')) (show ((2 :: Integer) ^ (1000 :: Integer)))
main = do
  putStrLn $ show ans
