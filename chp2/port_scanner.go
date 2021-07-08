package main

import (
	"fmt"
	"net"
	"sort"
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
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

func worker_multi(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

// to avoid inconsistencies with connecting to a lot of ports, use a worker pool
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

func multi_chan_port_scanner() {
	fmt.Println("here1")
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker_multi(ports, results)
	}

	fmt.Println("here2")
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	fmt.Println("here3")
	for i := 1; i <= 1000; i++ {
		fmt.Printf("here3.5 for port %f\n", i)
		port := <-results
		fmt.Println("here3.6")
		if port != 0 {
			openports = append(openports, port)
		}
	}

	fmt.Println("here4")
	close(ports)
	close(results)
	fmt.Println("before sort")
	fmt.Printf("%v\n", openports)
	sort.Ints(openports)
	fmt.Println("after sort")
	fmt.Printf("%v\n", openports)
	//for _, port := range openports {
	// fmt.Printf("%d open\n", port)
	//}
}

func main() {
	// basic_scan()
	//too_fast_scan()
	multi_chan_port_scanner()
}
