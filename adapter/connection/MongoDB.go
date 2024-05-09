package connection

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connn() *mongo.Database {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	basecodeuri := os.Getenv("MONGO_URI")
	if basecodeuri == "" {
		// Xử lý trường hợp không tìm thấy biến môi trường
		fmt.Println("NO ENV variable")
	}
	decodeuri, err := base64.StdEncoding.DecodeString(basecodeuri)
	if err != nil {
		// Xử lý lỗi khi giải mã
		fmt.Println("Lỗi khi giải mã base64:", err)
	}
	uri := string(decodeuri)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	database := client.Database("test")
	return database
}
