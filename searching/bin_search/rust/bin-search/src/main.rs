fn main() {
    println!("Hello, world!");
}

fn bin_search(elems: &Vec<i32>, n: i32) -> std::result::Result<i32, &str> {
    let mut start = 0;
    let mut end = elems.len() - 1;

    while start <= end {
        let middle = (end - start) / 2 + start;
        let current = elems[middle];

        if current == n {
            return Ok(middle as i32);
        }

        if current > n {
            end = middle - 1;
        } else {
            start = middle + 1;
        }
    }

    return Err("Not found.");
}

#[cfg(test)]
mod tests {
    use crate::bin_search;

    #[test]
    fn test_bin_search() {
        let elems = vec![1, 2, 3, 4, 5, 6];
        let n = 4;

        assert_eq!(bin_search(&elems, n), Ok(3));

        let n = 20;
        assert_eq!(bin_search(&elems, n), Err("Not found."));
    }
}
