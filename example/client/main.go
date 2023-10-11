// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Daz-3ux/dBlog/internal/pkg/log"
	pb "github.com/Daz-3ux/dBlog/pkg/proto/dazBlog/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

var (
	addr  = flag.String("addr", "localhost:9090", "the address to connect to.")
	limit = flag.Int64("limit", 10, "Limit to list users.")
)

func main() {
	flag.Parse()
	// connect to the server
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalw("Did not connect:", "err", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalw("Could not close connection:", "err", err)
		}
	}(conn)
	c := pb.NewDazBlogClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call the ListUsers method
	r, err := c.ListUsers(ctx, &pb.ListUsersRequest{Offset: 0, Limit: *limit})
	if err != nil {
		log.Fatalw("Could not list users:", "err", err)
	}

	// print the response
	fmt.Println("TotalCount:", r.TotalCount)
	for _, u := range r.Users {
		d, _ := json.Marshal(u)
		fmt.Println(string(d))
	}
}
