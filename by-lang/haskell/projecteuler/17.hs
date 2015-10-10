namedNumbers = ["zero", "one", "two", "three", "four", "five",
                "six", "seven", "eight", "nine", "ten", "eleven",
                "twelve", "thirteen", "fourteen", "fifteen",
                "sixteen", "seventeen", "eighteen", "nineteen", "twenty"]

tens = ["", "ten", "twenty", "thirty", "forty",
        "fifty", "sixty", "seventy", "eighty", "ninety"]

numToName n
  | n <= 20 = namedNumbers !! n
  | n < 100 = let n1 = n `div` 10
                  n0 = n `mod` 10
              in if n0 == 0
                 then
                   tens !! n1
                 else
                   (tens !! n1) ++ " " ++ (numToName n0)
  | n < 1000 = let n1 = n `div` 100
                   n0 = n `mod` 100
                   hundreds = (numToName n1) ++ " hundred"
               in if n0 == 0
                  then
                    hundreds
                  else
                    hundreds ++ " and " ++ (numToName n0)
  | n == 1000 = "one thousand"


ans = sum [countLetters $ numToName n | n <- [1..1000]]
  where
    countLetters s = sum [1 | c <- s, c /= ' ']

main = do
  putStrLn $ show ans
