defmodule Day02 do
  @data Path.join(Path.expand(".."), "/data/day02.txt")
        |> File.read!()
        |> String.split()

  def part_1() do
    ParsePart1.parse(@data)
    |> Enum.sum()
    |> IO.inspect(label: "Part 1")
  end

  def part_2() do
    ParsePart2.parse(@data)
    |> Enum.sum()
    |> IO.inspect(label: "Part 2")
  end
end

defmodule ParsePart1 do
  def parse([]), do: []

  def parse([head | tail]) do
    l =
      Regex.split(~r{[x]}, head)
      |> Enum.at(0)
      |> String.to_integer()

    w =
      Regex.split(~r{[x]}, head)
      |> Enum.at(1)
      |> String.to_integer()

    h =
      Regex.split(~r{[x]}, head)
      |> Enum.at(2)
      |> String.to_integer()

    side_1 = l * w
    side_2 = w * h
    side_3 = h * l

    extra_paper = Enum.min([side_1, side_2, side_3])

    area_1 = side_1 * 2
    area_2 = side_2 * 2
    area_3 = side_3 * 2

    box = Enum.sum([area_1 + area_2 + area_3 + extra_paper])

    Enum.to_list([box | parse(tail)])
  end
end

defmodule ParsePart2 do
  def parse([]), do: []

  def parse([head | tail]) do
    l =
      Regex.split(~r{[x]}, head)
      |> Enum.at(0)
      |> String.to_integer()

    w =
      Regex.split(~r{[x]}, head)
      |> Enum.at(1)
      |> String.to_integer()

    h =
      Regex.split(~r{[x]}, head)
      |> Enum.at(2)
      |> String.to_integer()

    bow = l * w * h

    ribbon =
      Enum.sort([l, w, h], :asc)
      |> Enum.slice(0, 2)
      |> Enum.map(fn x -> x + x end)
      |> Enum.sum()

    box = Enum.sum([bow, ribbon])

    Enum.to_list([box | parse(tail)])
  end
end

Day02.part_1()
Day02.part_2()
