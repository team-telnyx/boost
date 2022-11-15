package svc

import (
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	cl "github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/filecoin-project/boostd-data/client"
	"github.com/filecoin-project/boostd-data/model"
	"github.com/ipfs/go-cid"
	"golang.org/x/net/context"
)

func TestCouchbaseService(t *testing.T) {
	setupCouchbase(t)

	bdsvc := NewCouchbase(testCouchSettings)
	err := bdsvc.Start(context.Background(), 8042)
	if err != nil {
		t.Fatal(err)
	}

	pdcl := client.NewStore()
	err = pdcl.Dial(context.Background(), "http://localhost:8042")
	defer pdcl.Close(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	pieceCid, err := cid.Parse("baga6ea4seaqj2j4zfi2xk7okc7fnuw42pip6vjv2tnc4ojsbzlt3rfrdroa7qly")
	if err != nil {
		t.Fatal(err)
	}
	dealInfo := model.DealInfo{}
	err = pdcl.AddDealForPiece(context.TODO(), pieceCid, dealInfo)
	if err != nil {
		t.Fatal(err)
	}

	log.Debug("sleeping for a while.. running tests..")
	time.Sleep(2 * time.Second)
}

func setupCouchbase(t *testing.T) {
	ctx := context.Background()
	cli, err := cl.NewClientWithOpts()
	require.NoError(t, err)

	imageName := "couchbase"

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	require.NoError(t, err)
	_, err = io.Copy(os.Stdout, out)
	require.NoError(t, err)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		ExposedPorts: nat.PortSet{
			"8091":  struct{}{},
			"8092":  struct{}{},
			"8093":  struct{}{},
			"8094":  struct{}{},
			"8095":  struct{}{},
			"8096":  struct{}{},
			"11210": struct{}{},
			"11211": struct{}{},
		},
	}, &container.HostConfig{
		PortBindings: map[nat.Port][]nat.PortBinding{
			nat.Port("8091"):  {{HostIP: "127.0.0.1", HostPort: "8091"}},
			nat.Port("8092"):  {{HostIP: "127.0.0.1", HostPort: "8092"}},
			nat.Port("8093"):  {{HostIP: "127.0.0.1", HostPort: "8093"}},
			nat.Port("8094"):  {{HostIP: "127.0.0.1", HostPort: "8094"}},
			nat.Port("8095"):  {{HostIP: "127.0.0.1", HostPort: "8095"}},
			nat.Port("8096"):  {{HostIP: "127.0.0.1", HostPort: "8096"}},
			nat.Port("11210"): {{HostIP: "127.0.0.1", HostPort: "11210"}},
			nat.Port("11211"): {{HostIP: "127.0.0.1", HostPort: "11211"}},
		},
	}, nil, nil, "")
	require.NoError(t, err)

	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{Force: true})
	})

	command := []string{""}
	execCreateResp, err := cli.ContainerExecCreate(ctx, resp.ID, types.ExecConfig{Cmd: command})
	require.NoError(t, err)

	execResp, err := cli.ContainerExecAttach(ctx, execCreateResp.ID, types.ExecStartCheck{})
	require.NoError(t, err)
	defer execResp.Close()

	// TODO: setup admin password
	// TODO: create bucket -- see: https://docs.couchbase.com/server/current/manage/manage-buckets/create-bucket.html
}
