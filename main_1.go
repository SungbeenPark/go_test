package main

import (
	"fmt"
	accounts "github.com/SungbeenPark/go_test/accounts"
	"log"
)

//func multiply(a int, b int) int {
//	return a * b
//}
//
//func lenAndUpper(name string) (length int, uppercase string) {
//	defer fmt.Println("I'm done")
//	defer repeatMe("a", "b", "dd", "d", "e", "f")
//	length = len(name)
//	uppercase = strings.ToUpper(name)
//	return
//	//return len(name), strings.ToUpper(name)
//}
//
//func repeatMe(words ...string) {
//	fmt.Println(words)
//}
//
////for문
//func superAdd(numbers ...int) int {
//	total := 0
//	for _, number := range numbers {
//		total += number
//	}
//	return total
//}

//if/else문
//func canIDrink(age int) bool {
//
//	if koreanAge := age + 2; koreanAge < 18 {
//		fmt.Println(koreanAge)
//		return false
//	}
//	return true
//}
// switch문
func canIDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	//fmt.Println(multiply(2, 3))
	//totalLength, upperName := lenAndUpper("sungbeen")
	//fmt.Println(totalLength)
	//fmt.Println(upperName)

	//repeatMe("a", "b", "dd", "d", "e", "f")

	//result := superAdd(1, 2, 3, 4, 5, 6)
	//fmt.Println(result)
	//fmt.Println(canIDrink(16))

	//포인터
	//a := 2
	//b := &a
	//*b = 10
	//fmt.Println(a, *b)

	//names := []string{"a", "b", "c"}
	//names = append(names, "d")
	//fmt.Println(names)

	//been := map[string]string{"name": "been", "age": "12"}

	//for key, value := range been {
	//	fmt.Println(key, value)
	//}

	//sbpark := person{name: "sb", age: 37, favFood: []string{"ham", "pork"}}
	//fmt.Println(sbpark)

	account := accounts.NewAccount("been")
	//fmt.Println(accounts)
	//account2 := Account.NewAccount("bbbb")
	//fmt.Println(account2)
	//account3 := Account.NewAccount("222")
	//fmt.Println(account3)
	//fmt.Println(accounts)

	account.Deposit(30)
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)

}
