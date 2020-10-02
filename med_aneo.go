package main

import (
	"fmt"
	"math"
	"os"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func convertSpeedKH2MS(speed int) float64 {
	speef := float64(speed)
	return float64(speef*1000/60/60)

}

func convertSpeedMS2KH(speed float64) int {
	speef := float64(speed)
	return int(speef/1000*60*60)

}

/*
 *         -e    0   +e
 *  --------[----|----]----
 *  --------ok=============
 */
func trigoSolveError(amplitude float64) bool {
	// error := 1e-14
	error := 9.554722e-15
	// error := 0.0

	fmt.Fprintf(os.Stderr, "Amp is: %e\n", amplitude)
	if math.Abs(amplitude) < error {
		amplitude = 0
		fmt.Fprintf(os.Stderr, "  `->0\n")
	}

	if amplitude >= 0 {
		return true
	}
	return false

}

func trigoSolve(speed int, distance int, duration int) bool {
	// di := float64(distance)
	// ti := float64(duration)
	// time_spent := di / speed
	amplitude := math.Sin(math.Pi*float64(distance)/float64(speed)/float64(duration)/1000.0*3600.0)
	if trigoSolveError(amplitude) {
		return true
	}
	fmt.Fprintf(os.Stderr, "Failed for speed %d:\n->ti:%d\n->amp:%e\n", speed, duration, amplitude)
	return false
}

func solve(speed int, lightCount int, distances []int, durations []int) bool {
	for i := 0; i < lightCount; i++ {
		passed := trigoSolve(speed, distances[i], durations[i])
		if !passed {
			return false
		}
	}
	return true
}

func main() {
    var khMaxSpeed int
    fmt.Scan(&khMaxSpeed)


    var lightCount int
    fmt.Scan(&lightCount)

	distances := make([]int, lightCount)
	durations := make([]int, lightCount)

	total_distance := 0

    for i := 0; i < lightCount; i++ {
        var distance, duration int
        fmt.Scan(&distance, &duration)
		total_distance += distance
		distances[i] = distance
		durations[i] = duration
    }
	//Solve
	for i := khMaxSpeed; i > 0 ; i-- {
		fmt.Fprintf(os.Stderr, "Testing %d kh/h\n", i)
		// current_speed := convertSpeedKH2MS(i)
		all_passed := solve(i, lightCount, distances, durations)
		if all_passed {
			fmt.Println(i)
			fmt.Fprintf(os.Stderr, "Found for speed %d\n", i)
			break
		}

	}
	// best_time := total_distance / maxSpeed
    // fmt.Fprintln(os.Stderr, "Debug messages...")
}
