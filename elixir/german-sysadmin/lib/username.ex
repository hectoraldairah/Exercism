defmodule Username do
  def sanitize([]), do: []

  def sanitize([h | t]) do
    case h do
      ?\s -> sanitize(t)
      ?_ -> [h] ++ sanitize(t)
      h when h in ?a..?z -> [h] ++ sanitize(t) 
      h when h in [?ä, ?ö, ?ü, ?ß] -> check_german_letters(h) ++ sanitize(t) 
      _  -> sanitize(t)
    end
  end

  defp check_german_letters(letter) do
    case letter do
      ?ä -> [97, 101]
      ?ö -> [111, 101]
      ?ü -> [117, 101]
      ?ß -> [115, 115]
    end
  end
end
