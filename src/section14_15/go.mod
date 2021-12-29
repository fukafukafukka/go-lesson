module section14_15

go 1.17

replace config => ./config

replace utils => ./utils

replace bitflyer => ./bitflyer

require (
	bitflyer v0.0.0-00010101000000-000000000000 // indirect
	config v0.0.0-00010101000000-000000000000 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	utils v0.0.0-00010101000000-000000000000 // indirect
)
