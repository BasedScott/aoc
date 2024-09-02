defmodule Day03 do
  @data Path.join(Path.expand(".."), "/data/day03.txt")
        |> File.read!()
        |> String.split("", trim: true)

  def part_1() do
    # IO.inspect(@data)
    @data
    |> Parse.parse()
    |> Parse.path()
  end
end

defmodule Parse do
  def parse([]), do: []

  def parse([head | tail]) do
    cord = Convert.cord(head)

    Enum.to_list([cord | parse(tail)])
  end
  def path([]), do: []
  def path([head | tail]) do

  end
end

defmodule Convert do
  def cord("^"), do: [y: 1]
  def cord(">"), do: [x: 1]
  def cord("v"), do: [y: -1]
  def cord("<"), do: [x: -1]
  def step(:x), do: [x: ]
  def step(:y), do: [y: ]
end

IO.inspect(Day03.part_1())
