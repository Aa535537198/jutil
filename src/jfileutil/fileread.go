package jfile

import (
	"bufio"
	"bytes"
	"github.com/Aa535537198/jutil/src/base/characters"
	"github.com/Aa535537198/jutil/src/jpreconditionutil"
	"io"
	"os"
	"strings"
)

type JFile interface {
}

// 一次性读取文件
func ReadFileByOnce(path string) (string, []byte) {
	bytes, err := os.ReadFile(path)

	jpreconditionutil.IsEqual(err == nil, err)
	resStr := string(bytes)

	return resStr, bytes
}

// 逐行读取文件   (包含start 和  end)
func ReadFileByLines(path string, start, end int) []string {
	if start < 1 {
		start = 1
	}
	file, err := os.Open(path)
	jpreconditionutil.IsEqual(err == nil, err)

	var strs []string
	strs = make([]string, 1)
	reader := bufio.NewReader(file)

	index := 1
	f := true
	st := start
	for {
		if start > end && end != -1 {
			break
		}
		s, err := reader.ReadString(characters.CN)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if f && index < st {
			index++
			continue
		}
		f = false
		strs = append(strs, strings.Trim(s, characters.RN))
		start++
	}

	return strs
}

// 根据文件路径读取文件长度
func ReadFileLinesByPath(path string) (int, error) {
	r, _ := os.Open(path)
	return ReadFileLinesByFile(r)
}

// 读取文件的行数
func ReadFileLinesByFile(r *os.File) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
