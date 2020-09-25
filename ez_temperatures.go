// https://www.codingame.com/training/easy/temperatures
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    // n: the number of temperatures to analyse
    var n int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&n)


	if n == 0 {
		fmt.Fprintln(os.Stderr, "No data")
		fmt.Println("0")
		os.Exit(0)
	}

    scanner.Scan()
    inputs := strings.Split(scanner.Text()," ")
	var tmin int64
	tmin = 5526
    for i := 0; i < n; i++ {
        // t: a temperature expressed as an integer ranging from -273 to 5526
        t,_ := strconv.ParseInt(inputs[i],10,64)
        _ = t
		var tsq int64
		var tmsq int64
		tsq =  t*t
		tmsq = tmin*tmin
		if tsq < tmsq {
			tmin = t
		} else if tsq == tmsq  {
			if tmin < t {
				tmin = t
			}
		}
    }

    // fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Println(tmin)// Write answer to stdout
}

