package magicloop

import "os"

func FileExists(fname string) bool {
	_, err := os.Stat(fname)
	return !os.IsNotExist(err)
}
