package mutator

import "github.com/2er0/GoGp/iface"

// A SigKeep does not change the sigma over the time
type SigKeep struct {
}

func NewSigmaOri() iface.MutChanger {
	return &SigKeep{}
}

func (s *SigKeep) SetLambda(lam int) {}

func (s *SigKeep) Update(suc int, sig float32) float32 {
	return sig
}

// A SigNotKeep does change the sigma over time by the 1/5 rule
type SigNotKeep struct {
	limit int
}

func NewSigma15() iface.MutChanger {
	return &SigNotKeep{limit: 0}
}

func (s *SigNotKeep) SetLambda(lam int) {
	s.limit = int(float32(lam) * 0.2)
}

// Update changes the given sigma by the 1/5 rule
func (s *SigNotKeep) Update(suc int, sig float32) float32 {
	if suc < s.limit {
		return sig * 0.82
	} else if suc > s.limit {
		return sig * (1 / 0.82)
	}
	return sig
}
