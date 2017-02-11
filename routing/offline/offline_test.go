package offline

import (
	"bytes"
	"context"
	"github.com/ipfs/go-ipfs/thirdparty/testutil"
	ds "gx/ipfs/QmRWDav6mzWseLWeYfVd5fvUKiVe9xNH29YfMF438fG364/go-datastore"
	"testing"
)

func TestOfflineRouterStorage(t *testing.T) {
	ctx := context.Background()

	nds := ds.NewMapDatastore()
	privkey, _, _ := testutil.RandTestKeyPair(128)
	offline := NewOfflineRouter(nds, privkey)

	err := offline.PutValue(ctx, "key", []byte("testing 1 2 3"))
	if err != nil {
		t.Fatal(err)
	}

	val, err := offline.GetValue(ctx, "key")
	if !bytes.Equal([]byte("testing 1 2 3"), val) {
		t.Fatal("OfflineRouter does not properly store")
	}

	val, err = offline.GetValue(ctx, "notHere")
	if err == nil {
		t.Fatal("Router should throw errors for unfound records")
	}

	recVal, err := offline.GetValues(ctx, "key", 0)
	if err != nil {
		t.Fatal(err)
	}

	_, err = offline.GetValues(ctx, "notHere", 0)
	if err == nil {
		t.Fatal("Router should throw errors for unfound records")
	}

	local := recVal[0].Val
	if !bytes.Equal([]byte("testing 1 2 3"), local) {
		t.Fatal("OfflineRouter does not properly store")
	}
}

func TestOfflineRouterLocal(t *testing.T) {
	ctx := context.Background()

	nds := ds.NewMapDatastore()
	privkey, _, _ := testutil.RandTestKeyPair(128)
	offline := NewOfflineRouter(nds, privkey)

	id, _ := testutil.RandPeerID()
	_, err := offline.FindPeer(ctx, id)
	if err != ErrOffline {
		t.Fatal("OfflineRouting should alert that its offline")
	}

	cid, _ := testutil.RandCidV0()
	pChan := offline.FindProvidersAsync(ctx, cid, 1)
	p, ok := <-pChan
	if ok {
		t.Fatalf("FindProvidersAsync did not return a closed channel. Instead we got %+v !", p)
	}

	cid, _ = testutil.RandCidV0()
	err = offline.Provide(ctx, cid)
	if err != ErrOffline {
		t.Fatal("OfflineRouting should alert that its offline")
	}

	err = offline.Bootstrap(ctx)
	if err != nil {
		t.Fatal("You shouldn't be able to bootstrap offline routing.")
	}
}
