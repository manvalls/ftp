package ftp

// Auth is the interface that wraps the CheckPasswd method
type Auth interface {
	CheckPasswd(string, string) (bool, error)
}
