package main

type HTTPServer struct {
	config *Config
}

func (*HTTPServer) ListenAndServe(hostspec string, unknown interface{}) {
}

func MakeHTTPServer(config *Config) (*HTTPServer, error) {
	var server = HTTPServer{config: config}
	return &server, nil
}
