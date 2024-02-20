package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

func main() {
	config := clientv3.Config{
		Endpoints:   []string{"localhost:12379"},
		DialTimeout: 5 * time.Second,
	}
	// 建立连接
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	sender, receiver := "/sender_amount", "/receiver_amount"
	_, err = client.Put(context.Background(), sender, "1000")
	if err != nil {
		fmt.Printf("etcd put /sender_amount failed, err:%v\n", err)
		return
	}
	_, err = client.Put(context.Background(), receiver, "500")
	if err != nil {
		fmt.Printf("etcd put /receiver_amount failed, err:%v\n", err)
		return
	}

	err = txnStmTransfer(client, sender, receiver, 200)
	if err != nil {
		fmt.Printf("etcd txnTransfer failed, err:%v\n", err)
		return
	}
}

func txnStmTransfer(cli *clientv3.Client, from, to string, amount uint64) error {
	// NewSTM 创建了一个原子事务的上下文，并把我们的业务代码作为一个函数传进去
	_, err := concurrency.NewSTM(cli, func(stm concurrency.STM) error {
		// stm.Get 封装好了事务的读操作
		senderNum := StoUint64(stm.Get(from))
		receiverNum := StoUint64(stm.Get(to))
		if senderNum < amount {
			return fmt.Errorf("余额不足")
		}
		// stm.Put封装好了事务的写操作
		stm.Put(to, fromUint64(receiverNum+amount))
		stm.Put(from, fromUint64(senderNum-amount))
		return nil
	})
	return err
}
