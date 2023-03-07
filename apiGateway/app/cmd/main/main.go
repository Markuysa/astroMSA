package main

var (
	authService    = ":9000"
	messageService = ":9001"
	astroService   = ":9002"
)

func main() {

}

//func HTTPProxy(proxyaddr string) {
//	grpcGwMux := runtime.NewServeMux()
//
//	// Connecting to auth service
//	grpcAuthConn, err := grpc.Dial(
//		authService,
//		//grpc.WithPerRPCCredentials(&reqData{}),
//		grpc.WithTransportCredentials(insecure.NewCredentials()),
//	)
//	if err != nil {
//		log.Fatalln("Filed to connect to User service", err)
//	}
//	defer grpcAuthConn.Close()
//
//	err = protobuf.RegisterAuthServiceHandler(
//		context.Background(),
//		grpcGwMux,
//		grpcAuthConn,
//	)
//	if err != nil {
//		log.Fatalln("Filed to start HTTP server", err)
//	}
//	// Connecting to auth service
//	grpcMessagesConn, err := grpc.Dial(
//		messageService,
//		//grpc.WithPerRPCCredentials(&reqData{}),
//		grpc.WithTransportCredentials(insecure.NewCredentials()),
//	)
//	if err != nil {
//		log.Fatalln("Filed to connect to Messages sender service", err)
//	}
//	defer grpcMessagesConn.Close()
//
//	err = protobuf.RegisterMessageServiceHandlerClient(
//		context.Background(),
//		grpcGwMux,
//		grpcMessagesConn,
//	)
//	if err != nil {
//		log.Fatalln("Filed to start HTTP server", err)
//	}
//	// Connecting to auth service
//	grpcAstroConn, err := grpc.Dial(
//		astroService,
//		//grpc.WithPerRPCCredentials(&reqData{}),
//		grpc.WithTransportCredentials(insecure.NewCredentials()),
//	)
//	if err != nil {
//		log.Fatalln("Filed to connect to User service", err)
//	}
//	defer grpcAstroConn.Close()
//
//	err = protobuf.RegisterAstrologyServiceHandler(
//		context.Background(),
//		grpcGwMux,
//		grpcAstroConn,
//	)
//	if err != nil {
//		log.Fatalln("Filed to start HTTP server", err)
//	}
//
//	mux := http.NewServeMux()
//
//	mux.Handle("/api/v1/", grpcGwMux)
//	mux.HandleFunc("/", helloworld)
//
//	fmt.Println("starting HTTP server at " + proxyaddr)
//	log.Fatal(http.ListenAndServe(proxyaddr, mux))
//}
//func helloworld(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "URL:", r.URL.String())
//}
//func main() {
//	HTTPProxy(":8080")
//}
