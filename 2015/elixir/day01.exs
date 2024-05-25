defmodule Day01 do
  def data do
    Path.expand("aoc/2015/data/day01.txt", "~/projects")
    |> File.read!()
    |> String.split("", trim: true)
  end
end

defmodule Convert do
  def convert([]), do: []

  def convert([head | tail]) do
    item =
      if head == "(" do
        String.replace(head, "(", "1")
      else
        String.replace(head, ")", "-1", global: true)
      end
      |> String.to_integer()

    Enum.to_list([item | convert(tail)])
  end
end

defmodule Answer do
  def part_1() do
    Enum.sum(Convert.convert(Day01.data()))
    |> IO.inspect(label: "Part 1")
  end

  def part_2() do
  end
end

Answer.part_1()
