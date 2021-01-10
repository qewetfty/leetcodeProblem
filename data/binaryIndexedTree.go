package data

type BIT struct {
	numbers []int
	length  int
}

func NewBIT(length int) *BIT {
	number := make([]int, length)
	return &BIT{numbers: number, length: length}
}

func (b *BIT) Lowbit(i int) int {
	return i & (-i)
}

func (b *BIT) Update(index int, add int) {
	for i := index; i <= b.length; i += b.Lowbit(i) {
		b.numbers[i] += add
	}
}

func (b *BIT) GetSum(index int) int {
	result := 0
	for i := index; i > 0; i -= b.Lowbit(i) {
		result += b.numbers[i]
	}
	return result
}
