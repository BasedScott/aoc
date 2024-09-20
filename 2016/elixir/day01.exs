defmodule Day01 do
  @data Path.join(Path.expand(".."), "/data/day01.txt")

  def data do
    @data
    |> File.read!()
    |> String.split(", ", trim: true)
  end

  def part_1() do
    Day01.data()
    |> IO.inspect(label: "Part 1")
  end
end

Day01.part_1()
