#include <print>
#include <vector>

auto two_sum_1(std::vector<int>, int) -> bool;
auto two_sum_2(std::vector<int>, int) -> bool;


auto main() -> int {
  auto arr = std::vector<int>{10, 20, 30, 50};
  auto target = 80;

  auto result = two_sum_1(arr, target);

  if (result) {
    std::println("Pair found");
  } else {
    std::println("Not a sausage.");
  }

  return 0;
}
