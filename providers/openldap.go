package providers

// OpenLDAPProvider implements a BaseProvider struct
// for OpenLDAP backups
type OpenLDAPProvider struct {
	*BaseProvider
}

// GetName returns the provider name
func (p *OpenLDAPProvider) GetName() string {
	return "OpenLDAP"
}

// GetPrepareCommandToVolume returns the command to be executed before backup
func (p *OpenLDAPProvider) GetPrepareCommandToVolume(volDestination string) []string {
	return []string{
		"sh",
		"-c",
		"mkdir -p " + volDestination + "/backups && slapcat > " + volDestination + "/backups/all.ldif",
	}
}

// GetPrepareCommandToPipe returns the command to be executed before backup
func (p *OpenLDAPProvider) GetPrepareCommandToPipe() []string {
	return []string{
		"sh",
		"-c",
		"slapcat",
	}
}

// GetBackupDir returns the backup directory used by the provider
func (p *OpenLDAPProvider) GetBackupDir() string {
	return "backups"
}

// SetVolumeBackupDir sets the backup dir for the volume
func (p *OpenLDAPProvider) SetVolumeBackupDir() {
	p.vol.BackupDir = p.GetBackupDir()
}
