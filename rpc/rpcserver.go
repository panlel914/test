package rpcdemo


type RpcDemoServer struct {}

type Args struct {
	A,B int
}

func (RpcDemoServer) Add(args Args, result *int) error{
	*result = (args.A + args.B)
	return nil
}