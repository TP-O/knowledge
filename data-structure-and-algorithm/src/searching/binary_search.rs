#[allow(dead_code)]
pub fn binary_search<T: Ord>(arr: &[T], val: T) -> i128 {
    let mut left = 0;
    let mut right = arr.len() - 1;

    if val < arr[left] || val > arr[right] {
        return -1;
    }

    while left <= right {
        let middle = (right + left) / 2;

        if arr[middle] > val {
            right = middle - 1;
        } else if arr[middle] < val {
            left = middle + 1;
        } else {
            return middle as i128;
        }
    }

    -1
}
