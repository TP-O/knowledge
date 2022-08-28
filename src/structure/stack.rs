use crate::structure::NextNode;
use crate::structure::Node;

pub struct Stack<T> {
    size: u128,
    head: NextNode<T>,
}

#[allow(dead_code)]
impl<T> Stack<T> {
    pub fn new() -> Self {
        return Stack {
            size: 0,
            head: None,
        };
    }

    pub fn size(&self) -> u128 {
        self.size
    }

    pub fn is_empty(&self) -> bool {
        self.head.is_none()
    }

    pub fn peek(&self) -> Option<&T> {
        match self.head.as_ref() {
            None => None,
            Some(node) => Some(&node.data),
        }
    }

    pub fn push(&mut self, data: T) {
        let new_node = Node {
            data,
            next: self.head.take(),
        };

        self.size += 1;
        self.head = Some(Box::new(new_node));
    }

    pub fn pop(&mut self) -> Result<T, &str> {
        match self.head.take() {
            None => Err("Empty stack!"),
            Some(node) => {
                self.size -= 1;
                self.head = node.next;

                Ok(node.data)
            }
        }
    }
}

impl<T> Iterator for Stack<T> {
    type Item = T;

    fn next(&mut self) -> Option<Self::Item> {
        let peek = self.pop();

        peek.ok()
    }
}
