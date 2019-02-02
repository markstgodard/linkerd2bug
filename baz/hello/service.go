package hello

import "context"

type Service struct {
}

func (s *Service) Hello(context.Context, *SayHelloRequest) (*SayHelloResponse, error) {
	resp := &SayHelloResponse{
		Msg: "hello from baz",
	}

	return resp, nil
}
