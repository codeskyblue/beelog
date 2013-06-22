/*
	Example of beelog
*/
package main

import "beelog"

func main() {
	beelog.SetLevel(beelog.LevelWarning)
	beelog.Trace("I can see you")
	beelog.Debug("The air is fresh")
	beelog.Info("What a nice day.")
	beelog.Warn("It's raining outside")
	beelog.Error("Taifeng is comming")
	beelog.Critical("The end of world is comming")
}
