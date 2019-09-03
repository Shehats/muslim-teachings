package main

import (
	"log"
	"net"

	"github.com/muslim-teachings/quran-api/src/main/server"
	"google.golang.org/grpc"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/muslim-teachings/common/quran"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	db, err := gorm.Open("mysql", "root:ILoveAllah55*@tcp(127.0.0.1:3386)/Quran_db?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal("Failed to connect to DB, because of %v", err)
	}
	defer db.Close()
	server := new(server.QuranServer)
	server.DB = db
	// Creates a new gRPC server
	s := grpc.NewServer()
	quran.RegisterQuranServiceServer(s, server)
	s.Serve(lis)
}
