defmodule Day01 do
  def data do
    Path.expand("aoc/2015/data/day01.txt", "~/projects")
    |> File.read!()
    |> String.split("", trim: true)
  end

  def sum do
  end
end
