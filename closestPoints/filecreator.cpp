#include <algorithm>
#include <array>
#include <climits>
#include <cmath>
#include <fstream>
#include <iostream>
#include <random>
#include <set>
#include <string>
#include <utility>
#include <vector>

int main(int argc, char **argv) {

  if (argc != 3) {
    std::cerr << "ERROR: usage [fpath] [count]" << std::endl;
    return EXIT_FAILURE;
  }
  int count = atoi(*(argv + 2));
  int rc;
  std::ofstream ofs(*(argv + 1),
                    std::ios::out | std::ios::binary | std::ios::trunc);

  std::vector<std::array<int, 2>> pts;
  std::set<std::array<int, 2>> seen;
  for (int i = 0; i < count; ++i) {
    std::array<int, 2> pt;
    do {
      pt = {rand() % INT_MAX, rand() % INT_MAX};

    } while (seen.count(pt));
    seen.insert(pt);
    pts.push_back(pt);
  }
  if (!ofs) {
    std::cerr << "File failed to open" << std::endl;
    return EXIT_FAILURE;
  }

  for (const auto &a : pts) {
    ofs.write(reinterpret_cast<const char *>(&a[0]), sizeof(int) * 2);
  }
  ofs.close();

  return EXIT_SUCCESS;
}