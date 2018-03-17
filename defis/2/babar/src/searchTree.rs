#![allow(dead_code)]

#[derive(PartialEq)]
pub struct NodeST<'a>{
	value: &'a str,
	left: Option<Box<NodeST<'a>>>,
	right: Option<Box<NodeST<'a>>>,
}


impl<'a> NodeST<'a>{
	///insert return false if value already in tree.
	pub fn insert(&mut self, new_value: &'a str) -> bool
	{	
		if self.value == new_value{
			return false
		}
		let has_node = if self.value < new_value {&mut self.left} else {&mut self.right};
		match has_node {
			&mut Some(ref mut subnode) => return subnode.insert(new_value),
			&mut None => {
				let new_node = NodeST {value: new_value, left: None, right: None};
				let boxed_node = Some(Box::new(new_node));
				*has_node =  boxed_node;
				return true
			}
		}
	}
}