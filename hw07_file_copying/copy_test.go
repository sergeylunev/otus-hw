package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	outFile    = "out.txt"
	outNewFile = "out_new.txt"
)

func TestCopy(t *testing.T) {
	defer removeCreatedFiles()

	t.Run("no input file", func(t *testing.T) {
		err := Copy("no/such/file", "and/no/such/file", 0, 0)

		require.ErrorIs(t, err, ErrNoFileToCopy)
	})

	t.Run("to big offset", func(t *testing.T) {
		err := Copy("testdata/small_file.txt", "out.txt", 100000000, 0)

		require.ErrorIs(t, err, ErrOffsetExceedsFileSize)
	})

	t.Run("error if use some infinity file", func(t *testing.T) {
		err := Copy("/dev/urandom", "out.txt", 0, 0)

		require.ErrorIs(t, err, ErrUnsupportedFile)
	})

	t.Run("creating output file", func(t *testing.T) {
		err := Copy("testdata/small_file.txt", outFile, 0, 0)
		require.Nil(t, err)
		_, err = os.Open(outFile)
		require.Nil(t, err)

		err = Copy("testdata/small_file.txt", outNewFile, 0, 0)
		require.Nil(t, err)
		_, err = os.Open(outNewFile)
		require.Nil(t, err)
	})

	t.Run("copy file as is", func(t *testing.T) {
		err := Copy("testdata/small_file.txt", "out.txt", 0, 0)

		require.Nil(t, err)

		checkFiles(t, outFile, "testdata/asis.txt")
	})

	t.Run("copy with offset", func(t *testing.T) {
		err := Copy("testdata/small_file.txt", "out.txt", 4, 0)
		require.Nil(t, err)
		checkFiles(t, outFile, "testdata/offset_4.txt")
	})

	t.Run("copy with offset and limit", func(t *testing.T) {
		err := Copy("testdata/small_file.txt", "out.txt", 4, 1)
		require.Nil(t, err)
		checkFiles(t, outFile, "testdata/offset_4_limit_1.txt")
	})

	t.Run("copy with offset and big limit", func(t *testing.T) {
		err := Copy("testdata/small_file.txt", outFile, 4, 1000)
		require.Nil(t, err)
		checkFiles(t, outFile, "testdata/offset_4.txt")
	})

}

func checkFiles(t *testing.T, actual string, expected string) {
	afi, err := os.Stat(actual)
	require.Nil(t, err)

	efi, err := os.Stat(expected)
	require.Nil(t, err)

	require.Equal(t, efi.Size(), afi.Size())
}

func removeCreatedFiles() {
	os.Remove(outFile)
	os.Remove(outNewFile)
}
