defmodule Lasagna do
  def expected_minutes_in_oven, do: 40

  def remaining_minutes_in_oven(t), do: expected_minutes_in_oven() - t

  def preparation_time_in_minutes(number_of_layers), do: 2 * number_of_layers

  def total_time_in_minutes(number_of_layers, elapsed_bake_time) do
    preparation_time_in_minutes(number_of_layers) + elapsed_bake_time
  end

  def alarm, do: "Ding!"
end
