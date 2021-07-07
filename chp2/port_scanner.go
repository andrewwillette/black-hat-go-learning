package main

import (
	"fmt"
	"net"
	"sync"
)

func basic_scan() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("andrewwillette.com:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port is closed or filtered
			fmt.Printf("here1")
			continue
		}
		conn.Close()
		fmt.Printf("here2")
		fmt.Printf("%d open\n", i)
	}
}

func too_fast_scan() {
	for i := 1; i <= 1024; i++ {
		fmt.Printf("1")
		go func(j int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				// port is closed or filtered
				fmt.Printf("here1")
				return
			}
			conn.Close()
			fmt.Printf("here2")
			fmt.Printf("%d open\n", j)
		}(i)
	}
}

func wait_group_scan() {
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		fmt.Printf("1")
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				// port is closed or filtered
				fmt.Printf("here1")
				return
			}
			conn.Close()
			fmt.Printf("here2")
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait()
}

func worker(ports chan int, wg *sync.WaitGroup) {
	fmt.Printf("A")
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

func improved_wait_group_scan() {
	ports := make(chan int, 100)
	var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}

	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}

	wg.Wait()
	close(ports)
}
func main() {
	// basic_scan()
	//too_fast_scan()
	improved_wait_group_scan()
}
