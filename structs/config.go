package structs

type Config struct {
	Database struct {
		Addr         string
		User         string
		Passwd       string
		DatabaseName string
	}
	Log struct {
		File          string
		ConsoleOutput bool
		WriteError    bool
		WriteInfo     bool
	}
}
