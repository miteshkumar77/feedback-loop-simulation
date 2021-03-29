#ifndef FILEREADER_HPP
#define FILEREADER_HPP

#include <algorithm>
#include <array>
#include <climits>
#include <cmath>
#include <iostream>
#include <mpi/mpi.h>
#include <string>
#include <utility>
#include <vector>

struct Point {

  Point() : x(0), y(0) {}

  Point(int x, int y) : x(x), y(y) {}

  bool operator==(const Point &p) { return x == p.x && y == p.y; }

  int x;
  int y;
};

namespace FReaderUtil {
char errorString[MPI_MAX_ERROR_STRING];
void readPx(char *fname, int me, int numranks, std::vector<Point> &ret);
} // namespace FReaderUtil
#endif // <<< FILEREADER_HPP