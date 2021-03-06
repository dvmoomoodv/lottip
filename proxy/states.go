package proxy

// Cmd represents MySQL command to be executed.
type Cmd struct {
	ConnId int
	CmdId  int
	Query  string
}

// CmdResult represents MySQL command execution result.
type CmdResult struct {
	ConnId   int
	CmdId    int
	Result   byte
	Error    string
	Duration string
}

// ConnState represents tcp connection state.
type ConnState struct {
	ConnId int
	State  byte
}
