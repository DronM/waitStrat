**WaitStrategy** manages reconnection strategies. For example data base reconnections, network protocol reconnections etc.
On construction several strategies are defined. Every strategy includes maximum try count and wait interval.
On failur wait interval can be fetched with *NextWait()* function. For example *time.Sleep(time.Duration(WaitStrategys.NextWait()) * time.Millisecond)*
Zero try count means endlees try.

import github.com/DronM/
conStart =: waitStrat.WaitStrategy{
	Strategies: []waitStrat.WaitStrategyValues{
		waitStrat.WaitStrategyValues{10,1000},
		waitStrat.WaitStrategyValues{12,10000},
		waitStrat.WaitStrategyValues{0,30000},
	}},

