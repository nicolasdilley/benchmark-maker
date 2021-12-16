package main

// https://github.com/docker/swarmkit/blob/4b7bd25c452189a9d3164580e5db3d3d4c254ca4/protobuf/plugin/raftproxy/test/service.pb.go#L1472

func (p *raftProxyRouteGuideServer) RouteChat(stream RouteGuide_RouteChatServer) error {
	// code omitted
	errc := make(chan error, 1)
	go func() {
		msg, err := stream.Recv()
		if err == io.EOF {
			close(errc)
			return
		}
		if err != nil {
			errc <- err
			return
		}
		if err := clientStream.Send(msg); err != nil {
			errc <- err
			return
		}

		errc <- nil // make sure that we send something /*<\label{line:unblocked-send-1}>*/
	}()

	for {
		msg, err := clientStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if err := stream.Send(msg); err != nil {
			return err
		}
	}
	clientStream.CloseSend()
	return <-errc // block if no send
}