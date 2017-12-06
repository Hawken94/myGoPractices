package cardType

// CardType 牌类型
type CardType int32

const (
	Single        CardType = 1  // 单牌
	Pair          CardType = 2  // 对子，44、55
	ThreeSequence CardType = 3  // 3个牌的顺子，456、567
	FourSequence  CardType = 4  // 4个牌的顺子，4567
	FiveSequence  CardType = 5  // 5个牌的顺子，910JQK
	DoublePair    CardType = 6  // 两连对，4455
	ThreePair     CardType = 7  // 三连对，445566
	FourPair      CardType = 8  // 四连对，44556677
	FivePiar      CardType = 9  // 五连对，991010JJQQKK
	ThreeBomb     CardType = 10 // 3个相同牌的炸弹，444
	FourBomb      CardType = 11 // 4个相同牌的炸弹，4444
)

// CardsTypeMap 牌型map
var CardsTypeMap map[string]ValuePower

// ValuePower 牌值结构体
type ValuePower struct {
	cardType CardType
	value    int
}

func init() {
	// CardsTypeMap := make(map[string]ValuePower)
	CardsTypeMap = map[string]ValuePower{
		"4":  {Single, 4},
		"5":  {Single, 5},
		"6":  {Single, 6},
		"7":  {Single, 7},
		"9":  {Single, 9},
		"10": {Single, 10},
		"11": {Single, 11},
		"12": {Single, 12},
		"13": {Single, 13},

		"44":   {Pair, 4},
		"55":   {Pair, 5},
		"66":   {Pair, 6},
		"77":   {Pair, 7},
		"99":   {Pair, 9},
		"1010": {Pair, 10},
		"1111": {Pair, 11},
		"1212": {Pair, 12},
		"1313": {Pair, 13},

		"456":    {ThreeSequence, 4},
		"567":    {ThreeSequence, 5},
		"91011":  {ThreeSequence, 9},
		"101112": {ThreeSequence, 10},
		"111213": {ThreeSequence, 11},

		"4567":     {FourSequence, 4},
		"9101112":  {FourSequence, 9},
		"10111213": {FourSequence, 10},

		"910111213": {FiveSequence, 9},

		"4455":     {DoublePair, 4},
		"5566":     {DoublePair, 5},
		"6677":     {DoublePair, 6},
		"991010":   {DoublePair, 9},
		"10101111": {DoublePair, 10},
		"11111212": {DoublePair, 11},
		"12121313": {DoublePair, 12},

		"445566":       {ThreePair, 4},
		"556677":       {ThreePair, 5},
		"9910101111":   {ThreePair, 9},
		"101011111212": {ThreePair, 10},
		"111112121313": {ThreePair, 11},

		"44556677":         {FourPair, 4},
		"99101011111212":   {FourPair, 9},
		"1010111112121313": {FourPair, 10},

		"9910101111121213": {FivePiar, 9},

		"444":    {ThreeBomb, 4},
		"555":    {ThreeBomb, 5},
		"666":    {ThreeBomb, 6},
		"777":    {ThreeBomb, 7},
		"999":    {ThreeBomb, 9},
		"101010": {ThreeBomb, 10},
		"111111": {ThreeBomb, 11},
		"121212": {ThreeBomb, 12},
		"131313": {ThreeBomb, 13},

		"4444":     {FourBomb, 4},
		"5555":     {FourBomb, 5},
		"6666":     {FourBomb, 6},
		"7777":     {FourBomb, 7},
		"9999":     {FourBomb, 9},
		"10101010": {FourBomb, 10},
		"11111111": {FourBomb, 11},
		"12121212": {FourBomb, 12},
		"13131313": {FourBomb, 13},
	}
}
}
