package providers

// MySQLProvider implements a BaseProvider struct
// for MySQL backups
type MySQLProvider struct {
	*BaseProvider
}

// GetName returns the provider name
func (*MySQLProvider) GetName() string {
	return "MySQL"
}

// GetPrepareCommandToVolume returns the command to be executed before backup in order to store the backup into the current volume
func (p *MySQLProvider) GetPrepareCommandToVolume(volDestination string) []string {
	return []string{
		"sh",
		"-c",
		"mkdir -p " + volDestination + "/backups && mysqldump --all-databases --extended-insert --password=$MYSQL_ROOT_PASSWORD > " + volDestination + "/backups/dump.sql",
	}
}

// GetPrepareCommandToPipe returns the command to be executed before backup in order to send the backup through a pipe
func (p *MySQLProvider) GetPrepareCommandToPipe() []string {
	return []string{
		"sh",
		"-c",
		"mysqldump --all-databases --extended-insert --password=$MYSQL_ROOT_PASSWORD",
	}
}

// GetBackupDir returns the backup directory used by the provider
func (p *MySQLProvider) GetBackupDir() string {
	return "backups"
}

// SetVolumeBackupDir sets the backup dir for the volume
func (p *MySQLProvider) SetVolumeBackupDir() {
	p.vol.BackupDir = p.GetBackupDir()
}
