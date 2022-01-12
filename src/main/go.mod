module main

go 1.17

require "go_learn" v1.0.0
replace "go_learn" => "../go_learn" 

require "usual" v1.0.0
replace "usual" => "../usual"

// require "github.com/gorilla/mux" v1.8.0
// replace "github.com/gorilla/mux" => "../../pkg/mod/github.com/gorilla/mux@v1.8.0"

// require "github.com/gorilla/websocket" v1.4.2
// replace "github.com/gorilla/websocket" => "../../pkg/mod/github.com/gorilla/websocket@v1.4.2"
