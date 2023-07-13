package types

type LsvFeeders []LsvFeeder

func NewLsvFeeder(lsvAddr string, feederAddr string, updateHeight int64) *LsvFeeder {
	return &LsvFeeder{
		LsvAddr:      lsvAddr,
		FeederAddr:   feederAddr,
		UpdateHeight: updateHeight,
	}
}
