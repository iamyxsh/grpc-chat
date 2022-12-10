module github.com/iamyxsh/grpc-chat/chat

go 1.19

require (
	github.com/iamyxsh/grpc-chat/kafka v0.0.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/lib/pq v1.10.7
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
)

replace github.com/iamyxsh/grpc-chat/kafka => ../kafka

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/segmentio/kafka-go v0.4.38 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)
