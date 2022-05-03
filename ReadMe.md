**WaitStrategy** manages reconnection strategies. For example data base reconnections, network protocol reconnections etc.<br/>
Several strategies should be defined on construction. Every strategy includes maximum try count and wait interval.<br/>
On failur wait interval can be fetched with *NextWait()* function. For example *time.Sleep(time.Duration(WaitStrategys.NextWait()) * time.Millisecond)*<br/>
Zero try count means endlees try.<br/>
<br/>
import waitStrat<br/>
conStart =: waitStrat.WaitStrategy{<br/>
	Strategies: []waitStrat.WaitStrategyValues{<br/>
		waitStrat.WaitStrategyValues{10,1000},<br/>
		waitStrat.WaitStrategyValues{12,10000},<br/>
		waitStrat.WaitStrategyValues{0,30000},<br/>
	}},<br/>

