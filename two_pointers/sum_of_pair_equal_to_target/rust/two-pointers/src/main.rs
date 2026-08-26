// Sum of pair equal to target.
//
// Given a sorted array (asc) and a target, find if there exists any pair of elements (arr[i], arr[j])
// such that their sum is equal to the target.
//
// Input: arr = [10, 20, 35, 50]; target = 70
// Output: true ([20, 50] add up to 70)

fn two_sum(arr: Vec<i32>, target: i32) -> bool {
    let mut left = 0;
    let mut right = arr.len() - 1;

    while left < right {
        let sum = arr[left] + arr[right];

        if sum == target {
            return true;
        }

        if sum < target {
            left += 1;
        } else {
            right -= 1;
        }
    }

    return false;
}

fn main() {
    println!("Hello, world!");
}

#[cfg(test)]
mod tests {
    use crate::two_sum;

    #[test]
    fn test_two_sum() {
        let arr = vec![10, 20, 35, 50];
        let target = 70;

        assert!(two_sum(arr, target));

        let arr = vec![10, 20, 30];
        let target = 70;

        assert!(!two_sum(arr, target));

        let arr = vec![-8, 1, 4, 6, 10, 45];
        let target = 16;

        assert!(two_sum(arr, target));
    }
}
