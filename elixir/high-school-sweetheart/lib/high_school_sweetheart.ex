defmodule HighSchoolSweetheart do
  def first_letter(name), do: name |> String.trim() |> String.at(0) 

  def initial(name), do: name |> first_letter |> String.upcase() |> Kernel.<>(".")

  def initials(full_name) do
    [first_word, second_word] = String.split(full_name)

    "#{initial(first_word)} #{initial(second_word)}"
  end

  def pair(full_name1, full_name2) do
    pair_1 = initials(full_name1)
    par_2 = initials(full_name2)

    """
         ******       ******
       **      **   **      **
     **         ** **         **
    **            *            **
    **                         **
    **     #{pair_1}  +  #{par_2}     **
     **                       **
       **                   **
         **               **
           **           **
             **       **
               **   **
                 ***
                  *
    """

    # Please implement the pair/2 function
  end
end
