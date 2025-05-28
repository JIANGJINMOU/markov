package main

import (
	"encoding/csv"
	"os"
)

func LoadMarkovChainFromCSV(filename string) (*MarkovChain, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// 获取所有唯一状态
	stateMap := make(map[string]bool)
	for _, record := range records {
		stateMap[record[0]] = true
	}

	states := make([]string, 0, len(stateMap))
	for state := range stateMap {
		states = append(states, state)
	}

	// 初始化转移计数矩阵
	transitionCounts := make([][]int, len(states))
	for i := range transitionCounts {
		transitionCounts[i] = make([]int, len(states))
	}

	// 统计状态转移次数
	for i := 0; i < len(records)-1; i++ {
		currentState := records[i][0]
		nextState := records[i+1][0]

		currentIndex := -1
		nextIndex := -1
		for j, state := range states {
			if state == currentState {
				currentIndex = j
			}
			if state == nextState {
				nextIndex = j
			}
		}

		if currentIndex != -1 && nextIndex != -1 {
			transitionCounts[currentIndex][nextIndex]++
		}
	}

	// 计算转移概率矩阵
	transitionMatrix := make([][]float64, len(states))
	for i := range transitionMatrix {
		transitionMatrix[i] = make([]float64, len(states))

		total := 0
		for _, count := range transitionCounts[i] {
			total += count
		}

		if total > 0 {
			for j := range transitionMatrix[i] {
				transitionMatrix[i][j] = float64(transitionCounts[i][j]) / float64(total)
			}
		} else {
			// 如果没有数据，使用均匀分布
			for j := range transitionMatrix[i] {
				transitionMatrix[i][j] = 1.0 / float64(len(states))
			}
		}
	}

	return &MarkovChain{
		transitionMatrix: transitionMatrix,
		states:          states,
	}, nil
}