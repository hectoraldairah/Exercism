defmodule NameBadge do
  def print(id, name, department) do
     show_id = if id == nil, do: "", else: "[#{id}] - "
     show_deparment = if department == nil, do: "OWNER", else: "#{String.upcase(department)}"
    "#{show_id}#{name} - #{show_deparment}"
  end
end
