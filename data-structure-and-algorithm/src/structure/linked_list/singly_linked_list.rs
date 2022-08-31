use std::ptr::NonNull;

#[derive(Debug)]
pub struct Node<T> {
    data: T,
    next: NextNode<T>,
}

type NextNode<T> = Option<NonNull<Node<T>>>;

impl<T> Node<T> {
    pub fn new(data: T) -> Self {
        Node { data, next: None }
    }
}

pub struct SinglyLinkedList<T> {
    size: u128,
    head: NextNode<T>,
}

#[allow(dead_code)]
impl<T> SinglyLinkedList<T> {
    pub fn new() -> Self {
        SinglyLinkedList {
            size: 0,
            head: None,
        }
    }

    pub fn size(&self) -> u128 {
        self.size
    }

    pub fn first(&self) -> Option<&T> {
        self.get(0)
    }

    pub fn last(&self) -> Option<&T> {
        match self.size {
            0 => None,
            _ => self.get(self.size - 1),
        }
    }

    pub fn get(&self, index: u128) -> Option<&T> {
        match index >= self.size {
            true => None,
            false => self.get_from(self.head, index),
        }
    }

    fn get_from(&self, node: NextNode<T>, index: u128) -> Option<&T> {
        match node {
            None => None,
            Some(node_ptr) => match index {
                0 => Some(unsafe { &(*node_ptr.as_ptr()).data }),
                _ => self.get_from(unsafe { (*node_ptr.as_ptr()).next }, index - 1),
            },
        }
    }

    pub fn set(&mut self, index: u128, data: T) -> bool {
        match index >= self.size {
            true => false,
            false => self.set_from(self.head, index, data),
        }
    }

    fn set_from(&mut self, node: NextNode<T>, index: u128, data: T) -> bool {
        match node {
            None => false,
            Some(node_ptr) => match index {
                0 => unsafe {
                    (*node_ptr.as_ptr()).data = data;

                    true
                },
                _ => self.set_from(unsafe { (*node_ptr.as_ptr()).next }, index - 1, data),
            },
        }
    }

    pub fn unshift(&mut self, data: T) {
        let mut node = Box::new(Node::new(data));

        if self.head.is_some() {
            node.next = self.head.take();
        }

        self.head = Some(unsafe { NonNull::new_unchecked(Box::into_raw(node)) });
        self.size += 1;
    }

    pub fn push(&mut self, data: T) {
        self.insert(self.size, data)
    }

    pub fn insert(&mut self, index: u128, data: T) {
        if index > self.size {
            return;
        }

        match index {
            0 => self.unshift(data),
            _ => self.insert_from(self.head, index, data),
        }
    }

    fn insert_from(&mut self, node: NextNode<T>, index: u128, data: T) {
        match node {
            None => (),
            Some(node_ptr) => match index - 1 {
                0 => {
                    let mut inserted_node = Box::new(Node::new(data));

                    unsafe {
                        inserted_node.next = (*node_ptr.as_ptr()).next;
                        (*node_ptr.as_ptr()).next =
                            Some(NonNull::new_unchecked(Box::into_raw(inserted_node)));
                    }

                    self.size += 1;
                }
                _ => self.insert_from(unsafe { (*node_ptr.as_ptr()).next }, index - 1, data),
            },
        }
    }

    pub fn shift(&mut self) -> Option<T> {
        self.head.map(|head_ptr| {
            let old_head = unsafe { Box::from_raw(head_ptr.as_ptr()) };

            self.head = old_head.next;
            self.size -= 1;

            old_head.data
        })
    }

    pub fn pop(&mut self) -> Option<T> {
        match self.size {
            0 => None,
            _ => self.remove(self.size - 1),
        }
    }

    pub fn remove(&mut self, index: u128) -> Option<T> {
        match index {
            0 => self.shift(),
            _ => self.remove_from(self.head, index),
        }
    }

    fn remove_from(&mut self, node: NextNode<T>, index: u128) -> Option<T> {
        match node {
            None => None,
            Some(node_ptr) => match index - 1 {
                0 => {
                    // This node will be removed after function ends.
                    let removed_node = unsafe {
                        Box::from_raw((*node_ptr.as_ptr()).next.take().unwrap().as_ptr())
                    };

                    unsafe {
                        (*node_ptr.as_ptr()).next = removed_node.next;
                    }

                    self.size -= 1;

                    Some(removed_node.data)
                }
                _ => self.remove_from(unsafe { (*node_ptr.as_ptr()).next }, index - 1),
            },
        }
    }
}
