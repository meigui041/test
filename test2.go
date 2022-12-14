package main

import (
	"fmt"
	"time"
	"unsafe"
)

func ttt() {
	s := struct{}{}
	s1 := struct{}{}
	fmt.Println("空结构体内存使用情况", unsafe.Sizeof(s))
	fmt.Printf("空结构体指针使用情况:s =%p,s1=%p,两个指针比较结果%v", &s, &s1, &s == &s1)3
	strChan := make(chan string, 3)
	sigChan := make(chan struct{}, 2)  //接收数据信号
	sigChan1 := make(chan struct{}, 2) //操作完成信号

	go func() {
		//用来接收信息
		<-sigChan //阻塞协程，直到sigChan接收到值
		fmt.Println("11111")
		for value := range strChan {
			fmt.Println("接收到值为：", value)
		}
		sigChan1 <- struct{}{}
	}()
	go func() {
		//模拟发送数据
		for index, value := range []string{"1", "2", "3"} {
			fmt.Println("发送数据", value)
			strChan <- value
			if index == 2 {
				sigChan <- struct{}{}
				fmt.Println("sigChan", sigChan)
			}
		}
		close(strChan)
		sigChan1 <- struct{}{}
	}()
	fmt.Println("等待上面两个协程结束")
	//<-sigChan1
	//<-sigChan1 //阻塞直到上面两个协程完成
	time.Sleep(2 * time.Second)
}
