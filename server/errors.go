package server

import (
	"github.com/ryogrid/gord-overlay/chord"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func handleError(err error) error {
	code := status.Code(err)
	switch code {
	case codes.Unavailable:
		return chord.ErrNodeUnavailable
	case codes.NotFound:
		return chord.ErrNotFound
	default:
		return err
	}
}
