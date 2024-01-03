module Main where

floorCount :: [Integer] -> Int
floorCount = length . countEach (/= -1) . scanl1 (+)
    where
        countEach :: (a -> Bool) -> [a] -> [a]
        countEach _ [] = []
        countEach pred (x:xs)
            | pred x    = x : countEach pred xs
            | otherwise = [x]

floorConvert :: String -> [Integer]
floorConvert = map replace -- floorConvert x = map replace x
    where
        replace '(' = 1
        replace ')' = -1

main :: IO ()
main = do
    floors <- readFile "home/lenin/projects/aoc/2015/data/day01.txt"
    print $ floorCount $ floorConvert floors
