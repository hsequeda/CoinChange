package main

type CoinType int

const (
	TWENTY_FIVE CoinType = iota
	TEN
	FIVE
	ONE
)

type Individual struct {
	list [4]byte
}

func (ind Individual) GetUtility() int {
	indUtility := 0
	for coinType, val := range ind.list {
		switch coinType {
		case int(TWENTY_FIVE):
			indUtility += int(val)
			break
		case int(TEN):
			indUtility += int(val)
			break
		case int(FIVE):
			indUtility += int(val)
			break
		case int(ONE):
			indUtility += int(val)
			break
		}
	}
	return indUtility
}
