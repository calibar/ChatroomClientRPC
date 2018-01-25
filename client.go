package main

import (
	"net/rpc"
	"log"
	"fmt"

	"time"
)
type Args struct {
	A string
	UID string
	Stime time.Time
	Room string
}

type ReadingReply struct {
	ID int
	Content string
}

var err error
var client *rpc.Client
var name string
var Room string
var timeout int

func sending(client1 *rpc.Client)  {
	for {
		if timeout!=-1 {
			var message string
			var input string
			fmt.Scanln(&input)
			t:=time.Now()
			fmt.Println("      "+t.Format(time.RFC850))
			message = name+" says: "+input+"\n"+"      "+t.Format(time.RFC850)
			args := Args{message,name,t,Room}
			var reply string
			err = client1.Call("Arith.ReceiveMessage", args, &reply)
			if err != nil {
				log.Fatal("arith error:", err)
			}
		}else {
			fmt.Println("Time up! See you in other room!")
			break
		}

	}
}
func reading(client2 *rpc.Client)  {
	var currentID = 0
	var t1 time.Time
	args := Args{"hi",name,t1,Room}
	var readingreply ReadingReply
	for true  {
		err = client2.Call("Arith.Reading", args, &readingreply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		if readingreply.ID!=-1 {
			if readingreply.ID!=currentID {
				currentID=readingreply.ID
				fmt.Println(readingreply.Content)
			}
		}else {
			fmt.Println(readingreply.Content)
			timeout=-1
			break
		}

		time.Sleep(10*time.Millisecond)
	}
}
func callHistory(client3 *rpc.Client)  {
	var t1 time.Time
	args := Args{"hi",name,t1,Room}
	var reply3 string
	err = client3.Call("Arith.Show", args, &reply3)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println(reply3)
	fmt.Println("--------------------History--------------------")
}
func callShowroom(clients *rpc.Client){
	var t2 time.Time
	args := Args{"hi",name,t2,Room}
	var replys string
	err = clients.Call("Arith.Showroom", args, &replys)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println(replys)
}
func callCreatroom(clienty *rpc.Client)  {
	var t2 time.Time
	args := Args{"",name,t2,Room}
	var replys string
	err = clienty.Call("Arith.Creatroom", args, &replys)
	if err != nil {
		log.Fatal("arith error:", err)
	}
}

func main()  {

	serverAddress := "127.0.0.1"
	client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	fmt.Println("What is your name?")
	fmt.Scanln(&name)
	fmt.Println("which room do you want in?")
	callShowroom(client)
	fmt.Scanln(&Room)
	callCreatroom(client)
	fmt.Println("Welcome to Room: "+Room+"! "+name)
	callHistory(client)
	go reading(client)
	sending(client)
	time.Sleep(100*time.Minute)
}


