package logic

import (
	"context"
	"fmt"
	"github.com/GoCloudstorage/GoCloudstorage/pb/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestGrpcServer(t *testing.T) {
	dial, _ := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := storage.NewStorageClient(dial)
	resp, err := client.GetDownloadURL(context.Background(), &storage.GetDownloadURLReq{
		Hash:     "hhhhhhh",
		Filename: "",
		Ext:      "",
		Expire:   0,
	})
	fmt.Println(resp, err)
}
