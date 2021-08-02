package main

func main() {
	// #1
	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	// #2
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	// #3
	args := &example.Args{
		A: 10,
		B: 20,
	}

	// #4
	reply := &example.Reply{}

	// #5
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
