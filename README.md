# 马尔可夫链金融模拟器

基于马尔可夫链的金融市场状态模拟器，用于风险评估和分析。

## 功能特性

- 从CSV文件加载历史数据，自动计算状态转移概率矩阵
- 模拟金融市场状态变化（上涨、持平、下跌）
- 生成风险评估报告（上涨/下跌/持平概率、风险指数）
- 可视化模拟结果（状态变化图和风险分布饼图）

## 使用说明

1. 准备CSV数据文件 `financial_data.csv`，每行包含一个状态（上涨/持平/下跌）
2. 运行程序：
   ```bash
   go run .
   ```
3. 程序将输出模拟结果和风险评估报告，并生成可视化图表：
   - `state_changes.png` - 状态变化趋势图
   - `risk_distribution.png` - 风险分布饼图

## 项目结构

```
├── data_loader.go      # 数据加载和转移矩阵计算
├── markov_chain.go     # 马尔可夫链核心逻辑
├── visualization.go    # 结果可视化
├── main.go             # 主程序入口
├── go.mod              # 依赖管理
└── README.md           # 项目说明
```

## 依赖

- Go 1.24+
- github.com/wcharczuk/go-chart/v2
- github.com/golang/freetype

## 许可证

MIT