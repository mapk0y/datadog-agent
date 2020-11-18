// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2017-2020 Datadog, Inc.

package listeners

import (
	"github.com/DataDog/datadog-agent/pkg/autodiscovery/integration"
	"github.com/DataDog/datadog-agent/pkg/config"
	"github.com/DataDog/datadog-agent/pkg/util/containers"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

// EnvironmentListener implements a ServiceListener based on current environment
type EnvironmentListener struct {
	newService chan<- Service
	delService chan<- Service
}

// EnvironmentService represents services generated from EnvironmentListener
type EnvironmentService struct {
	adIdentifier string
}

// Make sure ECSService implements the Service interface
var _ Service = &EnvironmentService{}

func init() {
	Register("environment", NewEnvironmentListener)
}

// NewEnvironmentListener creates an EnvironmentListener
func NewEnvironmentListener() (ServiceListener, error) {
	return &EnvironmentListener{}, nil
}

// Listen checks
func (l *EnvironmentListener) Listen(newSvc chan<- Service, delSvc chan<- Service) {
	l.newService = newSvc
	l.delService = delSvc

	// ATM we consider environment as a fixed space
	// It may change in the future
	go l.createServices()
}

// Stop has nothing to do in this case
func (l *EnvironmentListener) Stop() {
}

func (l *EnvironmentListener) createServices() {
	if config.IsFeaturePresent(config.Docker) {
		log.Info("Listener created Docker service from environment")
		l.newService <- &EnvironmentService{adIdentifier: "_docker"}
	}

	if config.IsFeaturePresent(config.Kubernetes) {
		log.Info("Listener created Kubelet service from environment")
		l.newService <- &EnvironmentService{adIdentifier: "_kubelet"}
	}

	if config.IsFeaturePresent(config.ECSFargate) {
		log.Info("Listener created ECS Fargate service from environment")
		l.newService <- &EnvironmentService{adIdentifier: "_ecs_fargate"}
	}

	if config.IsFeaturePresent(config.EKSFargate) {
		log.Info("Listener created EKS Fargate service from environment")
		l.newService <- &EnvironmentService{adIdentifier: "_eks_fargate"}
	}

	if config.IsFeaturePresent(config.Cri) {
		log.Info("Listener created CRI service from environment")
		l.newService <- &EnvironmentService{adIdentifier: "_cri"}
	}

	if config.IsFeaturePresent(config.Containerd) {
		log.Info("Listener created Containerd service from environment")
		l.newService <- &EnvironmentService{adIdentifier: "_containerd"}
	}
}

// GetEntity returns the unique entity name linked to that service
func (s *EnvironmentService) GetEntity() string {
	return s.adIdentifier
}

// GetTaggerEntity returns the tagger entity
func (s *EnvironmentService) GetTaggerEntity() string {
	return ""
}

// GetADIdentifiers return the single AD identifier for an environment service
func (s *EnvironmentService) GetADIdentifiers() ([]string, error) {
	return []string{s.adIdentifier}, nil
}

// GetHosts is not supported
func (s *EnvironmentService) GetHosts() (map[string]string, error) {
	return nil, ErrNotSupported
}

// GetPorts returns nil and an error because port is not supported in Fargate-based ECS
func (s *EnvironmentService) GetPorts() ([]ContainerPort, error) {
	return nil, ErrNotSupported
}

// GetTags retrieves a container's tags
func (s *EnvironmentService) GetTags() ([]string, error) {
	return nil, nil
}

// GetPid inspect the container and return its pid
// TODO: not supported as pid is not in the metadata api
func (s *EnvironmentService) GetPid() (int, error) {
	return -1, ErrNotSupported
}

// GetHostname returns nil and an error because port is not supported in Fargate-based ECS
func (s *EnvironmentService) GetHostname() (string, error) {
	return "", ErrNotSupported
}

// GetCreationTime is always before for environment service
func (s *EnvironmentService) GetCreationTime() integration.CreationTime {
	return integration.Before
}

// IsReady is always true
func (s *EnvironmentService) IsReady() bool {
	return true
}

// GetCheckNames is not supported
func (s *EnvironmentService) GetCheckNames() []string {
	return nil
}

// HasFilter is not supported
func (s *EnvironmentService) HasFilter(filter containers.FilterType) bool {
	return false
}

// GetExtraConfig is not supported
func (s *EnvironmentService) GetExtraConfig(key []byte) ([]byte, error) {
	return []byte{}, ErrNotSupported
}
