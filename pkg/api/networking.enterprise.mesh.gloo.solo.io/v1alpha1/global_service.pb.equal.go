// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/enterprise/networking/v1alpha1/global_service.proto

package v1alpha1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"

	v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)

	_ = v1alpha2.ApprovalState(0)
)

// Equal function
func (m *GlobalServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlobalServiceSpec)
	if !ok {
		that2, ok := that.(GlobalServiceSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetHostname(), target.GetHostname()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetPort()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPort()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPort(), target.GetPort()) {
			return false
		}
	}

	switch m.ExportTo.(type) {

	case *GlobalServiceSpec_VirtualMesh:
		if _, ok := target.ExportTo.(*GlobalServiceSpec_VirtualMesh); !ok {
			return false
		}

		if h, ok := interface{}(m.GetVirtualMesh()).(equality.Equalizer); ok {
			if !h.Equal(target.GetVirtualMesh()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetVirtualMesh(), target.GetVirtualMesh()) {
				return false
			}
		}

	case *GlobalServiceSpec_MeshList_:
		if _, ok := target.ExportTo.(*GlobalServiceSpec_MeshList_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetMeshList()).(equality.Equalizer); ok {
			if !h.Equal(target.GetMeshList()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetMeshList(), target.GetMeshList()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.ExportTo != target.ExportTo {
			return false
		}
	}

	switch m.FailoverConfig.(type) {

	case *GlobalServiceSpec_Static:
		if _, ok := target.FailoverConfig.(*GlobalServiceSpec_Static); !ok {
			return false
		}

		if h, ok := interface{}(m.GetStatic()).(equality.Equalizer); ok {
			if !h.Equal(target.GetStatic()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetStatic(), target.GetStatic()) {
				return false
			}
		}

	case *GlobalServiceSpec_Localized:
		if _, ok := target.FailoverConfig.(*GlobalServiceSpec_Localized); !ok {
			return false
		}

		if h, ok := interface{}(m.GetLocalized()).(equality.Equalizer); ok {
			if !h.Equal(target.GetLocalized()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetLocalized(), target.GetLocalized()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.FailoverConfig != target.FailoverConfig {
			return false
		}
	}

	return true
}

// Equal function
func (m *GlobalServiceStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlobalServiceStatus)
	if !ok {
		that2, ok := that.(GlobalServiceStatus)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetObservedGeneration() != target.GetObservedGeneration() {
		return false
	}

	if m.GetState() != target.GetState() {
		return false
	}

	if len(m.GetMeshes()) != len(target.GetMeshes()) {
		return false
	}
	for k, v := range m.GetMeshes() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetMeshes()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetMeshes()[k]) {
				return false
			}
		}

	}

	if len(m.GetSelectedTrafficTargets()) != len(target.GetSelectedTrafficTargets()) {
		return false
	}
	for idx, v := range m.GetSelectedTrafficTargets() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelectedTrafficTargets()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSelectedTrafficTargets()[idx]) {
				return false
			}
		}

	}

	if len(m.GetErrors()) != len(target.GetErrors()) {
		return false
	}
	for idx, v := range m.GetErrors() {

		if strings.Compare(v, target.GetErrors()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *GlobalServiceSpec_Port) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlobalServiceSpec_Port)
	if !ok {
		that2, ok := that.(GlobalServiceSpec_Port)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetNumber() != target.GetNumber() {
		return false
	}

	if strings.Compare(m.GetProtocol(), target.GetProtocol()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *GlobalServiceSpec_MeshList) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlobalServiceSpec_MeshList)
	if !ok {
		that2, ok := that.(GlobalServiceSpec_MeshList)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetMeshes()) != len(target.GetMeshes()) {
		return false
	}
	for idx, v := range m.GetMeshes() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetMeshes()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetMeshes()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *GlobalServiceSpec_BackingServiceList) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlobalServiceSpec_BackingServiceList)
	if !ok {
		that2, ok := that.(GlobalServiceSpec_BackingServiceList)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetServices()) != len(target.GetServices()) {
		return false
	}
	for idx, v := range m.GetServices() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetServices()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetServices()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *GlobalServiceSpec_LocalityConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlobalServiceSpec_LocalityConfig)
	if !ok {
		that2, ok := that.(GlobalServiceSpec_LocalityConfig)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetServiceSelectors()) != len(target.GetServiceSelectors()) {
		return false
	}
	for idx, v := range m.GetServiceSelectors() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetServiceSelectors()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetServiceSelectors()[idx]) {
				return false
			}
		}

	}

	if len(m.GetLocalities()) != len(target.GetLocalities()) {
		return false
	}
	for idx, v := range m.GetLocalities() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetLocalities()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetLocalities()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *GlobalServiceSpec_BackingServiceList_BackingService) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlobalServiceSpec_BackingServiceList_BackingService)
	if !ok {
		that2, ok := that.(GlobalServiceSpec_BackingServiceList_BackingService)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.BackingServiceType.(type) {

	case *GlobalServiceSpec_BackingServiceList_BackingService_KubeService:
		if _, ok := target.BackingServiceType.(*GlobalServiceSpec_BackingServiceList_BackingService_KubeService); !ok {
			return false
		}

		if h, ok := interface{}(m.GetKubeService()).(equality.Equalizer); ok {
			if !h.Equal(target.GetKubeService()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetKubeService(), target.GetKubeService()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.BackingServiceType != target.BackingServiceType {
			return false
		}
	}

	return true
}

// Equal function
func (m *GlobalServiceSpec_LocalityConfig_Locality) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlobalServiceSpec_LocalityConfig_Locality)
	if !ok {
		that2, ok := that.(GlobalServiceSpec_LocalityConfig_Locality)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetFrom(), target.GetFrom()) != 0 {
		return false
	}

	if len(m.GetTo()) != len(target.GetTo()) {
		return false
	}
	for idx, v := range m.GetTo() {

		if strings.Compare(v, target.GetTo()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *GlobalServiceStatus_SelectedTrafficTarget) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlobalServiceStatus_SelectedTrafficTarget)
	if !ok {
		that2, ok := that.(GlobalServiceStatus_SelectedTrafficTarget)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRef(), target.GetRef()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetService()).(equality.Equalizer); ok {
		if !h.Equal(target.GetService()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetService(), target.GetService()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GlobalServiceStatus_SelectedTrafficTarget_BackingService) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GlobalServiceStatus_SelectedTrafficTarget_BackingService)
	if !ok {
		that2, ok := that.(GlobalServiceStatus_SelectedTrafficTarget_BackingService)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.BackingServiceType.(type) {

	case *GlobalServiceStatus_SelectedTrafficTarget_BackingService_KubeService:
		if _, ok := target.BackingServiceType.(*GlobalServiceStatus_SelectedTrafficTarget_BackingService_KubeService); !ok {
			return false
		}

		if h, ok := interface{}(m.GetKubeService()).(equality.Equalizer); ok {
			if !h.Equal(target.GetKubeService()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetKubeService(), target.GetKubeService()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.BackingServiceType != target.BackingServiceType {
			return false
		}
	}

	return true
}