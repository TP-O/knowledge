mod structure;

fn main() {
    let mut s = structure::Stack::<u8>::new();

    s.push(1);
    s.push(2);
    s.push(3);
    s.push(4);
    s.push(5);

    loop {
        match s.next() {
            None => break,
            Some(data) => {
                println!("{}", data);
            }
        }
    }
}
