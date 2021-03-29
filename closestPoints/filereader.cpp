#include "filereader.hpp"

char errorString[MPI_MAX_ERROR_STRING];
int errorLen;
void FReaderUtil::readPx(char *fname, int me, int numranks,
                         std::vector<Point> &ret) {
  int rc;
  MPI_File fh;
  rc =
      MPI_File_open(MPI_COMM_WORLD, fname, MPI_MODE_RDONLY, MPI_INFO_NULL, &fh);
  if (rc != 0) {
    MPI_Error_string(rc, errorString, &errorLen);
    fprintf(stderr,
            "ERROR RANK %d: MPI_File_open() failed with code (%d) --- %s\n", me,
            rc, errorString);
    exit(EXIT_FAILURE);
  }

  MPI_Offset fsize;
  rc = MPI_File_get_size(fh, &fsize);
  if (rc != 0) {
    MPI_Error_string(rc, errorString, &errorLen);
    fprintf(stderr,
            "ERROR RANK %d: MPI_File_get_size() failed with code (%d) --- %s\n",
            me, rc, errorString);
    exit(EXIT_FAILURE);
  }
  MPI_Offset num_points = fsize / sizeof(Point);
  MPI_Offset byte_delta = sizeof(Point) * (num_points / (MPI_Offset)numranks);
  MPI_Offset byte_ofs = byte_delta * me;
  MPI_Offset read_bytes = (me + 1 == numranks ? fsize - byte_ofs : byte_delta);

  char *data = (char *)malloc(read_bytes * sizeof(Point));
  if (data == NULL) {
    fprintf(stderr, "ERROR RANK %d: malloc() failed\n");
    exit(EXIT_FAILURE);
  }

  rc = MPI_File_read_at(fh, byte_ofs, data, read_bytes, MPI_CHAR, NULL);
  if (rc != 0) {
    MPI_Error_string(rc, errorString, &errorLen);
    fprintf(stderr,
            "ERROR RANK %d: MPI_File_read_at() failed with code (%d) --- %s\n",
            me, rc, errorString);
  }

  rc = MPI_File_close(&fh);
  if (rc != 0) {
    MPI_Error_string(rc, errorString, &errorLen);
    fprintf(stderr,
            "ERROR RANK %d: MPI_File_close() failed with code (%d) --- %s\n",
            me, rc, errorString);
  }
  data = NULL;

  ret.assign(data, data + read_bytes);
}