// bin_search.cpp
#include <cstdint>
#include <optional>
#include <print>
#include <vector>

// Search for `n` in a sorted list of numbers.  If present, return its offset.
auto bin_search(std::vector<int32_t>, int32_t) -> std::optional<int32_t>;

auto main() -> int
{

    auto elems = std::vector<int32_t>{10, 20, 32, 44};

    auto n = 32;

    auto is_in = bin_search(elems, n);

    if (is_in.has_value())
    {
        std::println("{} found at offset {}", n, is_in.value());
    }
    else
    {
        std::println("{} was not found ...", n);
    }

    return 0;
}
