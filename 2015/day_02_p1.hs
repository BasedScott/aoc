{- cabal:
	build-depends: base, split, regex-tdfa
-}
module Main where

main :: IO()
main = do
    boxes <- readFile "day_02_rd.txt"
    putStrLn boxes
