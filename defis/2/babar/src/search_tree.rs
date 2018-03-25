#[derive(PartialEq)]
pub struct NodeSearchTree<'a> {
    pub value: &'a str,
    pub left: Option<Box<NodeSearchTree<'a>>>,
    pub right: Option<Box<NodeSearchTree<'a>>>,
}

impl<'a> NodeSearchTree<'a> {
    ///insert return false if value already in tree.
    /// /!\ WARNING /!\
    /// NOT CURRENTLY USED, BUT WILL BE SOON™
    #[allow(dead_code)]
    pub fn insert(&mut self, new_value: &'a str) -> bool {
        if self.value == new_value {
            return false;
        }
        let has_node = if self.value < new_value {
            &mut self.left
        } else {
            &mut self.right
        };
        match has_node {
            &mut Some(ref mut subnode) => return subnode.insert(new_value),
            &mut None => {
                let new_node = NodeSearchTree {
                    value: new_value,
                    left: None,
                    right: None,
                };
                let boxed_node = Some(Box::new(new_node));
                *has_node = boxed_node;
                return true;
            }
        }
    }
}
