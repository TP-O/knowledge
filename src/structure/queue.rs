use std::ptr::NonNull;

pub struct Node<T> {
    data: T,
    next: NextNode<T>,
}

pub type NextNode<T> = Option<NonNull<Node<T>>>;

impl<T> Node<T> {
    pub fn new(data: T) -> Self {
        Node { data, next: None }
    }
}

pub struct Queue<T> {
    size: u128,
    head: NextNode<T>,
    tail: NextNode<T>,
}

#[allow(dead_code)]
impl<T> Queue<T> {
    pub fn new() -> Self {
        return Queue {
            size: 0,
            head: None,
            tail: None,
        };
    }

    pub fn size(&self) -> u128 {
        self.size
    }

    pub fn is_empty(&self) -> bool {
        self.head.is_none()
    }

    pub fn enqueue(&mut self, data: T) {
        let node = Box::new(Node::new(data));
        let node_ptr = Some(unsafe { NonNull::new_unchecked(Box::into_raw(node)) });

        match self.tail {
            None => {
                self.head = node_ptr;
            }
            Some(tail) => unsafe {
                (*tail.as_ptr()).next = node_ptr;
            },
        }

        self.tail = node_ptr;
        self.size += 1;
    }

    pub fn dequeue(&mut self) -> Result<T, &str> {
        match self.head.take() {
            None => Err("Empty queue!"),
            Some(head) => unsafe {
                let old_head = Box::from_raw(head.as_ptr());

                self.head = old_head.next;
                self.size -= 1;

                Ok(old_head.data)
            },
        }
    }
}

impl<T> Iterator for Queue<T> {
    type Item = T;

    fn next(&mut self) -> Option<Self::Item> {
        let first = self.dequeue();

        first.ok()
    }
}
