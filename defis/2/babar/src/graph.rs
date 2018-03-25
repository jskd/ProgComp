#![allow(dead_code)]

use cell;

pub enum Color {
    White,
    Black,
    Grey,
    Red,
}

pub struct Node<'a> {
    pub value: Box<&'a cell::Formula>,
    pub c: Color,
    pub child_list: Vec<Node<'a>>,
}

fn evaluate(mut node: Node) {
    match node.c {
        Color::White => {
            node.c = Color::Grey;
            for mut n in node.child_list {
                match n.c {
                    Color::White => {
                        println!("child white, evaluate(child)");
                        evaluate(n);
                    }
                    Color::Black => {
                        println!("child black");
                    }
                    _ => {
                        node.c = Color::Red;
                        println!("child red or grey, this node goes red");
                    }
                }
            }
        }
        _ => println!("this node red or black -> no look children"),
    }

    //evaluate content of cell if not already done
    match node.c {
        Color::Grey | Color::White => {
            // charger en mémoire la zone de la formule on est pas rouge donc on peut la calculer
            println!("evaluate this node value");
            node.c = Color::Black; // on, a évaluer on passe le node en noir
        }
        Color::Red => println!("this node value = P"),
        _ => println!("already evaluate"), //normalement case inaccessible.
    }
}

/*fn main(){
	let mut child = Vec::new();
	for i in 0..10{
		if i < 4 {
			child.push(Node{value:i, c:Color::White, child_list: Vec::new()});
		}
		else if i < 7 {
			child.push(Node{value:i, c:Color::Red, child_list: Vec::new()});
		}
		else {
			child.push(Node{value:i, c:Color::Black, child_list: Vec::new()});
		}
	}
	let mut node = Node {value: 11, c:Color::White, child_list: child};

	evaluate(node);
}*/
