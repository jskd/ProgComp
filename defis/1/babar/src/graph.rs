#![allow(dead_code)]

use cell;


enum Color
{
    White,
    Black,
    Grey,
    Red,
}

struct Node{
    value: Box<cell::Cell>,
    c: Color,
    child_list: Vec<Node>,
}

fn evaluate(node: Node)
{

}