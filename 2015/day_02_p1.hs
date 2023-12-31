{- cabal:
	build-depends: base, regex-base
-}
module Main where
import Text.Regex.*

main :: IO()
main = do
    boxes <- readFile "day_02_rd.txt"
    putStrLn boxes