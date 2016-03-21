package main
import (
	"fmt"
	"os"
	"encoding/csv"
)

func main(){
	f,err:=os.Open("./data.txt")
	if err!=nil{
		panic(err)
	}
	defer f.Close()

	reader:=csv.NewReader(f)
	rows,err:=reader.ReadAll()
	if err!=nil{
		panic(err)
	}

	for _,row:=range rows{
		fmt.Println(row)
	}

	fmt.Print(f)
}
