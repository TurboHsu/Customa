package structs

// Structs for messages communicating with different functions.

type ErrMsg struct {
	IsErr bool
	Msg   string
}

type RetMsg struct {
	Msg string
	ErrMsg
}
