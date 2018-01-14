(* Some utilities *)
module Util = struct
  (* Transform a string into a list of chars.  *)
  let rec chars = function
    | "" -> []
    | s -> String.get s 0 :: chars String.(sub s 1 (length s - 1))

  (* Apply [f] to the start value [acc] and all the lines of the given
     file.  *)
  let fold_file f acc filename =
    let rec aux f acc ic =
      match input_line ic with
      | line -> aux f (f acc line) ic
      | exception End_of_file ->
          close_in ic;
          acc
    in
    filename |> open_in |> aux f acc

  (* Transform the given list to a string by applying [to_string] to
     each element of the list.  *)
  let string_of_list to_string elements =
    elements |> List.map to_string |> String.concat "\n"
end

(* An anagram dictionary is just a mapping from a word to its known
   anagrams.  *)
module AnagramDict : sig
  type t

  (* Return an empty anagram dictionary for the given words.  *)
  val init : string list -> t

  (* Return a string representation of the anagram dictionary.  *)
  val to_string : t -> string

  (* Add the needed bindings for the given word in the anagram
     dictionary.  *)
  val enrich : t -> string -> t
end = struct
  module StringMap = Map.Make(String)

  type t = string list StringMap.t

  (* Normalize a word, i.e. return the list of its lower-case letters
     sorted alphabetically.  *)
  let normalize word =
    word |> String.lowercase |> Util.chars |> List.sort compare

  (* Return true if and only if [word] and [word'] are anagrams each
     other.  *)
  let anagram word word' = normalize word = normalize word'

  let init words =
    List.fold_left (fun acc word ->
        StringMap.add word [] acc
      ) StringMap.empty words


  let to_string anagram_dict =
    let entries =
      StringMap.bindings anagram_dict
      |> List.map (fun (word, anagrams) -> (word, List.sort compare anagrams))
    in
    let string_of_entry (word, anagrams) =
      word ^ ":" ^ (if anagrams = [] then "" else "\n")  ^
      Util.string_of_list (fun x -> x) anagrams
    in
    entries |> Util.string_of_list string_of_entry

  let enrich anagram_dict word =
    StringMap.mapi (fun word' anagrams ->
        if anagram word word' then word :: anagrams else anagrams
      ) anagram_dict
end

(* Entry point *)
let main =
  try
    let argc = Array.length Sys.argv in
    if argc < 2 then raise Exit else
      (* Initialize the anagram dictionary.  *)
      let anagram_dict =
        Array.sub Sys.argv 2 (argc - 2)
        |> Array.to_list
        |> AnagramDict.init
      in

      (* Populate the anagram dictionary.  *)
      let anagram_dict =
        Util.fold_file AnagramDict.enrich anagram_dict Sys.argv.(1)
      in

      (* Print the result.  *)
      anagram_dict |> AnagramDict.to_string |> print_endline
  with
  | Exit ->
      Printf.printf "Usage: %s <dict-file> [word...]\n" Sys.argv.(0);
      exit 1
  | Sys_error msg ->
      print_endline msg;
      exit 1
