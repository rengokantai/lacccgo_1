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

	//var atTotal,bmTotal, wsTotal,counter float64

//	for i,row:=range rows{
//		//fmt.Println(row)
//		if i !=0 &&i<5{
//			//fmt.Println(row[0])
//			//fmt.Printf("%T ",row[2])
//			at,_:=strconv.ParseFloat(row[1],64)
//			bm,_:=strconv.ParseFloat(row[2],64)
//			ws,_:=strconv.ParseFloat(row[7],64)
//			//fmt.Printf("%T",at)
//
//			atTotal+=at
//			bmTotal+=bm
//			wsTotal+=ws
//			counter++
//
//
//		}
//	}
	fmt.Println("Total Records",len(rows)-1)
	fmt.Println("Mean Air Temp", mean(rows,1))
	fmt.Println("Mean Barometric",mean(rows,2))
	fmt.Println("Mean Wind Speed",mean(rows,7))

	//fmt.Print(f)
}


func mean(rows[][] string, idx int) float64{
	var total float64
	for i, row:=range rows{
		if i!=0{
			value,_:=strconv.ParseFloat(row[idx],64)
			total+=value
		}
	}
	return  total/float64(len(rows)-1)  //remove header line
}
