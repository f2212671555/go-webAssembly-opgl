package main

import (
	"math/rand"
	"syscall/js"
	"time"
)

const (
	width  = 400
	height = 400
)

func sum(this js.Value, args []js.Value) interface{} {
	var sum int
	for _, val := range args {
		sum += val.Int()
	}
	return sum
}
func registerCallbacks() {
	js.Global().Set("sum", js.FuncOf(sum))
}
func getRandomNum() float32 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	n := float32(rand.Intn(10000))
	return n / 10000.0
}
func draw() {
	var canvas js.Value = js.
		Global().
		Get("document").
		Call("getElementById", "canvas")

	var context js.Value = canvas.Call("getContext", "2d")

	// reset
	canvas.Set("height", height)
	canvas.Set("width", width)
	context.Call("clearRect", 0, 0, width, height)

	// 随机绘制 50 条直线
	for i := 0; i < 50; i++ {
		context.Call("beginPath")
		context.Call("moveTo", getRandomNum()*width, getRandomNum()*height)
		context.Call("lineTo", getRandomNum()*width, getRandomNum()*height)
		context.Call("stroke")
	}
}

func main() {
	draw()
	c := make(chan struct{}, 0)
	println("Hello, WebAssembly!")
	registerCallbacks()
	<-c

}
