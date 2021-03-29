#include <algorithm>
#include <array>
#include <climits>
#include <cmath>
#include <fstream>
#include <iostream>
#include <string>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>
#include <utility>
#include <vector>
using namespace std;

struct Point;
double dist(const Point &p1, const Point &p2);
array<Point, 2> closestPair(vector<Point> points);
array<Point, 2> closestPairDivC(const vector<Point> &Px, const vector<int> &Py,
                                int l, int r);
pair<vector<int>, vector<int>> divide(const vector<int> &Py, int mid);
vector<int> getStrip(const vector<Point> &Px, const vector<int> &Py,
                     double center, double d);

struct Point {

  Point() : x(0), y(0) {}

  Point(int x, int y) : x(x), y(y) {}

  bool operator==(const Point &p) { return x == p.x && y == p.y; }

  int x;
  int y;
};

array<Point, 2> closestPairBruteForce(const vector<Point> &points, int l,
                                      int r) {
  array<Point, 2> ans{Point(INT_MIN, INT_MIN), Point(INT_MAX, INT_MAX)};
  for (int i = l; i <= r; ++i) {
    for (int j = i + 1; j <= r; ++j) {
      if (dist(ans[0], ans[1]) > dist(points[i], points[j])) {
        ans = {points[i], points[j]};
      }
    }
  }

  return ans;
}

array<Point, 2> closestPairDivC(const vector<Point> &Px, const vector<int> &Py,
                                int l, int r) {
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
  array<Point, 2> res;

  if (dl < dr) {
    res = resultLeft;
    d = dl;
  } else {
    res = resultRight;
    d = dr;
  }

  double center = ((double)Px[mid].x + Px[mid + 1].x) / 2;

  vector<int> strip = getStrip(Px, Py, center, d);

  for (int i = 0; i < strip.size(); ++i) {
    for (int j = i + 1; j < min((int)strip.size(), i + 16); ++j) {
      double di = dist(Px[strip[i]], Px[strip[j]]);
      if (di < d) {
        d = di;
        res = {Px[strip[i]], Px[strip[j]]};
      }
    }
  }

  return res;
}

pair<vector<int>, vector<int>> divide(const vector<int> &Py, int mid) {
  auto ret = make_pair(vector<int>(), vector<int>());
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

vector<int> getStrip(const vector<Point> &Px, const vector<int> &Py,
                     double center, double d) {
  vector<int> ans;
  double lb = center - d;
  double ub = center + d;
  for (int i : Py) {
    if (Px[i].x >= lb && Px[i].x <= ub) {
      ans.push_back(i);
    }
  }
  return ans;
}

array<Point, 2> closestPair(vector<Point> points) {
  sort(points.begin(), points.end(),
       [](const Point &a, const Point &b) -> bool { return a.x < b.x; });

  vector<int> Py(points.size());

  for (int i = 0; i < points.size(); ++i) {
    Py[i] = i;
  }

  sort(Py.begin(), Py.end(),
       [&](int a, int b) -> bool { return points[a].y < points[b].y; });

  return closestPairDivC(points, Py, 0, points.size() - 1);
}

int main(int argc, char **argv) {

  if (argc != 2) {
    std::cerr << "ERROR: invalid file" << std::endl;
    return EXIT_FAILURE;
  }
  struct stat st;
  if (-1 == lstat(*(argv + 1), &st)) {
    std::cerr << "ERROR: invalid faile" << std::endl;
    return EXIT_FAILURE;
  }

  size_t sz = st.st_size / sizeof(Point);
  std::ifstream ifs(*(argv + 1), ios::in | ios::binary);
  if (!ifs.good()) {
    std::cerr << "ERROR: invalid file" << std::endl;
    return EXIT_FAILURE;
  }

  vector<Point> points(sz);
  ifs.read((char *)&points[0], sz * sizeof(Point));
  std::cout << points.size() << std::endl;
  // for (const auto &p : points) {
  //   std::cout << p.x << ' ' << p.y << std::endl;
  // }
  auto ans = closestPair(points);
  // auto expectedans = closestPairBruteForce(points, 0, points.size() - 1);
  // cout << "ANSWER:   (" << ans[0].x << "," << ans[0].y << "); (" << ans[1].x
  //      << "," << ans[1].y << ")" << endl;
  // cout << "EXPECTED: (" << expectedans[0].x << "," << expectedans[0].y << ");
  // ("
  //      << expectedans[1].x << "," << expectedans[1].y << ")" << endl;
  cout << "ANSWER: " << dist(ans[0], ans[1]) << endl;
  // cout << "EXPECTED: " << dist(expectedans[0], expectedans[1]) << endl;
  return 0;
}