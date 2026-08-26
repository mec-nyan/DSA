"""
Sum of pair equal to target.

Given a sorted array (asc) and a target, find if there exists any pair of elements (arr[i], arr[j])
such that their sum is equal to the target.

Input: arr = [10, 20, 35, 50]; target = 70
Output: true ([20, 50] add up to 70)

"""


def two_sum(arr: list[int], target: int) -> bool:
    left, right = 0, len(arr) - 1

    while left < right:
        sum = arr[left] + arr[right]

        if sum == target:
            return True
        elif sum < target:
            left += 1
        else:
            right -= 1

    return False
