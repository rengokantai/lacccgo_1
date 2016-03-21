package main
import (
	"fmt"
	"os"
	"encoding/csv"
	"strconv"
)

func main(){
	f,err:=os.Open("./data.txt")
	if err!=nil{
		panic(err)
	}
	defer f.Close()

	reader:=csv.NewReader(f)
	reader.Comma = '\t'
	reader.TrimLeadingSpace = true
	rows,err:=reader.ReadAll()
	if err!=nil{
		panic(err)
	}

	for i,row:=range rows{
		//fmt.Println(row)
		if i !=0 &&i<5{
			fmt.Println(row[0])
			fmt.Printf("%T ",row[2])
			at,_:=strconv.ParseFloat(row[1],64)
			fmt.Printf("%T",at)


		}
	}

	//fmt.Print(f)
}
