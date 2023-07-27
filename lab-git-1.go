package main
import "fmt"

func sumInt(ar ...int) (int, int) {
    var i int
    i = 0
    var sum int
    sum = 0
    for _, elem := range ar{
        sum += elem
        i++
    }
    return i, sum
}

func main(){
	fmt.Println("fff")
	a,b := sumInt(5,6,8)
	fmt.Println(a,b)
	fmt.Println(b,a)
}
