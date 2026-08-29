import unittest

from two_sum import two_sum_1, two_sum_2


class TestTwoSum(unittest.TestCase):

    def test_two_sum_1(self):
        arr = [10, 20, 35, 50]
        target = 70
        self.assertTrue(two_sum_1(arr, target))

        arr = [10, 20, 30]
        target = 70
        self.assertFalse(two_sum_1(arr, target))

        arr = [-8, 1, 4, 6, 10, 45]
        target = 16
        self.assertTrue(two_sum_1(arr, target))

    def test_two_sum_2(self):
        arr = [10, 20, 35, 50]
        target = 70
        self.assertTrue(two_sum_2(arr, target))

        arr = [10, 20, 30]
        target = 70
        self.assertFalse(two_sum_2(arr, target))

        arr = [-8, 1, 4, 6, 10, 45]
        target = 16
        self.assertTrue(two_sum_2(arr, target))


if __name__ == "__main__":
    unittest.main()
