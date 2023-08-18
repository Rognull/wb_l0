package main

import(
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
)

func main(){
sc, err := stan.Connect("test-cluster", "publisher", stan.NatsURL("nats://localhost:4222"))

	if err != nil {
		panic(err)
		 
	}

	b,err := ioutil.ReadFile("model.json")

	if err != nil {
		fmt.Println("cant open file")
		return
	}


	err = sc.Publish("addOrder", b)

	if err != nil{
		fmt.Println(err.Error())
	}
	
	sc.Close()
}