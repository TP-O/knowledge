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
    let mut q = structure::Queue::<u8>::new();

    q.enqueue(1);
    q.enqueue(2);
    q.enqueue(3);
    q.enqueue(4);
    q.enqueue(5);

    loop {
        match q.next() {
            None => break,
            Some(data) => {
                println!("{}", data);
            }
        }
    }
}
