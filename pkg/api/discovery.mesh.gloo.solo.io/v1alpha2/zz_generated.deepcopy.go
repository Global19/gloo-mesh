// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for discovery.mesh.gloo.solo.io/v1alpha2 resources

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for TrafficTarget

func (in *TrafficTarget) DeepCopyInto(out *TrafficTarget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *TrafficTarget) DeepCopy() *TrafficTarget {
	if in == nil {
		return nil
	}
	out := new(TrafficTarget)
	in.DeepCopyInto(out)
	return out
}

func (in *TrafficTarget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *TrafficTargetList) DeepCopyInto(out *TrafficTargetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TrafficTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *TrafficTargetList) DeepCopy() *TrafficTargetList {
	if in == nil {
		return nil
	}
	out := new(TrafficTargetList)
	in.DeepCopyInto(out)
	return out
}

func (in *TrafficTargetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for Workload

func (in *Workload) DeepCopyInto(out *Workload) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *Workload) DeepCopy() *Workload {
	if in == nil {
		return nil
	}
	out := new(Workload)
	in.DeepCopyInto(out)
	return out
}

func (in *Workload) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *WorkloadList) DeepCopyInto(out *WorkloadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Workload, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *WorkloadList) DeepCopy() *WorkloadList {
	if in == nil {
		return nil
	}
	out := new(WorkloadList)
	in.DeepCopyInto(out)
	return out
}

func (in *WorkloadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for Mesh

func (in *Mesh) DeepCopyInto(out *Mesh) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *Mesh) DeepCopy() *Mesh {
	if in == nil {
		return nil
	}
	out := new(Mesh)
	in.DeepCopyInto(out)
	return out
}

func (in *Mesh) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *MeshList) DeepCopyInto(out *MeshList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Mesh, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *MeshList) DeepCopy() *MeshList {
	if in == nil {
		return nil
	}
	out := new(MeshList)
	in.DeepCopyInto(out)
	return out
}

func (in *MeshList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}