package main

import "math/rand"

type MarkovChain struct {
	transitionMatrix [][]float64
	states          []string
}

func NewMarkovChain(states []string, transitionMatrix [][]float64) *MarkovChain {
	return &MarkovChain{
		transitionMatrix: transitionMatrix,
		states:          states,
	}
}

func (mc *MarkovChain) NextState(currentState string) string {
	currentIndex := -1
	for i, state := range mc.states {
		if state == currentState {
			currentIndex = i
			break
		}
	}
	if currentIndex == -1 {
		return ""
	}

	r := rand.Float64()
	cumulativeProb := 0.0
	for nextIndex, prob := range mc.transitionMatrix[currentIndex] {
		cumulativeProb += prob
		if r <= cumulativeProb {
			return mc.states[nextIndex]
		}
	}
	return ""
}