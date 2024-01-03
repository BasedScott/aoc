{- cabal:
	build-depends: base, split, regex-tdfa
-}
module Main where

main :: IO()
main = do
    boxes <- readFile "home/lenin/projects/aoc/2015/data/day02.txt"
    putStrLn boxes
