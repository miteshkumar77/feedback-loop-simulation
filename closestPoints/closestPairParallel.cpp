#include "filereader.hpp"
#include <algorithm>
#include <array>
#include <climits>
#include <cmath>
#include <iostream>
#include <mpi.h>
#include <string>
#include <utility>
#include <vector>
int myrank, numranks;

double dist(const Point &p1, const Point &p2);
std::array<Point, 2> closestPair(std::vector<Point> points);
std::array<Point, 2> closestPair(const std::vector<Point> &Px,
                                 const std::vector<Point> &Py);
std::array<Point, 2> closestPairDivC(const std::vector<Point> &Px,
                                     const std::vector<int> &Py, int l, int r);
std::pair<std::vector<int>, std::vector<int>> divide(const std::vector<int> &Py,
                                                     int mid);
std::vector<int> getStrip(const std::vector<Point> &Px,
                          const std::vector<int> &Py, double center, double d);

std::array<Point, 2> closestPairBruteForce(const std::vector<Point> &points,
                                           int l, int r) {
  std::array<Point, 2> ans{Point(INT_MIN, INT_MIN), Point(INT_MAX, INT_MAX)};
  for (int i = l; i <= r; ++i) {
    for (int j = i + 1; j <= r; ++j) {
      if (dist(ans[0], ans[1]) > dist(points[i], points[j])) {
        ans = {points[i], points[j]};
      }
    }
  }

  return ans;
}

std::array<Point, 2> closestPairDivC(const std::vector<Point> &Px,
                                     const std::vector<int> &Py, int l, int r) {
  if (l + 2 >= r) {
    return closestPairBruteForce(Px, l, r);
  }

  int mid = l + (r - l) / 2;
  auto splitPy = divide(Py, mid);

  auto resultLeft = closestPairDivC(Px, splitPy.first, l, mid);
  auto resultRight = closestPairDivC(Px, splitPy.second, mid + 1, r);

  double dl = dist(resultLeft[0], resultLeft[1]);
  double dr = dist(resultRight[0], resultRight[1]);

  double d;
  std::array<Point, 2> res;

  if (dl < dr) {
    res = resultLeft;
    d = dl;
  } else {
    res = resultRight;
    d = dr;
  }

  double center = ((double)Px[mid].x + Px[mid + 1].x) / 2;

  std::vector<int> strip = getStrip(Px, Py, center, d);

  for (int i = 0; i < strip.size(); ++i) {
    for (int j = i + 1; j < std::min((int)strip.size(), i + 16); ++j) {
      double di = dist(Px[strip[i]], Px[strip[j]]);
      if (di < d) {
        d = di;
        res = {Px[strip[i]], Px[strip[j]]};
      }
    }
  }

  return res;
}

std::pair<std::vector<int>, std::vector<int>> divide(const std::vector<int> &Py,
                                                     int mid) {
  auto ret = std::make_pair(std::vector<int>(), std::vector<int>());
  ret.first.reserve(Py.size() / 2);
  ret.second.reserve(Py.size() / 2);

  for (int i : Py) {
    if (i <= mid) {
      ret.first.push_back(i);
    } else {
      ret.second.push_back(i);
    }
  }

  return ret;
}

double dist(const Point &p1, const Point &p2) {
  double dx = (double)p1.x - (double)p2.x;
  double dy = (double)p1.y - (double)p2.y;
  return sqrt(dx * dx + dy * dy);
}

std::vector<int> getStrip(const std::vector<Point> &Px,
                          const std::vector<int> &Py, double center, double d) {
  std::vector<int> ans;
  double lb = center - d;
  double ub = center + d;
  for (int i : Py) {
    if (Px[i].x >= lb && Px[i].x <= ub) {
      ans.push_back(i);
    }
  }
  return ans;
}

std::array<Point, 2> closestPair(std::vector<Point> points) {
  sort(points.begin(), points.end(),
       [](const Point &a, const Point &b) -> bool { return a.x < b.x; });

  std::vector<int> Py(points.size());

  for (int i = 0; i < points.size(); ++i) {
    Py[i] = i;
  }

  std::sort(Py.begin(), Py.end(),
            [&](int a, int b) -> bool { return points[a].y < points[b].y; });

  return closestPairDivC(points, Py, 0, points.size() - 1);
}
std::array<Point, 2> closestPair(const std::vector<Point> &Px,
                                 const std::vector<int> &Py) {

  return closestPairDivC(Px, Py, 0, Px.size() - 1);
}

int main(int argc, char **argv) {
  // Initialize MPI
  MPI_Init(&argc, &argv);
  MPI_Comm_rank(MPI_COMM_WORLD, &myrank);
  MPI_Comm_size(MPI_COMM_WORLD, &numranks);
  if (myrank == 0 && argc != 2) {
    fprintf(stderr, "Usage: %s datapath\n", argv[0]);
    exit(EXIT_FAILURE);
  }

  std::vector<Point> Px;
  std::cout << "Reading file..." << std::endl;
  FReaderUtil::readPx(argv[1], myrank, numranks, Px);
  std::cout << "Finished reading file..." << std::endl;
  size_t N = Px.size();
  std::vector<int> Py(N);
  for (size_t i = 0; i < N; ++i) {
    Py[i] = i;
  }
  int tag = 123;

  double ownd;
  double recvd;
  std::cout << "Computing local closest pair..." << std::endl;
  std::sort(Py.begin(), Py.end(),
            [&](int a, int b) -> bool { return Px[a].y < Px[b].y; });
  std::array<Point, 2> res = closestPair(Px, Py);
  std::cout << "Computed local closest pair..." << std::endl;
  ownd = dist(res[0], res[1]);
  bool wasrecv = true;
  for (unsigned long long int i = 1; i <= numranks;) {
    i <<= 1;
    if (myrank % i == 0 && myrank + (i >> 1) < numranks) {
      MPI_Recv(&recvd, 1, MPI_DOUBLE, myrank + (i >> 1), tag, MPI_COMM_WORLD,
               MPI_STATUS_IGNORE);
      ownd = std::min(ownd, recvd);
    } else if (myrank >= (i >> 1) && ((myrank - (i >> 1)) % i) == 0) {
      MPI_Send(&ownd, 1, MPI_DOUBLE, myrank - (i >> 1), tag, MPI_COMM_WORLD);
    }
  }
  if (myrank == 0) {
    printf("RESULT: %lf\n", d);
  }
  MPI_Finalize();
  return EXIT_SUCCESS;
}
