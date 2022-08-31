mod searching;
mod sorting;
mod structure;

fn main() {
    // STACK
    // let mut s = structure::Stack::<u8>::new();

    // s.push(1);
    // s.push(2);
    // s.push(3);
    // s.push(4);
    // s.push(5);

    // loop {
    //     match s.next() {
    //         None => break,
    //         Some(data) => {
    //             println!("{}", data);
    //         }
    //     }
    // }

    // QUEUE
    // let mut q = structure::Queue::<u8>::new();

    // q.enqueue(1);
    // q.enqueue(2);
    // q.enqueue(3);
    // q.enqueue(4);
    // q.enqueue(5);

    // loop {
    //     match q.next() {
    //         None => break,
    //         Some(data) => {
    //             println!("{}", data);
    //         }
    //     }
    // }

    // BUBBLE SORT
    // let mut arr: [i32; 5] = [5, 1, 2, 6, -2];

    // sorting::bubble_sort(&mut arr, true);

    // for e in arr {
    //     println!("{}", e);
    // }

    // SELECTION SORT
    // let mut arr: [i32; 5] = [5, 1, 2, 6, -2];

    // sorting::selection_sort(&mut arr, false);

    // for e in arr {
    //     println!("{}", e);
    // }

    // SINGLY LINKED LIST
    // let mut sll = structure::SinglyLinkedList::<u8>::new();

    // sll.push(33);
    // sll.push(4);
    // sll.push(5);
    // sll.unshift(2);
    // sll.unshift(1);
    // sll.set(2, 3);

    // for i in 0..sll.size() {
    //     print!("{}, ", sll.get(i).unwrap());
    // }

    // println!();

    // sll.shift();
    // sll.pop();

    // for i in 0..sll.size() {
    //     print!("{}, ", sll.get(i).unwrap());
    // }

    // println!();

    // BINARY SEARCH
    // let arr: [i32; 6] = [1, 2, 3, 4, 5, 6];
    // let found = searching::binary_search(&arr, 5);

    // println!("Found at index: {}", found);

    // LINEAR SEARCH
    // let arr: [i32; 6] = [1, 2, 3, 4, 5, 6];
    // let found = searching::linear_search(&arr, 5);

    // println!("Found at index: {}", found);

    // INTERPOLATION SEARCH
    // let arr: [u128; 6] = [1, 2, 3, 4, 5, 6];
    // let found = searching::interpolation_search(&arr, 5);

    // println!("Found at index: {}", found);

    // JUMP SEARCH
    // let arr: [u128; 6] = [1, 2, 3, 4, 5, 6];
    // let found = searching::jump_search(&arr, 5);

    // println!("Found at index: {}", found);
}
