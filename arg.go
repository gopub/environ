package environ

import (
	"github.com/gopub/conv"
	"github.com/gopub/types"
	"log"
	"os"
)

func IntArg(index int) (int, error) {
	if index >= len(os.Args) {
		return 0, types.ErrNotExist
	}
	return conv.ToInt(os.Args[index])
}

func MustIntArg(index int) int {
	n, err := IntArg(index)
	if err != nil {
		log.Panicf("No int64 arg at %d: %v", index, err)
	}
	return n
}

func Int64Arg(index int) (int64, error) {
	if index >= len(os.Args) {
		return 0, types.ErrNotExist
	}
	return conv.ToInt64(os.Args[index])
}

func MustInt64Arg(index int) int64 {
	n, err := Int64Arg(index)
	if err != nil {
		log.Panicf("No int64 arg at %d: %v", index, err)
	}
	return n
}

func StringArg(index int) (string, error) {
	if index >= len(os.Args) {
		return "", types.ErrNotExist
	}
	return os.Args[index], nil
}

func MustStringArg(index int) string {
	s, err := StringArg(index)
	if err != nil {
		log.Panicf("No string arg at %d: %v", index, err)
	}
	return s
}
