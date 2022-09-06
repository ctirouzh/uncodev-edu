package helpers

import (
	"strings"
	"math"
	"math/rand"
	"time"
	"encoding/json"
)


func PrettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
}



func Randint(max int,min int)int{
	rand.Seed(time.Now().UnixNano())
    return rand.Intn(max - min) + min
}


func NumberFormat(number string)string{
	str := ""
	arr := strings.Split(number, "")
	count := len(arr)
	newarr := []string{}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	for i:=0;i<count;i++{
		if math.Mod(float64(i),3)==0{
			newarr = append(newarr, ",")
		}
		newarr = append(newarr, arr[i])
	}
	for i:=len(newarr)-1;i>0;i--{
		str += newarr[i]
	}
	return str
}


