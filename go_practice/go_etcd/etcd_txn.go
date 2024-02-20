package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
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

	err = txnTransfer(client, sender, receiver, 200)
	if err != nil {
		fmt.Printf("etcd txnTransfer failed, err:%v\n", err)
		return
	}
}

func txnTransfer(etcd *clientv3.Client, sender, receiver string, amount uint64) error {
	// 失败重试
	for {
		if ok, err := doTxn(etcd, sender, receiver, amount); err != nil {
			return err
		} else if ok {
			return nil
		}
	}
}

func doTxn(etcd *clientv3.Client, sender, receiver string, amount uint64) (bool, error) {
	// 第一个事务，利用事务的原子性，同时获取发送和接收者的余额以及 ModRevision
	getResp, err := etcd.Txn(context.TODO()).Then(clientv3.OpGet(sender), clientv3.OpGet(receiver)).Commit()
	if err != nil {
		return false, err
	}
	senderKV := getResp.Responses[0].GetResponseRange().Kvs[0]
	receiverKV := getResp.Responses[1].GetResponseRange().Kvs[0]
	senderNum, receiverNum := BtoUInt64(senderKV.Value), BtoUInt64(receiverKV.Value)
	// 验证账户余额是否充足
	if senderNum < amount {
		return false, fmt.Errorf("资金不足")
	}
	// 发起转账事务，冲突判断 ModRevision 是否发生变化
	txn := etcd.Txn(context.TODO()).If(
		clientv3.Compare(clientv3.ModRevision(sender), "=", senderKV.ModRevision),
		clientv3.Compare(clientv3.ModRevision(receiver), "=", receiverKV.ModRevision))
	// ModRevision 未发生变化，即 If 判断条件成功
	txn = txn.Then(
		clientv3.OpPut(sender, fromUint64(senderNum-amount)),     // 更新发送者账户余额
		clientv3.OpPut(receiver, fromUint64(receiverNum+amount))) // 更新接收者账户余额
	resp, err := txn.Commit() // 提交事务
	if err != nil {
		return false, err
	}
	return resp.Succeeded, nil
}
