defmodule KitchenCalculator do

  @conversion_factors %{
    :cup => 240,
    :fluid_ounce => 30,
    :teaspoon => 5,
    :tablespoon => 15,
    :milliliter => 1,
  }

  def get_volume({_unit, volume}), do: volume 

  def to_milliliter({unit, volume}), do: {:milliliter, @conversion_factors[unit] * volume}

  def from_milliliter({:milliliter, volume}, unit), do: {unit, volume / @conversion_factors[unit]}

  def convert(volume_pair, unit_to_convert) when tuple_size(volume_pair) == 2 do
    volume_pair
    |> to_milliliter()
    |> from_milliliter(unit_to_convert)
  end
end
