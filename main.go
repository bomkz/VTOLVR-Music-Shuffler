package main

import (
	"fmt"
)

var Version = "dev"

func main() {

	fmt.Print(
		`Music Scrambler` + Version + `
		
This shit will "shuffle" your music every time you die.
It will do this by renaming everything in your music folder random characters to make them appear in different order.
So please, make sure you only have MP3's, and either back them up in their original state, or if you don't care, leave'em as is.
			
------------------------------------------------------------------------------------------------------
Licensed under MPLv2
------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------
! ! Music Scrambler is starting up, please make sure you are currently not spawned in an aircraft if VTOL VR is already running ! !
------------------------------------------------------------------------------------------------------


`)

	go readLog()
	fmt.Println("Music Scrambler is now listening to log events.")

	wait := make(chan bool)

	<-wait
}
