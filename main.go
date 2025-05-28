package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 从CSV文件加载历史数据
	mc, err := LoadMarkovChainFromCSV("financial_data.csv")
	if err != nil {
		fmt.Println("无法加载历史数据，使用默认参数:", err)
		// 使用默认参数
		states := []string{"上涨", "持平", "下跌"}
		transitionMatrix := [][]float64{
			{0.55, 0.25, 0.20}, // 从上涨转移到其他状态的概率
			{0.30, 0.40, 0.30}, // 从持平转移到其他状态的概率
			{0.15, 0.25, 0.60}, // 从下跌转移到其他状态的概率
		}
		mc = NewMarkovChain(states, transitionMatrix)
	}

	// 模拟100天的市场状态变化
	currentState := "持平"
	simulatedStates := make([]string, 0, 100)
	for i := 0; i < 100; i++ {
		simulatedStates = append(simulatedStates, currentState)
		fmt.Printf("第%d天: %s\n", i+1, currentState)
		currentState = mc.NextState(currentState)
	}

	// 计算风险指标
	calculateRiskMetrics(simulatedStates)
}