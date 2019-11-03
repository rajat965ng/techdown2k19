package main

import (
	"fmt"
	"math"
)

var (
	buff = make(chan Box)
)

type (
	Product struct {
		name string
		l    int
		w    int
		h    int
	}

	Box struct {
		l int
		w int
		h int
	}
)

func getLeastBestLength(availableBox []Box, product Product, lengthBuff chan Box) {
	diff, len, final := math.MaxInt64, product.l, Box{}
	for _, v := range availableBox {
		if v.l-len > 0 && v.l-len < diff {
			diff = v.l - len
			final = v
		}
	}
	lengthBuff <- final
}

func getLeastBestWidth(availableBox []Box, product Product, widthBuff chan Box) {
	diff, len, final := math.MaxInt64, product.w, Box{}
	for _, v := range availableBox {
		if v.w-len > 0 && v.w-len < diff {
			diff = v.w - len
			final = v
		}
	}
	widthBuff <- final
}

func getLeastBestHieght(availableBox []Box, product Product, hieghtBuff chan Box) {
	diff, len, final := math.MaxInt64, product.h, Box{}
	for _, v := range availableBox {
		if v.h-len > 0 && v.h-len < diff {
			diff = v.h - len
			final = v
		}
	}
	hieghtBuff <- final
}

func getLargestBoxOfLeastSize(buff chan Box) Box {
	temp, final := math.MinInt64, Box{}

	for i := 0; i < 3; i++ {
		v := <-buff
		fmt.Println("Boxes: ", v)
		if v.l+v.w+v.h > temp {
			temp = v.l + v.w + v.h
			final = v
		}
	}
	fmt.Println("final: ", final, " temp:", temp)
	return final
}

func getBestBox(availableBox []Box, product Product) Box {

	go getLeastBestLength(availableBox, product, buff)
	go getLeastBestWidth(availableBox, product, buff)
	go getLeastBestHieght(availableBox, product, buff)

	final := getLargestBoxOfLeastSize(buff)

	return final
}

func main() {
	availableBox := []Box{
		Box{l: 10, w: 10, h: 20},
		Box{l: 10, w: 20, h: 20},
		Box{l: 15, w: 20, h: 25},
		Box{l: 15, w: 30, h: 50},
		Box{l: 30, w: 30, h: 60},
		Box{l: 40, w: 40, h: 40},
		Box{l: 50, w: 40, h: 45},
		Box{l: 60, w: 60, h: 50},
	}

	products := Product{name: "Whiskey", l: 7, w: 30, h: 37}

	finalBox := getBestBox(availableBox, products)
	fmt.Println("~~~~ The result is ~~~~")
	fmt.Println(finalBox)
}
