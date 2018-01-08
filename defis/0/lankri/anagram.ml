(* Transform a string into a list of chars.  *)
let rec chars = function
  | "" -> []
  | s when String.length s = 1 -> [String.get s 0]
  | s ->
      String.get s 0 :: chars (String.sub s 1 (String.length s - 1))

(* Normalize a word, i.e. return the list of its chars sorted in
   alphabetically.  *)
let normalize word = word |> chars |> List.sort compare

(* Return true iff [word] and [word'] are anagrams each other.  *)
let anagram word word' = normalize word = normalize word'

let ( =?= ) = anagram

let _ =
  try
    let in_ch = open_in Sys.argv.(1) in
    ();
    close_in in_ch
  with
  | Invalid_argument _ -> exit(1)
