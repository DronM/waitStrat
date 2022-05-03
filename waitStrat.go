package waitStrat


type WaitStrategyValues struct {
	MaxTries int   //max try count
	Wait int 	//wait value
}
type WaitStrategy struct {
	CurTry int	//current try  count
	CurInd int	//current Strategies index
	Strategies []WaitStrategyValues
}
func (s *WaitStrategy) Init() {
	s.CurInd = 0
	s.CurTry = 0
}
func (s *WaitStrategy) NextWait() int {
	if s.Strategies == nil || s.CurInd >= len(s.Strategies) {
		return 0
	}
	s.CurTry++
	if s.Strategies[s.CurInd].MaxTries > 0 && s.CurTry > s.Strategies[s.CurInd].MaxTries {
		if  s.CurInd+1 < len(s.Strategies) {
			s.CurInd++
			s.CurTry = 0
		}else{
			//last
			s.CurTry = 0			
		}		
	}
	return s.Strategies[s.CurInd].Wait
}

