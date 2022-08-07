package main

import "fmt"
import "time"

func send(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
}

func receive(ch chan int) {
	var rec int

	for {
		rec = <-ch
		fmt.Println(rec)
	}
}

func main() {

	ch := make(chan int)

	go send(ch)
	go receive(ch)

	time.Sleep(time.Duration(2) * time.Second)

	//var s *Member
	//s = new(Member)
	//s.name = "aa"
	//fmt.Println(s.name)
	//	r := gin.Default()
	//	r.GET("/:message", func(c *gin.Context) {
	//		message := c.Param("message")
	//		query := c.Query("s")
	//
	//		c.JSON(http.StatusOK, gin.H{"message": message, "query": query})
	//	})
	//
	//	r.Run(":80")
}
