{- cabal:
	build-depends: base, diagrams-lib
-}
module Main where
import Diagrams.Coordinates
import Data.Char (Char)

arrows :: String -> (Int,Int)
arrows directions = zip directions [(x, y) | y <- [0..], x <- [0..length (lines directions) - 1]]

main :: IO()
main = do
    directions <- readFile "/home/lenin/projects/aoc/2015/data/day03.txt"
    let arrowsMapped = arrows directions
    print arrowsMapped
