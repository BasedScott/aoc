defmodule Day01 do
  def data do
    Path.expand("aoc/2015/data/day01.txt", "~/projects")
    |> File.read!()
    |> String.split("", trim: true)
  end
end

defmodule Convert do
  def convert([]), do: 0

  def convert([head | tail]) do
    if head == "(" do
      String.replace(head, "(", "1")
      |> String.to_integer()
      |> IO.puts()
    else
      String.replace(head, ")", "-1", global: true)
    end

    convert(tail)
  end
end
