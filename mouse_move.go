package main

import (
	"fmt"
	"math/rand"
	"syscall"
	"time"
	"unsafe"
)

var (
	user32              = syscall.NewLazyDLL("user32.dll")
	getCursorPos        = user32.NewProc("GetCursorPos")
	setCursorPos        = user32.NewProc("SetCursorPos")
)

type POINT struct {
	X int32
	Y int32
}

func getMousePos() (int32, int32, error) {
	var pt POINT
	ret, _, err := getCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return 0, 0, err
	}
	return pt.X, pt.Y, nil
}

func setMousePos(x, y int32) error {
	ret, _, err := setCursorPos.Call(uintptr(x), uintptr(y))
	if ret == 0 {
		return err
	}
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("鼠标自动移动工具已启动...")

	for {
		x, y, err := getMousePos()
		if err != nil {
			//fmt.Println("获取鼠标位置失败:", err)
			continue
		}

		// 偏移范围 -50 到 +50
		offsetX := int32(rand.Intn(101) - 50)
		offsetY := int32(rand.Intn(101) - 50)

		newX := x + offsetX
		newY := y + offsetY

		err = setMousePos(newX, newY)
		if err != nil {
			//fmt.Println("移动鼠标失败:", err)
		} else {
			//fmt.Printf("鼠标已移动到：(%d, %d)\n", newX, newY)
		}

		sleepSeconds := 10 + rand.Intn(21) // 10~30秒
		//fmt.Printf("等待 %d 秒后再次移动...\n", sleepSeconds)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}
}
