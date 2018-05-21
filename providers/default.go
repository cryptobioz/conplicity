package providers

import (
	log "github.com/Sirupsen/logrus"
)

// DefaultProvider implements a BaseProvider struct
// for simple filesystem backups
type DefaultProvider struct {
	*BaseProvider
}

// GetName returns the provider name
func (*DefaultProvider) GetName() string {
	return "Default"
}

// PrepareBackup sets up the data before backup
func (p *DefaultProvider) PrepareBackup() error {
	log.WithFields(log.Fields{
		"provider": p.GetName(),
	}).Debug("Provider does not implement a prepare method")
	return nil
}

// GetPrepareCommandToVolume returns the command to be executed before backup
func (p *DefaultProvider) GetPrepareCommandToVolume(volDestination string) []string {
	return nil
}

// GetPrepareCommandToPipe returns the command to be executed before backup
func (p *DefaultProvider) GetPrepareCommandToPipe() []string {
	return nil
}
