import unittest

from two_pointers import two_sum


class TestTwoSum(unittest.TestCase):

    def test_two_sum(self):
        arr = [10, 20, 35, 50]
        target = 70
        self.assertTrue(two_sum(arr, target))

        arr = [10, 20, 30]
        target = 70
        self.assertFalse(two_sum(arr, target))

        arr = [-8, 1, 4, 6, 10, 45]
        target = 16
        self.assertTrue(two_sum(arr, target))


if __name__ == "__main__":
    unittest.main()
