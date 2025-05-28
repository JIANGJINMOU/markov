package main

import (
	"fmt"
	"os"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/golang/freetype"
)

func calculateRiskMetrics(states []string) {
	var upCount, downCount, stableCount int
	for _, state := range states {
		switch state {
		case "上涨":
			upCount++
		case "下跌":
			downCount++
		case "持平":
			stableCount++
		}
	}

	total := len(states)
	fmt.Printf("风险评估报告:\n")
	fmt.Printf("上涨概率: %.2f%%\n", float64(upCount)/float64(total)*100)
	fmt.Printf("下跌概率: %.2f%%\n", float64(downCount)/float64(total)*100)
	fmt.Printf("持平概率: %.2f%%\n", float64(stableCount)/float64(total)*100)
	fmt.Printf("风险指数: %.2f\n", float64(downCount)/float64(upCount+1))

	// 可视化结果
	visualizeResults(states, upCount, downCount, stableCount)
}

func visualizeResults(states []string, upCount, downCount, stableCount int) {
	// 创建状态变化图表
	xValues := make([]float64, len(states))
	yValues := make([]float64, len(states))
	for i := range states {
		xValues[i] = float64(i)
		switch states[i] {
		case "上涨":
			yValues[i] = 1
		case "持平":
			yValues[i] = 0
		case "下跌":
			yValues[i] = -1
		}
	}

	// 加载中文字体
	fontBytes, err := os.ReadFile("C:\\Windows\\Fonts\\msyh.ttf")
	if err != nil {
		fontBytes, err = os.ReadFile("C:\\Windows\\Fonts\\simsun.ttc")
		if err != nil {
			fmt.Println("警告: 无法加载中文字体，将使用默认字体")
			return
		}
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println("警告: 无法解析字体，将使用默认字体")
		return
	}

	graph := chart.Chart{
		Title: "金融市场状态模拟",
		TitleStyle: chart.Style{
			Font: font,
		},
		XAxis: chart.XAxis{
			Name: "天数",
			NameStyle: chart.Style{
				Font: font,
			},
			Style: chart.Style{
				Font: font,
			},
		},
		YAxis: chart.YAxis{
			Name: "状态",
			NameStyle: chart.Style{
				Font: font,
			},
			Style: chart.Style{
				Font: font,
			},
			Ticks: []chart.Tick{
				{Value: -1, Label: "下跌"},
				{Value: 0, Label: "持平"},
				{Value: 1, Label: "上涨"},
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xValues,
				YValues: yValues,
			},
		},
	}

	f, _ := os.Create("state_changes.png")
	defer f.Close()
	graph.Render(chart.PNG, f)

	// 创建风险分布饼图
	pie := chart.PieChart{
		Title: "风险分布",
		TitleStyle: chart.Style{
			Font: font,
		},
		Values: []chart.Value{
			{Value: float64(upCount), Label: "上涨"},
			{Value: float64(downCount), Label: "下跌"},
			{Value: float64(stableCount), Label: "持平"},
		},
		Font: font,
	}

	f2, _ := os.Create("risk_distribution.png")
	defer f2.Close()
	pie.Render(chart.PNG, f2)

	fmt.Println("已生成可视化图表: state_changes.png 和 risk_distribution.png")
}