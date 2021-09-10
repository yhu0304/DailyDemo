package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5*time.Second,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(),2*time.Second)
	_, err = cli.Put(ctx, "test", "henry")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	ctx, cancel = context.WithTimeout(context.Background(),2*time.Second)
	resp, err := cli.Get(ctx, "test")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, val := range resp.Kvs {
		fmt.Printf("%s:%s\n", val.Key, val.Value)
	}
}
