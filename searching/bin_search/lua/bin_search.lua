--[[
--
--  Binary search.
--
--]]

local function binary_search(arr, n)
  return 0, nil
end

local function test(name, fun)
	local ok, err = pcall(fun)

	if ok then
		print('PASS: ' .. name)
	else
		print('FAIL: ' .. name)
		print('....  ' .. err)
		os.exit(1)
	end
end

test('Test binary search', function()
  local test_cases = {
    {
      name = 'Element is present.',
      arr = {1, 2, 3, 4, 5},
      n = 4,
      want = {
        index = 4,
        error = nil,
      },
    },
    {
      name = 'Element isn\'t present.',
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
    index, error = binary_search(tc.arr, n)
    assert(error == tc.want.error)
    assert(index == tc.want.index)
  end
end)
