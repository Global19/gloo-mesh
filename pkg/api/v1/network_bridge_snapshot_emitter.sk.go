// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sync"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.uber.org/zap"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	skstats "github.com/solo-io/solo-kit/pkg/stats"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errutils"
)

var (
	// Deprecated. See mNetworkBridgeResourcesIn
	mNetworkBridgeSnapshotIn = stats.Int64("network_bridge.zephyr.solo.io/emitter/snap_in", "Deprecated. Use network_bridge.zephyr.solo.io/emitter/resources_in. The number of snapshots in", "1")

	// metrics for emitter
	mNetworkBridgeResourcesIn    = stats.Int64("network_bridge.zephyr.solo.io/emitter/resources_in", "The number of resource lists received on open watch channels", "1")
	mNetworkBridgeSnapshotOut    = stats.Int64("network_bridge.zephyr.solo.io/emitter/snap_out", "The number of snapshots out", "1")
	mNetworkBridgeSnapshotMissed = stats.Int64("network_bridge.zephyr.solo.io/emitter/snap_missed", "The number of snapshots missed", "1")

	// views for emitter
	// deprecated: see networkBridgeResourcesInView
	networkBridgesnapshotInView = &view.View{
		Name:        "network_bridge.zephyr.solo.io/emitter/snap_in",
		Measure:     mNetworkBridgeSnapshotIn,
		Description: "Deprecated. Use network_bridge.zephyr.solo.io/emitter/resources_in. The number of snapshots updates coming in.",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}

	networkBridgeResourcesInView = &view.View{
		Name:        "network_bridge.zephyr.solo.io/emitter/resources_in",
		Measure:     mNetworkBridgeResourcesIn,
		Description: "The number of resource lists received on open watch channels",
		Aggregation: view.Count(),
		TagKeys: []tag.Key{
			skstats.NamespaceKey,
			skstats.ResourceKey,
		},
	}
	networkBridgesnapshotOutView = &view.View{
		Name:        "network_bridge.zephyr.solo.io/emitter/snap_out",
		Measure:     mNetworkBridgeSnapshotOut,
		Description: "The number of snapshots updates going out",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	networkBridgesnapshotMissedView = &view.View{
		Name:        "network_bridge.zephyr.solo.io/emitter/snap_missed",
		Measure:     mNetworkBridgeSnapshotMissed,
		Description: "The number of snapshots updates going missed. this can happen in heavy load. missed snapshot will be re-tried after a second.",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
)

func init() {
	view.Register(
		networkBridgesnapshotInView,
		networkBridgesnapshotOutView,
		networkBridgesnapshotMissedView,
		networkBridgeResourcesInView,
	)
}

type NetworkBridgeSnapshotEmitter interface {
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *NetworkBridgeSnapshot, <-chan error, error)
}

type NetworkBridgeEmitter interface {
	NetworkBridgeSnapshotEmitter
	Register() error
	MeshBridge() MeshBridgeClient
	Mesh() MeshClient
	MeshIngress() MeshIngressClient
}

func NewNetworkBridgeEmitter(meshBridgeClient MeshBridgeClient, meshClient MeshClient, meshIngressClient MeshIngressClient) NetworkBridgeEmitter {
	return NewNetworkBridgeEmitterWithEmit(meshBridgeClient, meshClient, meshIngressClient, make(chan struct{}))
}

func NewNetworkBridgeEmitterWithEmit(meshBridgeClient MeshBridgeClient, meshClient MeshClient, meshIngressClient MeshIngressClient, emit <-chan struct{}) NetworkBridgeEmitter {
	return &networkBridgeEmitter{
		meshBridge:  meshBridgeClient,
		mesh:        meshClient,
		meshIngress: meshIngressClient,
		forceEmit:   emit,
	}
}

type networkBridgeEmitter struct {
	forceEmit   <-chan struct{}
	meshBridge  MeshBridgeClient
	mesh        MeshClient
	meshIngress MeshIngressClient
}

func (c *networkBridgeEmitter) Register() error {
	if err := c.meshBridge.Register(); err != nil {
		return err
	}
	if err := c.mesh.Register(); err != nil {
		return err
	}
	if err := c.meshIngress.Register(); err != nil {
		return err
	}
	return nil
}

func (c *networkBridgeEmitter) MeshBridge() MeshBridgeClient {
	return c.meshBridge
}

func (c *networkBridgeEmitter) Mesh() MeshClient {
	return c.mesh
}

func (c *networkBridgeEmitter) MeshIngress() MeshIngressClient {
	return c.meshIngress
}

func (c *networkBridgeEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *NetworkBridgeSnapshot, <-chan error, error) {

	if len(watchNamespaces) == 0 {
		watchNamespaces = []string{""}
	}

	for _, ns := range watchNamespaces {
		if ns == "" && len(watchNamespaces) > 1 {
			return nil, nil, errors.Errorf("the \"\" namespace is used to watch all namespaces. Snapshots can either be tracked for " +
				"specific namespaces or \"\" AllNamespaces, but not both.")
		}
	}

	errs := make(chan error)
	var done sync.WaitGroup
	ctx := opts.Ctx
	/* Create channel for MeshBridge */
	type meshBridgeListWithNamespace struct {
		list      MeshBridgeList
		namespace string
	}
	meshBridgeChan := make(chan meshBridgeListWithNamespace)

	var initialMeshBridgeList MeshBridgeList
	/* Create channel for Mesh */
	type meshListWithNamespace struct {
		list      MeshList
		namespace string
	}
	meshChan := make(chan meshListWithNamespace)

	var initialMeshList MeshList
	/* Create channel for MeshIngress */
	type meshIngressListWithNamespace struct {
		list      MeshIngressList
		namespace string
	}
	meshIngressChan := make(chan meshIngressListWithNamespace)

	var initialMeshIngressList MeshIngressList

	currentSnapshot := NetworkBridgeSnapshot{}

	for _, namespace := range watchNamespaces {
		/* Setup namespaced watch for MeshBridge */
		{
			meshBridges, err := c.meshBridge.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial MeshBridge list")
			}
			initialMeshBridgeList = append(initialMeshBridgeList, meshBridges...)
		}
		meshBridgeNamespacesChan, meshBridgeErrs, err := c.meshBridge.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting MeshBridge watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, meshBridgeErrs, namespace+"-meshBridges")
		}(namespace)
		/* Setup namespaced watch for Mesh */
		{
			meshes, err := c.mesh.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial Mesh list")
			}
			initialMeshList = append(initialMeshList, meshes...)
		}
		meshNamespacesChan, meshErrs, err := c.mesh.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Mesh watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, meshErrs, namespace+"-meshes")
		}(namespace)
		/* Setup namespaced watch for MeshIngress */
		{
			meshIngresses, err := c.meshIngress.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial MeshIngress list")
			}
			initialMeshIngressList = append(initialMeshIngressList, meshIngresses...)
		}
		meshIngressNamespacesChan, meshIngressErrs, err := c.meshIngress.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting MeshIngress watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, meshIngressErrs, namespace+"-meshIngresses")
		}(namespace)

		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-ctx.Done():
					return
				case meshBridgeList := <-meshBridgeNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case meshBridgeChan <- meshBridgeListWithNamespace{list: meshBridgeList, namespace: namespace}:
					}
				case meshList := <-meshNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case meshChan <- meshListWithNamespace{list: meshList, namespace: namespace}:
					}
				case meshIngressList := <-meshIngressNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case meshIngressChan <- meshIngressListWithNamespace{list: meshIngressList, namespace: namespace}:
					}
				}
			}
		}(namespace)
	}
	/* Initialize snapshot for MeshBridges */
	currentSnapshot.MeshBridges = initialMeshBridgeList.Sort()
	/* Initialize snapshot for Meshes */
	currentSnapshot.Meshes = initialMeshList.Sort()
	/* Initialize snapshot for MeshIngresses */
	currentSnapshot.MeshIngresses = initialMeshIngressList.Sort()

	snapshots := make(chan *NetworkBridgeSnapshot)
	go func() {
		// sent initial snapshot to kick off the watch
		initialSnapshot := currentSnapshot.Clone()
		snapshots <- &initialSnapshot

		timer := time.NewTicker(time.Second * 1)
		previousHash, err := currentSnapshot.Hash(nil)
		if err != nil {
			contextutils.LoggerFrom(ctx).Panicw("error while hashing, this should never happen", zap.Error(err))
		}
		sync := func() {
			currentHash, err := currentSnapshot.Hash(nil)
			// this should never happen, so panic if it does
			if err != nil {
				contextutils.LoggerFrom(ctx).Panicw("error while hashing, this should never happen", zap.Error(err))
			}
			if previousHash == currentHash {
				return
			}

			sentSnapshot := currentSnapshot.Clone()
			select {
			case snapshots <- &sentSnapshot:
				stats.Record(ctx, mNetworkBridgeSnapshotOut.M(1))
				previousHash = currentHash
			default:
				stats.Record(ctx, mNetworkBridgeSnapshotMissed.M(1))
			}
		}
		meshBridgesByNamespace := make(map[string]MeshBridgeList)
		meshesByNamespace := make(map[string]MeshList)
		meshIngressesByNamespace := make(map[string]MeshIngressList)

		for {
			record := func() { stats.Record(ctx, mNetworkBridgeSnapshotIn.M(1)) }

			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				close(snapshots)
				done.Wait()
				close(errs)
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case meshBridgeNamespacedList := <-meshBridgeChan:
				record()

				namespace := meshBridgeNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"mesh_bridge",
					mNetworkBridgeResourcesIn,
				)

				// merge lists by namespace
				meshBridgesByNamespace[namespace] = meshBridgeNamespacedList.list
				var meshBridgeList MeshBridgeList
				for _, meshBridges := range meshBridgesByNamespace {
					meshBridgeList = append(meshBridgeList, meshBridges...)
				}
				currentSnapshot.MeshBridges = meshBridgeList.Sort()
			case meshNamespacedList := <-meshChan:
				record()

				namespace := meshNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"mesh",
					mNetworkBridgeResourcesIn,
				)

				// merge lists by namespace
				meshesByNamespace[namespace] = meshNamespacedList.list
				var meshList MeshList
				for _, meshes := range meshesByNamespace {
					meshList = append(meshList, meshes...)
				}
				currentSnapshot.Meshes = meshList.Sort()
			case meshIngressNamespacedList := <-meshIngressChan:
				record()

				namespace := meshIngressNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"mesh_ingress",
					mNetworkBridgeResourcesIn,
				)

				// merge lists by namespace
				meshIngressesByNamespace[namespace] = meshIngressNamespacedList.list
				var meshIngressList MeshIngressList
				for _, meshIngresses := range meshIngressesByNamespace {
					meshIngressList = append(meshIngressList, meshIngresses...)
				}
				currentSnapshot.MeshIngresses = meshIngressList.Sort()
			}
		}
	}()
	return snapshots, errs, nil
}