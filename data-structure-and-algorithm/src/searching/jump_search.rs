#[allow(dead_code)]
pub fn jump_search(arr: &[u128], val: u128) -> i128 {
    let len = arr.len();

    if val < arr[0] || val > arr[len - 1] {
        return -1;
    }

    let mut prev_position = 0;
    let mut position = 0;
    let step = (arr.len() as f64).sqrt() as usize;

    while position <= len {
        if arr[position] < val {
            prev_position = position;
            position = prev_position + step;
        } else if arr[position] > val {
            break;
        } else {
            return position as i128;
        }
    }

    for i in prev_position..position {
        if arr[i] == val {
            return i as i128;
        }
    }

    -1
}
