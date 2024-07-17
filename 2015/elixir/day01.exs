defmodule Day01 do
  @data Path.join(Path.expand(".."), "/data/day01.txt")

  def data do
    @data
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

defmodule Count do
  @answer Convert.convert(Day01.data())
  def count([head | tail], acc) when acc > -1 do
    acc = head + acc
    count(tail, acc)
  end

  def count(tail, _) do
    pos = Enum.count(tail) - Enum.count(@answer)
    Kernel.abs(pos)
  end
end

defmodule Answer do
  @answer Convert.convert(Day01.data())
  def part_1() do
    Enum.sum(@answer)
    |> IO.inspect(label: "Part 1")
  end

  def part_2() do
    Count.count(@answer, 0)
    |> IO.inspect(label: "Part 2")
  end
end

Answer.part_1()
Answer.part_2()
