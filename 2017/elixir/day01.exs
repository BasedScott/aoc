defmodule Day01 do
  @data Path.join(Path.expand(".."), "/data/day01.txt")
        |> File.read!()
        |> String.split("", trim: true)
  def answer_part1([]), do: []

  def answer_part1([@data, second | tail]) do
    if first = second do
      Day01.answer_part1(tail, [first, second])
    else
      Day01.answer_part1(tail)
    end
  end
end

Day01.answer_part1([])
