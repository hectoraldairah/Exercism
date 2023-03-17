defmodule BirdCount do
  def today(list), do: List.first(list) 

  def increment_day_count([]), do: [1]
  def increment_day_count([h | t]), do: [h + 1 | t]

  def has_day_without_birds?([]), do: false
  def has_day_without_birds?([0 | _t]), do: true
  def has_day_without_birds?([_ | t]), do: has_day_without_birds?(t)

  def sum(acc, x), do: acc + x

  def total(list) do
    Enum.reduce(list, 0, &sum/2)
  end

  def busy_days(list) do
    list
    |> Enum.filter(fn x -> x >= 5 end)
    |> length
  end
end
