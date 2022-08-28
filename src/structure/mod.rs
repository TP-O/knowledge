pub mod stack;

pub use stack::Stack;

pub struct Node<T> {
    data: T,
    next: NextNode<T>,
}

pub type NextNode<T> = Option<Box<Node<T>>>;
