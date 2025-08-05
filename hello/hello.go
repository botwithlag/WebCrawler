package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type vehicle struct {
	make   string
	year   int16
	wheels int8
}

type f struct {
	x, y int
}

type notastruct struct {
	honda vehicle
}

type Order struct {
	id     int
	status string
}

func updatingOrders(orders []*Order) {
	for i := 0; i < len(orders); i++ {
		status := [3]string{"completed", "failure", "Aborted"}[rand.Intn(3)]
		orders[i].status = status
		fmt.Printf("completed Updating Order for %v \n", orders[i].id)

	}

}

func recieptorders(orders []*Order) {
	for index, order := range orders {
		time.Sleep(time.Duration(100))
		fmt.Printf("Order :%v status : %v \n", index, order.status)
	}
}
func stringarrayconverted(str string) []string {
	s := ""
	var str_arr []string
	count := 0
	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " {
			if s == "" {
				continue
			}
			str_arr = append(str_arr, s)
			count++
			s = ""
			continue
		}

		if int(str[i]) == (len(str) - 1) {
			if s == "" {
				break
			}
			str_arr = append(str_arr, s)
		}
		s += string(str[i])

	}
	return str_arr
}

func generateOrder(count int) []*Order {
	orders := make([]*Order, count)
	for index := range orders {
		time.Sleep(time.Duration(100))
		orders[index] = &Order{
			id:     index + 1,
			status: "pending",
		}
		fmt.Printf("Complpeted generating order %v \n", orders[index].id)

	}
	return orders
}

func routineexample(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

var wg sync.WaitGroup
var mut sync.Mutex

func main() {

	chanel := make(chan string, 10)
	go func() {
		for i := 0; i < 10; i++ {
			chanel <- strconv.Itoa(i)
		}
		close(chanel)
	}()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go chanelremover(chanel)
	}
	wg.Wait()

	// recieptorders(orders)
	// var intptr *int
	// p := 20
	// intptr = &p
	// fmt.Println(*intptr)

	// var f_arr1 = []f{{1, 2}}
	// f_arr1[0] = f{0, 2}

	// str2 := stringarrayconverted("    1235t4     Greeting ")
	// fmt.Print(str2)
	// str3 := Channels2()

	// fmt.Println(str3)

	// d := make([]int, 5)
	// d = d[:0]
	// fmt.Println(d)
	// d = d[1:3]
	// fmt.Println(d)

	// // Get a greeting message and print it.
	// message, err := greetings.Hello("2378")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// str := strings.Fields("This           is a message ")
	// fmt.Print(str)
	// fmt.Print(len(str))
	// defer fmt.Println(message)
	// var num int = 2
	// fmt.Println(num)
	// var car = vehicle{"make", 291, 2}
	// car.make = "SUV"

}
func chanelremover(chanel chan string) {
	defer wg.Done()
	str := <-chanel
	fmt.Println(str)

}

func webCrawler(starturl string, depth int8) {
	if depth == 0 {
		return
	}

	//http get response
	resp, err := http.Get(starturl)
	if err != nil {
		fmt.Println("OOPS unable to get ", starturl)
	}
	defer resp.Body.Close()
	links := linkcollector(resp)

	for _, link := range links {
		wg.Add(1)
		go webCrawler(link, depth-1)
	}

	// linkcollector

}

func linkcollector(resp *http.Response) (links []string) {
	return
}
