package providers

// PostgreSQLProvider implements a BaseProvider struct
// for PostgreSQL backups
type PostgreSQLProvider struct {
	*BaseProvider
}

// GetName returns the provider name
func (p *PostgreSQLProvider) GetName() string {
	return "PostgreSQL"
}

// GetPrepareCommandToVolume returns the command to be executed before backup
func (p *PostgreSQLProvider) GetPrepareCommandToVolume(volDestination string) []string {
	return []string{
		"sh",
		"-c",
		"mkdir -p " + volDestination + "/backups && pg_dumpall --clean -Upostgres > " + volDestination + "/backups/all.sql",
	}
}

// GetPrepareCommandToPipe returns the command to be executed before backup
func (p *PostgreSQLProvider) GetPrepareCommandToPipe() []string {
	return []string{
		"sh",
		"-c",
		"pg_dumpall --clean -Upostgres",
	}
}

// GetBackupDir returns the backup directory used by the provider
func (p *PostgreSQLProvider) GetBackupDir() string {
	return "backups"
}

// SetVolumeBackupDir sets the backup dir for the volume
func (p *PostgreSQLProvider) SetVolumeBackupDir() {
	p.vol.BackupDir = p.GetBackupDir()
}
