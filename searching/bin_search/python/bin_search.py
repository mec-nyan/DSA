"""
Binary search (simple implementation).

"""


def binary_search(arr: list[int], n: int) -> tuple[int, str]:
    left, right = 0, len(arr) - 1

    while left <= right:
        middle = (right - left) // 2 + left
        current = arr[middle]

        if current == n:
            return middle, ""

        if current < n:
            left = middle + 1
        else:
            right = middle - 1

    return -1, "Not found"
