//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Agent) DeepCopyInto(out *Agent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Agent.
func (in *Agent) DeepCopy() *Agent {
	if in == nil {
		return nil
	}
	out := new(Agent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Agent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentInitParameters) DeepCopyInto(out *AgentInitParameters) {
	*out = *in
	if in.AgentName != nil {
		in, out := &in.AgentName, &out.AgentName
		*out = new(string)
		**out = **in
	}
	if in.AgentResourceRoleArn != nil {
		in, out := &in.AgentResourceRoleArn, &out.AgentResourceRoleArn
		*out = new(string)
		**out = **in
	}
	if in.AgentResourceRoleArnRef != nil {
		in, out := &in.AgentResourceRoleArnRef, &out.AgentResourceRoleArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AgentResourceRoleArnSelector != nil {
		in, out := &in.AgentResourceRoleArnSelector, &out.AgentResourceRoleArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomerEncryptionKeyArn != nil {
		in, out := &in.CustomerEncryptionKeyArn, &out.CustomerEncryptionKeyArn
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FoundationModel != nil {
		in, out := &in.FoundationModel, &out.FoundationModel
		*out = new(string)
		**out = **in
	}
	if in.IdleSessionTTLInSeconds != nil {
		in, out := &in.IdleSessionTTLInSeconds, &out.IdleSessionTTLInSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Instruction != nil {
		in, out := &in.Instruction, &out.Instruction
		*out = new(string)
		**out = **in
	}
	if in.PrepareAgent != nil {
		in, out := &in.PrepareAgent, &out.PrepareAgent
		*out = new(bool)
		**out = **in
	}
	if in.PromptOverrideConfiguration != nil {
		in, out := &in.PromptOverrideConfiguration, &out.PromptOverrideConfiguration
		*out = make([]PromptOverrideConfigurationInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentInitParameters.
func (in *AgentInitParameters) DeepCopy() *AgentInitParameters {
	if in == nil {
		return nil
	}
	out := new(AgentInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentList) DeepCopyInto(out *AgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Agent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentList.
func (in *AgentList) DeepCopy() *AgentList {
	if in == nil {
		return nil
	}
	out := new(AgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentObservation) DeepCopyInto(out *AgentObservation) {
	*out = *in
	if in.AgentArn != nil {
		in, out := &in.AgentArn, &out.AgentArn
		*out = new(string)
		**out = **in
	}
	if in.AgentID != nil {
		in, out := &in.AgentID, &out.AgentID
		*out = new(string)
		**out = **in
	}
	if in.AgentName != nil {
		in, out := &in.AgentName, &out.AgentName
		*out = new(string)
		**out = **in
	}
	if in.AgentResourceRoleArn != nil {
		in, out := &in.AgentResourceRoleArn, &out.AgentResourceRoleArn
		*out = new(string)
		**out = **in
	}
	if in.AgentVersion != nil {
		in, out := &in.AgentVersion, &out.AgentVersion
		*out = new(string)
		**out = **in
	}
	if in.CustomerEncryptionKeyArn != nil {
		in, out := &in.CustomerEncryptionKeyArn, &out.CustomerEncryptionKeyArn
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FoundationModel != nil {
		in, out := &in.FoundationModel, &out.FoundationModel
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IdleSessionTTLInSeconds != nil {
		in, out := &in.IdleSessionTTLInSeconds, &out.IdleSessionTTLInSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Instruction != nil {
		in, out := &in.Instruction, &out.Instruction
		*out = new(string)
		**out = **in
	}
	if in.PrepareAgent != nil {
		in, out := &in.PrepareAgent, &out.PrepareAgent
		*out = new(bool)
		**out = **in
	}
	if in.PromptOverrideConfiguration != nil {
		in, out := &in.PromptOverrideConfiguration, &out.PromptOverrideConfiguration
		*out = make([]PromptOverrideConfigurationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentObservation.
func (in *AgentObservation) DeepCopy() *AgentObservation {
	if in == nil {
		return nil
	}
	out := new(AgentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentParameters) DeepCopyInto(out *AgentParameters) {
	*out = *in
	if in.AgentName != nil {
		in, out := &in.AgentName, &out.AgentName
		*out = new(string)
		**out = **in
	}
	if in.AgentResourceRoleArn != nil {
		in, out := &in.AgentResourceRoleArn, &out.AgentResourceRoleArn
		*out = new(string)
		**out = **in
	}
	if in.AgentResourceRoleArnRef != nil {
		in, out := &in.AgentResourceRoleArnRef, &out.AgentResourceRoleArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AgentResourceRoleArnSelector != nil {
		in, out := &in.AgentResourceRoleArnSelector, &out.AgentResourceRoleArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomerEncryptionKeyArn != nil {
		in, out := &in.CustomerEncryptionKeyArn, &out.CustomerEncryptionKeyArn
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FoundationModel != nil {
		in, out := &in.FoundationModel, &out.FoundationModel
		*out = new(string)
		**out = **in
	}
	if in.IdleSessionTTLInSeconds != nil {
		in, out := &in.IdleSessionTTLInSeconds, &out.IdleSessionTTLInSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Instruction != nil {
		in, out := &in.Instruction, &out.Instruction
		*out = new(string)
		**out = **in
	}
	if in.PrepareAgent != nil {
		in, out := &in.PrepareAgent, &out.PrepareAgent
		*out = new(bool)
		**out = **in
	}
	if in.PromptOverrideConfiguration != nil {
		in, out := &in.PromptOverrideConfiguration, &out.PromptOverrideConfiguration
		*out = make([]PromptOverrideConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentParameters.
func (in *AgentParameters) DeepCopy() *AgentParameters {
	if in == nil {
		return nil
	}
	out := new(AgentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentSpec) DeepCopyInto(out *AgentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentSpec.
func (in *AgentSpec) DeepCopy() *AgentSpec {
	if in == nil {
		return nil
	}
	out := new(AgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentStatus) DeepCopyInto(out *AgentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentStatus.
func (in *AgentStatus) DeepCopy() *AgentStatus {
	if in == nil {
		return nil
	}
	out := new(AgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferenceConfigurationInitParameters) DeepCopyInto(out *InferenceConfigurationInitParameters) {
	*out = *in
	if in.MaxLength != nil {
		in, out := &in.MaxLength, &out.MaxLength
		*out = new(float64)
		**out = **in
	}
	if in.StopSequences != nil {
		in, out := &in.StopSequences, &out.StopSequences
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Temperature != nil {
		in, out := &in.Temperature, &out.Temperature
		*out = new(float64)
		**out = **in
	}
	if in.TopK != nil {
		in, out := &in.TopK, &out.TopK
		*out = new(float64)
		**out = **in
	}
	if in.TopP != nil {
		in, out := &in.TopP, &out.TopP
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferenceConfigurationInitParameters.
func (in *InferenceConfigurationInitParameters) DeepCopy() *InferenceConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(InferenceConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferenceConfigurationObservation) DeepCopyInto(out *InferenceConfigurationObservation) {
	*out = *in
	if in.MaxLength != nil {
		in, out := &in.MaxLength, &out.MaxLength
		*out = new(float64)
		**out = **in
	}
	if in.StopSequences != nil {
		in, out := &in.StopSequences, &out.StopSequences
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Temperature != nil {
		in, out := &in.Temperature, &out.Temperature
		*out = new(float64)
		**out = **in
	}
	if in.TopK != nil {
		in, out := &in.TopK, &out.TopK
		*out = new(float64)
		**out = **in
	}
	if in.TopP != nil {
		in, out := &in.TopP, &out.TopP
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferenceConfigurationObservation.
func (in *InferenceConfigurationObservation) DeepCopy() *InferenceConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(InferenceConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InferenceConfigurationParameters) DeepCopyInto(out *InferenceConfigurationParameters) {
	*out = *in
	if in.MaxLength != nil {
		in, out := &in.MaxLength, &out.MaxLength
		*out = new(float64)
		**out = **in
	}
	if in.StopSequences != nil {
		in, out := &in.StopSequences, &out.StopSequences
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Temperature != nil {
		in, out := &in.Temperature, &out.Temperature
		*out = new(float64)
		**out = **in
	}
	if in.TopK != nil {
		in, out := &in.TopK, &out.TopK
		*out = new(float64)
		**out = **in
	}
	if in.TopP != nil {
		in, out := &in.TopP, &out.TopP
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InferenceConfigurationParameters.
func (in *InferenceConfigurationParameters) DeepCopy() *InferenceConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(InferenceConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromptConfigurationsInitParameters) DeepCopyInto(out *PromptConfigurationsInitParameters) {
	*out = *in
	if in.BasePromptTemplate != nil {
		in, out := &in.BasePromptTemplate, &out.BasePromptTemplate
		*out = new(string)
		**out = **in
	}
	if in.InferenceConfiguration != nil {
		in, out := &in.InferenceConfiguration, &out.InferenceConfiguration
		*out = make([]InferenceConfigurationInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ParserMode != nil {
		in, out := &in.ParserMode, &out.ParserMode
		*out = new(string)
		**out = **in
	}
	if in.PromptCreationMode != nil {
		in, out := &in.PromptCreationMode, &out.PromptCreationMode
		*out = new(string)
		**out = **in
	}
	if in.PromptState != nil {
		in, out := &in.PromptState, &out.PromptState
		*out = new(string)
		**out = **in
	}
	if in.PromptType != nil {
		in, out := &in.PromptType, &out.PromptType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromptConfigurationsInitParameters.
func (in *PromptConfigurationsInitParameters) DeepCopy() *PromptConfigurationsInitParameters {
	if in == nil {
		return nil
	}
	out := new(PromptConfigurationsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromptConfigurationsObservation) DeepCopyInto(out *PromptConfigurationsObservation) {
	*out = *in
	if in.BasePromptTemplate != nil {
		in, out := &in.BasePromptTemplate, &out.BasePromptTemplate
		*out = new(string)
		**out = **in
	}
	if in.InferenceConfiguration != nil {
		in, out := &in.InferenceConfiguration, &out.InferenceConfiguration
		*out = make([]InferenceConfigurationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ParserMode != nil {
		in, out := &in.ParserMode, &out.ParserMode
		*out = new(string)
		**out = **in
	}
	if in.PromptCreationMode != nil {
		in, out := &in.PromptCreationMode, &out.PromptCreationMode
		*out = new(string)
		**out = **in
	}
	if in.PromptState != nil {
		in, out := &in.PromptState, &out.PromptState
		*out = new(string)
		**out = **in
	}
	if in.PromptType != nil {
		in, out := &in.PromptType, &out.PromptType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromptConfigurationsObservation.
func (in *PromptConfigurationsObservation) DeepCopy() *PromptConfigurationsObservation {
	if in == nil {
		return nil
	}
	out := new(PromptConfigurationsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromptConfigurationsParameters) DeepCopyInto(out *PromptConfigurationsParameters) {
	*out = *in
	if in.BasePromptTemplate != nil {
		in, out := &in.BasePromptTemplate, &out.BasePromptTemplate
		*out = new(string)
		**out = **in
	}
	if in.InferenceConfiguration != nil {
		in, out := &in.InferenceConfiguration, &out.InferenceConfiguration
		*out = make([]InferenceConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ParserMode != nil {
		in, out := &in.ParserMode, &out.ParserMode
		*out = new(string)
		**out = **in
	}
	if in.PromptCreationMode != nil {
		in, out := &in.PromptCreationMode, &out.PromptCreationMode
		*out = new(string)
		**out = **in
	}
	if in.PromptState != nil {
		in, out := &in.PromptState, &out.PromptState
		*out = new(string)
		**out = **in
	}
	if in.PromptType != nil {
		in, out := &in.PromptType, &out.PromptType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromptConfigurationsParameters.
func (in *PromptConfigurationsParameters) DeepCopy() *PromptConfigurationsParameters {
	if in == nil {
		return nil
	}
	out := new(PromptConfigurationsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromptOverrideConfigurationInitParameters) DeepCopyInto(out *PromptOverrideConfigurationInitParameters) {
	*out = *in
	if in.OverrideLambda != nil {
		in, out := &in.OverrideLambda, &out.OverrideLambda
		*out = new(string)
		**out = **in
	}
	if in.PromptConfigurations != nil {
		in, out := &in.PromptConfigurations, &out.PromptConfigurations
		*out = make([]PromptConfigurationsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromptOverrideConfigurationInitParameters.
func (in *PromptOverrideConfigurationInitParameters) DeepCopy() *PromptOverrideConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(PromptOverrideConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromptOverrideConfigurationObservation) DeepCopyInto(out *PromptOverrideConfigurationObservation) {
	*out = *in
	if in.OverrideLambda != nil {
		in, out := &in.OverrideLambda, &out.OverrideLambda
		*out = new(string)
		**out = **in
	}
	if in.PromptConfigurations != nil {
		in, out := &in.PromptConfigurations, &out.PromptConfigurations
		*out = make([]PromptConfigurationsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromptOverrideConfigurationObservation.
func (in *PromptOverrideConfigurationObservation) DeepCopy() *PromptOverrideConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(PromptOverrideConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromptOverrideConfigurationParameters) DeepCopyInto(out *PromptOverrideConfigurationParameters) {
	*out = *in
	if in.OverrideLambda != nil {
		in, out := &in.OverrideLambda, &out.OverrideLambda
		*out = new(string)
		**out = **in
	}
	if in.PromptConfigurations != nil {
		in, out := &in.PromptConfigurations, &out.PromptConfigurations
		*out = make([]PromptConfigurationsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromptOverrideConfigurationParameters.
func (in *PromptOverrideConfigurationParameters) DeepCopy() *PromptOverrideConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(PromptOverrideConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}
