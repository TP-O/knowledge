#[allow(dead_code)]
pub fn linear_search<T: Ord>(arr: &[T], val: T) -> i128 {
    let len = arr.len();

    for i in 0..len {
        if arr[i] == val {
            return i as i128;
        }
    }

    -1
}
