package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"

	"golang.org/x/net/html"
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
var visited = make(map[string]bool)
var semaphore = make(chan struct{}, 10)

func main() {

	// chanel := make(chan string, 10)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		chanel <- strconv.Itoa(i)
	// 	}
	// 	close(chanel)
	// }()
	// wg.Add(10)
	// for i := 0; i < 10; i++ {
	// 	go chanelremover(chanel)
	// }
	// wg.Wait()

	webCrawler("https://google.com", 2)

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

func webCrawler(starturl string, depth int) {
	fmt.Println("Currently Crawling := ", starturl)
	visited[starturl] = true
	//http get response

	resp, err := http.Get(starturl)
	if err != nil {
		fmt.Println("OOPS unable to get ", starturl)
	}

	links := linkcollector(resp, starturl)
	defer resp.Body.Close()

	for _, link := range links {
		wg.Add(1)
		go crawl(link, depth-1)

	}
	wg.Wait()

	// linkcollector

}
func crawl(link string, depth int) {

	if depth == 0 {
		fmt.Println("Depth Reached -> 0")
		return

	}

	mut.Lock()
	if visited[link] {
		mut.Unlock()
		fmt.Println("Already Visited this")
		return
	}
	fmt.Println("Currently Crawling := ", link)
	visited[link] = true
	mut.Unlock()
	wg.Done()

	//http get response
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("OOPS unable to get ", link)
		return
	}
	defer resp.Body.Close()
	links := linkcollector(resp, link)

	for _, url := range links {
		wg.Add(1)
		go crawl(url, depth-1)

	}

}
func linkcollector(resp *http.Response, link string) []string {
	links := []string{}
	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			t := tokenizer.Token()
			if t.Data != "a" {
				continue
			}
			for _, atr := range t.Attr {
				if atr.Key == "href" {
					href := atr.Val
					url := resolveUrl(href, link)
					if url != "" {
						fmt.Println("Link one", link, href)
						links = append(links, url)
					}

				}
			}
		}
	}
}

func resolveUrl(href, link string) string {
	u, err := url.Parse(link)
	if err != nil {
		fmt.Println(err)
	}
	url, err := u.Parse(href)

	if err != nil {
		fmt.Println(err)
	}
	return url.String()

}
