package cmd

type Module struct {
	Config Config
	Server *Server
}

func NewModule(config Config) *Module {
	srv := NewServer(config)

	return &Module{
		Config:             config,
		Server:             srv,
	}
}
