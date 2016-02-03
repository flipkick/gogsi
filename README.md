# gogsi [![Build Status](http://img.shields.io/travis/mammothbane/gogsi.svg?style=flat-square)](https://travis-ci.org/mammothbane/gogsi) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/mammothbane/gogsi) [![Go report](https://img.shields.io/badge/go_report-A+-brightgreen.svg?style=flat-square)](http://goreportcard.com/report/mammothbane/gogsi) 

`go get github.com/mammothbane/gogsi`

The Dota 2 client is capable of streaming JSON-formatted updates about the game to an arbitrary HTTP endpoint. This library listens to those updates and converts them to a Go representation that you can manipulate as convenient.

## Usage
A barebones example might look like this:

```go
package main

import "github.com/mammothbane/gogsi"

func main() {
	gogsi.Listen("localhost:3000", func(s *gogsi.State) error {
		if s.Added.Map.NightstalkerNight {
			/* adjust your phillips hue light, play a scary noise, update statistics, log something, etc. */
		}
		return nil
	})
}
```

The notion of change in the Dota 2 GSI is represented by the `State.Added` and `State.Previous` fields, so, as might be expected, this snippet responds to the value of `Map.NightstalkerNight` changing to `true`. Further detail about `Added` and `Previous` can be found in the [docs](https://godoc.org/github.com/mammothbane/gogsi#State).

## Config
A sample configuration is provided below, set up to send all game updates to `localhost:3000`. Save it (along with any desired changes) to `game/dota/cfg/gamestate_integration/gamestate_integration_[integration_name].cfg` in order to register your integration with the Dota 2 client.

```cfg
"Dota 2 Sample Integration"
{
	"uri"		"http://localhost:3000"
	"timeout"	"5.0"
	"buffer"  	"0.1"
	"throttle"	"0.1"
	"heartbeat"	"30.0"
	"data"
	{
		"provider"		"1"
		"map"			"1"
		"player"		"1"
		"hero"			"1"
		"abilities"		"1"
		"items"			"1"
		"allplayers"	"1"
	}	
}
```

The first line (the name of the integration) should be unique.

`uri` is the URI of your server (usually on `localhost`).

`buffer` is the amount of time (in seconds) for which the client should wait once it has a delta to send, collecting together all deltas in the meantime in order to reduce traffic. A value of `"0"` causes every delta to be sent as soon as it is available.

`throttle` is the minimum amount of time (in seconds) the client waits before sending a new delta after receiving a 2XX response from the server.

`heartbeat` specifies the maximum amount of time the client waits before sending an update, even if no events have occurred.  

`data` options are more or less self-explanatory (`"1"` indicates that the game client should send updates about the given item, `"0"` indicates that it shouldn't). 

Additionally, an `auth` block can be specified as follows:

```cfg
"Dota 2 Sample Integration"
{
	// ...
	"auth"
	{
		"token" "an_auth_token"
	}
}
```

This token will be transmitted with every request as `state.Auth.Token`, so it's easiest to use as a shared key and/or unique ID per client for a remote `gogsi` instance.

These options are also covered in more detail at the [CS: GO GSI page](https://developer.valvesoftware.com/wiki/Counter-Strike:_Global_Offensive_Game_State_Integration).

## Notes and caveats
GSI does not begin sending updates until the first time the player loads into a game. Normally (and according to Valve's spec and various observations), the heartbeat is intended to be respected while in the client UI (i.e. not loaded into a game). This does not occur until **after** a game is loaded for the first time. I have also come across instances in testing where the client did not send GSI updates until the 0:00 horn in the first game.

On another note, **this is not a scripting or hacking tool.** Anyone looking for an easy way to break into the game should look somewhere else. The GSI is provided by Valve, and it does not allow any input. No information about other heroes or creeps is available, and position data is similarly not accessible.

### Credits
This package was inspired by [antonpup's C# implementation](https://github.com/antonpup/Dota2GSI).
