package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
 * 基于Channel编写一个简单的单线程生产者消费者模型：
 * 队列：
 *   队列长度10，队列元素类型为int
 * 生产者：
 *   每1秒往队列中放入一个类型为int的元素，队列满时生产者可以阻塞
 * 消费者：
 *   每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
*/


func main() {
	c := make(chan int, 10)

	go func() {  // 消费者
		for {
			num := <-c
			fmt.Printf("消费者读到数据： %d\n", num)
			time.Sleep(time.Second)
		}
	}()

	// 生成者
	for {
		c <- rand.Int()
		time.Sleep(time.Second)
	}
}
