package fleet

import (
	"fmt"
	"math"
)

const (
	minNumDistricts          = 1
	maxNumDistricts          = 100
	minNumScootersInDistrict = 0
	maxNumScootersInDistrict = 1000
	minNumScootersForFM      = 1
	maxNumScootersForFM      = 999
	minNumScootersForFE      = 1
	maxNumScootersForFE      = 1000
)

var (
	ErrNumDistricts          = fmt.Errorf("must have between %d and %d districts", minNumDistricts, maxNumDistricts)
	ErrNumScootersForFM      = fmt.Errorf("FM must be able to maintain between %d and %d scooters", minNumScootersForFM, maxNumScootersForFM)
	ErrNumScootersForFE      = fmt.Errorf("FE must be able to maintain between %d and %d scooters", minNumScootersForFE, maxNumScootersForFE)
	ErrNumScootersInDistrict = fmt.Errorf("each district must have between %d and %d scooters", minNumScootersInDistrict, maxNumScootersInDistrict)
)

// MinNumFE returns the minimun number of fleet engineers (FE) required to help
// the fleet manager (FM) so that every scooter in every district Coup operates
// in is maintained.
//
// []int scooters has as many elements as there are districts. For each i,
// scooters[i] denotes the number of scooters in that district. scooters must
// contain between 1 and 100 elements. Each element must be between 0 and 1000.
//
// int c denotes the maximum number of scooters the FM is able to maintain
//
// int p denotes the maximum number of scooters an FE is able to maintain
//
// Returns an error for invalid inputs.
func MinNumFE(scooters []int, c, p int) (int, error) {
	if err := assertInputs(scooters, c, p); err != nil {
		return 0, err
	}

	total := 0
	maxSavings := 0

	for _, numScooters := range scooters {
		if err := assertNumScootersInDistrict(numScooters); err != nil {
			return 0, err
		}

		numFE := int(math.Ceil(float64(numScooters) / float64(p)))
		total += numFE

		numFEWithFM := int(math.Ceil(math.Max(float64(numScooters-c)/float64(p), 0)))
		if potentialSavings := numFE - numFEWithFM; potentialSavings > maxSavings {
			maxSavings = potentialSavings
		}
	}

	return total - maxSavings, nil
}

// assertInputs verifies the number of districts as well as the number of
// scooters the FM and any FE are able to maintain.
func assertInputs(scooters []int, c, p int) error {
	if n := len(scooters); n < minNumDistricts || n > maxNumDistricts {
		return ErrNumDistricts
	}

	if c < minNumScootersForFM || c > maxNumScootersForFM {
		return ErrNumScootersForFM
	}

	if p < minNumScootersForFE || p > maxNumScootersForFE {
		return ErrNumScootersForFE
	}

	return nil
}

// assertNumScootersInDistrict verifies the number of scooters in a given district
func assertNumScootersInDistrict(numScooters int) error {
	if numScooters < minNumScootersInDistrict || numScooters > maxNumScootersInDistrict {
		return ErrNumScootersInDistrict
	}

	return nil
}
