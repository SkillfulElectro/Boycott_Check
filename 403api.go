package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
thanks to www.anti403.ir for providing this API

this application just recives data from anti403 database and checks if the website is under boycott or not
based on the result recived from database !
*/

func main() {
	var WebAddr string

	switch {
	case len(os.Args) == 1:
		fmt.Println("run the application with a web address as the last argument.\nexample: anti403.exe www.google.com")
		return
	default:
		WebAddr = os.Args[len(os.Args)-1]
	}

	res, err := http.Get("https://api.anti403.ir/api/search-filter?url=" + WebAddr)
	if err != nil {
		fmt.Println("checking failed : ", err)
		return
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("checking failed : ", err)
		return
	}

	if len(b) > 130 {
		fmt.Println("This website is under boycott")
	} else {
		fmt.Println("This website is not under boycott")
		fmt.Println("even though it is not under boycott, checking if the service is reachable! . . .")
		fmt.Println("this method is not specified just for Iran")
		res, err = http.Get(WebAddr)

		if err != nil {
			fmt.Println("this website is banned in your country : ", err)
			return
		}

		if res.StatusCode != 200 {
			fmt.Println("This website might not be reachable!")
		} else {
			fmt.Println("This website wont cause you any problem")
		}
	}
}

/// Garbage dump:
/*
func stacks_onWar() {
	for {
		res, err := http.Post("https://api.anti403.ir/api/search-filter?url=www.ygyuff.com", "form", strings.NewReader("I want to learn more and more about networks"))
		if err != nil {
			fmt.Println(err)
			return
		}
		b, _ := io.ReadAll(res.Body)
		fmt.Println(string(b))
		time.Sleep(time.Second)
	}
}
*/
/*
	for {
		time.Sleep(time.Second * 10)
		go stacks_onWar()
	}
*/
//	res, _ = http.Get("https://anti403.ir/")
//	b, _ = io.ReadAll(res.Body)
//	fmt.Println(string(b))
