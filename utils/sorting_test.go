package utils

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	elems:=[]int{6,5,4,3,2,1}
	BubbleSort(elems)
	assert.Equal(t,6,len(elems),"The length of elems is 6")
	assert.Equal(t,1,elems[0],"The elems[0] is 1")
	assert.Equal(t,6,elems[5],"The elems[5] is 6")

}

//for Fewer elements this bubble sort is good.
func BenchmarkBubbleSort1000(b *testing.B) {
	elems:=[]int{6,5,4,3,2,1}
	for i:=0;i<b.N;i++{
		BubbleSort(elems)
	}
}

func CreateData(n int) []int {
	elem:=make([]int,n)
	for i:=n;i>0;i--{
		elem[n-i]=i
	}
	return elem
}

//for very large set of elements the native sort function is better.
func BenchmarkBubbleSortForNativeSort10000(b *testing.B) {
	elems := CreateData(10000)
	for i:=0;i<b.N;i++  {
		sort.Ints(elems)
	}
}