package main
import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)




func main() {
	res,err:=http.Get("http://www.github.com")
	if err!=nil{
		log.Fatal(err)
	}
	page,err:=ioutil.ReadAll(res.Body)  //compare to main.go line 21
	res.Body.Close()
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s",page)  //original
	fmt.Printf("%v",page)   //print bytes, like 111 101 111
}