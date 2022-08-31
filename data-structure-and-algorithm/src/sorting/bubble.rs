#[allow(dead_code)]
pub fn bubble_sort<T: Ord>(arr: &mut [T], asc: bool) {
    let mut n = arr.len();
    let mut is_sorted = false;

    // Array is sorted if it remains unchanged in 1 round.
    while !is_sorted {
        is_sorted = true;

        for i in 0..n - 1 {
            if should_swap(&arr[i], &arr[i + 1], asc) {
                arr.swap(i, i + 1);

                is_sorted = false;
            }
        }

        n -= 1;
    }
}

fn should_swap<T: Ord>(a: &T, b: &T, asc: bool) -> bool {
    if asc {
        a > b
    } else {
        a < b
    }
}
