// +build !ignore_autogenerated

/*
Copyright 2019 Gaurav Gupta.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v2beta2 "k8s.io/api/autoscaling/v2beta2"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1beta2 "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1beta2"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChangeThreshold) DeepCopyInto(out *ChangeThreshold) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.Percentage != nil {
		in, out := &in.Percentage, &out.Percentage
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChangeThreshold.
func (in *ChangeThreshold) DeepCopy() *ChangeThreshold {
	if in == nil {
		return nil
	}
	out := new(ChangeThreshold)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HpaStatus) DeepCopyInto(out *HpaStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HpaStatus.
func (in *HpaStatus) DeepCopy() *HpaStatus {
	if in == nil {
		return nil
	}
	out := new(HpaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HpaTemplateSpec) DeepCopyInto(out *HpaTemplateSpec) {
	*out = *in
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]v2beta2.MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HpaTemplateSpec.
func (in *HpaTemplateSpec) DeepCopy() *HpaTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(HpaTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hvpa) DeepCopyInto(out *Hvpa) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hvpa.
func (in *Hvpa) DeepCopy() *Hvpa {
	if in == nil {
		return nil
	}
	out := new(Hvpa)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Hvpa) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HvpaCurrentStatus) DeepCopyInto(out *HvpaCurrentStatus) {
	*out = *in
	if in.LastScaleTime != nil {
		in, out := &in.LastScaleTime, &out.LastScaleTime
		*out = (*in).DeepCopy()
	}
	out.LastScaleType = in.LastScaleType
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HvpaCurrentStatus.
func (in *HvpaCurrentStatus) DeepCopy() *HvpaCurrentStatus {
	if in == nil {
		return nil
	}
	out := new(HvpaCurrentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HvpaList) DeepCopyInto(out *HvpaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Hvpa, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HvpaList.
func (in *HvpaList) DeepCopy() *HvpaList {
	if in == nil {
		return nil
	}
	out := new(HvpaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HvpaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HvpaSpec) DeepCopyInto(out *HvpaSpec) {
	*out = *in
	in.HpaTemplate.DeepCopyInto(&out.HpaTemplate)
	in.VpaTemplate.DeepCopyInto(&out.VpaTemplate)
	if in.WeightBasedScalingIntervals != nil {
		in, out := &in.WeightBasedScalingIntervals, &out.WeightBasedScalingIntervals
		*out = make([]WeightBasedScalingInterval, len(*in))
		copy(*out, *in)
	}
	if in.TargetRef != nil {
		in, out := &in.TargetRef, &out.TargetRef
		*out = new(v2beta2.CrossVersionObjectReference)
		**out = **in
	}
	if in.ScaleUpStabilizationWindow != nil {
		in, out := &in.ScaleUpStabilizationWindow, &out.ScaleUpStabilizationWindow
		*out = new(string)
		**out = **in
	}
	if in.ScaleDownStabilizationWindow != nil {
		in, out := &in.ScaleDownStabilizationWindow, &out.ScaleDownStabilizationWindow
		*out = new(string)
		**out = **in
	}
	if in.MinCPUChange != nil {
		in, out := &in.MinCPUChange, &out.MinCPUChange
		*out = new(ChangeThreshold)
		(*in).DeepCopyInto(*out)
	}
	if in.MinMemChange != nil {
		in, out := &in.MinMemChange, &out.MinMemChange
		*out = new(ChangeThreshold)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HvpaSpec.
func (in *HvpaSpec) DeepCopy() *HvpaSpec {
	if in == nil {
		return nil
	}
	out := new(HvpaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HvpaStatus) DeepCopyInto(out *HvpaStatus) {
	*out = *in
	out.HpaStatus = in.HpaStatus
	in.VpaStatus.DeepCopyInto(&out.VpaStatus)
	in.HvpaStatus.DeepCopyInto(&out.HvpaStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HvpaStatus.
func (in *HvpaStatus) DeepCopy() *HvpaStatus {
	if in == nil {
		return nil
	}
	out := new(HvpaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastScaleType) DeepCopyInto(out *LastScaleType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastScaleType.
func (in *LastScaleType) DeepCopy() *LastScaleType {
	if in == nil {
		return nil
	}
	out := new(LastScaleType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpaTemplateSpec) DeepCopyInto(out *VpaTemplateSpec) {
	*out = *in
	if in.ResourcePolicy != nil {
		in, out := &in.ResourcePolicy, &out.ResourcePolicy
		*out = new(v1beta2.PodResourcePolicy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpaTemplateSpec.
func (in *VpaTemplateSpec) DeepCopy() *VpaTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(VpaTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeightBasedScalingInterval) DeepCopyInto(out *WeightBasedScalingInterval) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeightBasedScalingInterval.
func (in *WeightBasedScalingInterval) DeepCopy() *WeightBasedScalingInterval {
	if in == nil {
		return nil
	}
	out := new(WeightBasedScalingInterval)
	in.DeepCopyInto(out)
	return out
}
