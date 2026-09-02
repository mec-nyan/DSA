--[[
--
--  Binary search.
--
--]]

local function binary_search(arr, n)
  local left, right = 1, #arr

  while left <= right do
    local middle = math.floor((right - left) / 2) + left

    local current = arr[middle]

    if current == n then
      return middle, nil
    end

    if current < n then
      left = middle + 1
    else
      right = middle - 1
    end
  end

  return 0, 'Not found.'

end

local function test(name, fun)
	local ok, err = pcall(fun)

	if ok then
		print('PASS: ' .. name .. '\t\x1b[32mOk ✔\x1b[0m')
	else
		print('FAIL: ' .. name .. '\t\x1b[31m:( ✖\x1b[0m')
		print('....  ' .. err)
		os.exit(1)
	end
end

test('Test binary search', function()
  local test_cases = {
    {
      name = 'Should be found.',
      arr = {1, 2, 3, 4, 5},
      n = 4,
      want = {
        index = 4,
        error = nil,
      },
    },
    {
      name = "There's nothing here ...",
      arr = {1, 2, 3, 4, 5},
      n = 7,
      want = {
        index = 0,
        error = 'Not found.',
      },
    },
  }

  for _, tc in pairs(test_cases) do
    print('==', tc.name)
    index, error = binary_search(tc.arr, tc.n)
    assert(error == tc.want.error, "Errors don't match!")
    assert(index == tc.want.index, "Indices don't match!")
  end
end)
