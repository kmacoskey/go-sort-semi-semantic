package main

import (
	"bufio"
	"fmt"
	"github.com/cppforlife/go-semi-semantic/version"
	"io/ioutil"
	"os"
	"strings"
)

func ReadInput() []string {

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) == 2 {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return lines
}

func main() {

	stdin, _ := ioutil.ReadAll(os.Stdin)
	version_strings := strings.Split(string(stdin), "\n")

	versions := []version.Version{}
	for _, v := range version_strings {
		if len(v) > 0 {
			versions = append(versions, version.MustNewVersionFromString(v))
		}
	}

	for _, v := range versions {
		fmt.Println(v)
	}

}
