#[allow(dead_code)]
pub fn selection_sort<T: Ord>(arr: &mut [T], asc: bool) {
    let n = arr.len();

    for i in 0..n - 1 {
        let mut swapped_index = i;

        for j in i + 1..n {
            if should_update_swapped_index(&i, &j, asc) {
                swapped_index = j;
            }
        }

        arr.swap(i, swapped_index);
    }
}

fn should_update_swapped_index<T: Ord>(a: &T, b: &T, asc: bool) -> bool {
    if asc {
        a < b
    } else {
        a > b
    }
}
