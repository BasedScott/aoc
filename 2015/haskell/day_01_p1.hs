module Main where

floorConvert :: String -> Int
floorConvert = sum . map replace -- floorConvert x = sum AND map replace x
    where
        replace '(' = 1
        replace ')' = -1

main :: IO ()
main = do
    floors <- readFile "/home/lenin/projects/aoc/2015/data/day01.txt"
    print $ floorConvert floors
