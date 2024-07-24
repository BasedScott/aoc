defmodule Day03 do
  @data Path.join(Path.expand(".."), "/data/day03.txt")
        |> File.read!()
        |> String.split("", trim: true)

  def part_1() do
    # IO.inspect(@data)
    ParsePart1.parse(@data)
  end
end

defmodule ParsePart1 do
  def parse([]), do: []

  def parse([head | tail]) do
    map = [
      %{"^" => [x: 1]},
      %{">" => [y: 1]},
      %{"v" => [x: -1]},
      %{"<" => [y: -1]}
    ]

    cord = Map.values(map)

    Enum.to_list([cord | parse(tail)])
  end
end

Day03.part_1()
