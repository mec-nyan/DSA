import unittest

from bin_search import binary_search


class TestBinarySearch(unittest.TestCase):

    def test_binary_search_found(self):
        arr = [1, 2, 3, 4, 5]
        n = 4
        expected = 3
        actual = binary_search(arr, n)

        self.assertEqual(actual, expected, f"Want {expected} but got {actual}")

    def test_binary_search_not_found(self):
        arr = [1, 2, 3, 4, 5]
        n = 6
        expected = -1
        actual = binary_search(arr, n)

        self.assertEqual(actual, expected, f"Want {expected} but got {actual}")


if __name__ == "__main__":
    unittest.main()
