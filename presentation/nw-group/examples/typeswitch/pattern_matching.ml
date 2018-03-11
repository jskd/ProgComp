type expr =
  | Literal of int
  | Add of expr * expr
  | Mult of expr * expr

let rec to_string = function
  | Literal i -> Printf.sprintf "%d" i
  | Add (left, right) -> to_string left ^ " + " ^ to_string right
  | Mult (left, right) -> to_string left ^ " * " ^ to_string right

let rec eval = function
  | Literal i -> i
  | Add (left, right) -> eval left + eval right
  | Mult (left, right) -> eval left * eval right


let main =
  let one = Literal 1 in
  let two = Literal 2 in
  let three = Literal 3 in
  let expr = Add (one, Mult (two, three)) in
  Printf.printf "%s = %d\n" (to_string expr) (eval expr)
