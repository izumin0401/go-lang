package main

import (
	"log"
	"strconv"
	"time"
)

// スリープさせる
func sleep(i int) {
	log.Print(strconv.Itoa(i + 1) + "回目")

	time.Sleep(1 * time.Second)
}

// メイン関数
func main() {
	const sleepSecondTime int = 3

	log.Print("start")

	for i := 0; i < sleepSecondTime; i ++ {
		go sleep(i)  // ゴルーチン使用
	}

	time.Sleep(4 * time.Second)

	log.Print("end")
}
