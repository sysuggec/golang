package main

import (
    "fmt"
    "time"
    "github.com/garyburd/redigo/redis"
)


func main(){

   conn,err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Second,time.Second,time.Second) 
   if err == nil {
	//test set
	v, err := conn.Do("SET", "name", "red")
	fmt.Println(err)
	fmt.Println(v)
	
	//test get
	v,err = redis.String(conn.Do("GET","name"))
	fmt.Println(err)
	fmt.Println(v)
	
	//test pipe
	conn.Send("SET","name1","yellow")
	conn.Send("GET","name1")
	conn.Flush()
	v,err = conn.Receive()
	v2,err2 := redis.String(conn.Receive())
	fmt.Println(err)
	fmt.Println(v)
	fmt.Println(err2)
	fmt.Println(v2)

	v,err = conn.Do("ZADD","test_zset",100,"john")
	fmt.Println(err)
	fmt.Println(v)
        v,err = conn.Do("ZADD","test_zset",10,"jim")
        fmt.Println(err)
        fmt.Println(v)
	
	values,_ := redis.Values(conn.Do("ZRANGE","test_zset",0,-1,"withscores"))
	fmt.Println(values)
	for _ , x := range values {
		fmt.Println(string(x.([]byte)))
	}


   }else{
	fmt.Println("connect to redis fail")
	fmt.Println(err)
   }
}


