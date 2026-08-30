-- Sum of pair equal to target.
--
-- Given a sorted array (asc) and a target, find if there exists any pair of elements (arr[i],
-- arr[j]) such that their sum is equal to the target.
--
-- Input: arr = [10, 20, 35, 50]; target = 70
-- Output: true ([20, 50] add up to 70)

local function two_sum_1(arr, target)
	local left, right = 1, #arr

	while left < right do
		local sum = arr[left] + arr[right]

		if sum == target then
			return true
		end

		if sum < target then
			left = left + 1
		else
			right = right - 1
		end
	end

	return false
end

local function two_sum_2(arr, target)
	local complements = {}

	for _, n in pairs(arr) do
		if complements[n] ~= nil then
			return true
		end

		complements[target - n] = true
	end

	return false
end

local function test(name, fun)
	local ok, err = pcall(fun)

	if ok then
		print("PASS: " .. name)
	else
		print("FAIL: " .. name)
		print("      " .. err)
		os.exit(1)
	end
end

test("two_sum_1", function()
	local test_cases = {
		{
			arr = {10, 20, 35, 50},
			target = 70,
			expected = true,
		},
		{
			arr = {10, 20, 30},
			target = 70,
			expected = false,
		},
		{
			arr = {-8, 1, 4, 6, 10, 45},
			target = 16,
			expected = true,
		},
	}

	for _, tc in pairs(test_cases) do
		assert(two_sum_1(tc.arr, tc.target) == tc.expected, "Oops!")
	end
end)

test("two_sum_2", function()
	local test_cases = {
		{
			arr = {10, 20, 35, 50},
			target = 70,
			expected = true,
		},
		{
			arr = {10, 20, 30},
			target = 70,
			expected = false,
		},
		{
			arr = {-8, 1, 4, 6, 10, 45},
			target = 16,
			expected = true,
		},
	}

	for _, tc in pairs(test_cases) do
		assert(two_sum_2(tc.arr, tc.target) == tc.expected, "Oops!")
	end
end)
