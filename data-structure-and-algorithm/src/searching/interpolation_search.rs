#[allow(dead_code)]
pub fn interpolation_search(arr: &[u128], val: u128) -> i128 {
    let mut low = 0;
    let mut high = arr.len() - 1;

    while low <= high && val >= arr[low] && val <= arr[high] {
        // low == high causes position calculation to fail,
        // so we handle this case here
        if low == high {
            if arr[low] == val {
                return low as i128;
            }

            break;
        }

        let position = low + (high - low) / (((arr[high] - arr[low]) * (val - arr[low])) as usize);

        if arr[position] > val {
            high = position - 1;
        } else if arr[position] < val {
            low = position + 1;
        } else {
            return position as i128;
        }
    }

    -1
}
