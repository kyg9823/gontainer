package clientutil

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/containerd/containerd/namespaces"
	containerd "github.com/containerd/containerd/v2/client"
	"golang.org/x/sys/unix"
)

func NewClient(address string, namespace string, opts ...containerd.Opt) (*containerd.Client, context.Context, context.CancelFunc, error) {
	ctx := namespaces.WithNamespace(context.Background(), namespace)

	address = strings.TrimPrefix(address, "unix://")
	const containerdAddress = "/var/run/containerd/containerd.sock"

	if err := IsSocketAccessible(address); err != nil {
		if IsSocketAccessible(containerdAddress) == nil {
			err = fmt.Errorf("cannot access containerd socket %q, but can access default socket %q: %w", address, containerdAddress, err)
		} else {
			err = fmt.Errorf("cannot access containerd socket %q: %w", address, err)
		}
		return nil, nil, nil, err
	}

	client, err := containerd.New(address, opts...)
	if err != nil {
		return nil, nil, nil, err
	}

	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	return client, ctx, cancel, nil
}

func IsSocketAccessible(s string) error {
	abs, err := filepath.Abs(s)
	if err != nil {
		return err
	}

	return unix.Faccessat(-1, abs, unix.R_OK|unix.W_OK, unix.AT_EACCESS)
}
