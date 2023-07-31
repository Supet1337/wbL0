package main

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"log"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	sc, err := stan.Connect("test-cluster", "client2")
	if err != nil {
		log.Fatalf("can not connect to nats streaming: %v", err)
	}
	defer sc.Close()

	data := []byte("{\n    \"ordenr_uid\": \"1234564\",\n    \"track_number\": \"ABC123\",\n    \"entry\": \"entry\",\n    \"delivery\": {\n        \"name\": \"John Doe\",\n        \"phone\": \"555-555-5555\",\n        \"zip\": \"12345\",\n        \"city\": \"New York\",\n        \"address\": \"123 Main St\",\n        \"region\": \"NY\",\n        \"email\": \"johndoe@example.com\"\n    },\n    \"payment\": {\n        \"transaction\": \"123456\",\n        \"request_id\": \"7890\",\n        \"currency\": \"USD\",\n        \"provider\": \"PayPal\",\n        \"amount\": 100,\n        \"payment_dt\": 1627470326,\n        \"bank\": \"Bank of America\",\n        \"delivery_cost\": 10,\n        \"goods_total\": 90,\n        \"custom_fee\": 5\n    },\n    \"items\": [\n        {\n            \"chrt_id\": 1,\n            \"track_number\": \"ABC123\",\n            \"price\": 50,\n            \"rid\": \"54321\",\n            \"name\": \"Product 1\",\n            \"sale\": 10,\n            \"size\": \"M\",\n            \"total_price\": 55,\n            \"nm_id\": 123,\n            \"brand\": \"Brand 1\",\n            \"status\": 1\n        },\n        {\n            \"chrt_id\": 2,\n            \"track_number\": \"ABC123\",\n            \"price\": 40,\n            \"rid\": \"98765\",\n            \"name\": \"Product 2\",\n            \"sale\": 0,\n            \"size\": \"L\",\n            \"total_price\": 40,\n            \"nm_id\": 456,\n            \"brand\": \"Brand 2\",\n            \"status\": 2\n        }\n    ],\n    \"locale\": \"en_US\",\n    \"internal_signature\": \"abc123\",\n    \"customer_id\": \"789\",\n    \"delivery_service\": \"FedEx\",\n    \"shardkey\": \"shardkey\",\n    \"sm_id\": 12345,\n    \"date_created\": \"2022-01-01T00:00:00Z\",\n    \"oof_shard\": \"oof_shard\"\n}")
	err = sc.Publish("topic", data)
	log.Println("PUBLISHING")
	if err != nil {
		log.Println(err)
	}
}
