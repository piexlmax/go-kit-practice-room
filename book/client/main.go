
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"napodate/protobuf/book"
)

func main() {
	serviceAddress := "127.0.0.1:50052"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()
	bookClient := book.NewBookServiceClient(conn)
	bi, _ := bookClient.GetBookInfo(context.Background(), &book.BookInfoParams{BookId: 1})
	fmt.Print("获取书籍信息:\t")
	fmt.Println("bookInfo:", bi.BookName,bi.BookId)

	bl, _ := bookClient.GetBookList(context.Background(), &book.BookListParams{Page: 1, Limit: 10})
	fmt.Println("获取书籍列表:\t")
	for _, b := range bl.BookList {
		fmt.Println("bookId:", b.BookId, " => ", "bookName:", b.BookName)
	}
}