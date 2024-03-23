package main

import (
	"log"
	"os"

	"github.com/andygrunwald/vdf"
	lnk "github.com/parsiya/golnk"
)

// Looks up where VTOL VR is installed by parsing libraryfolders.vdf and finding out where steam has it stored based on that data.
func getVTOLDir() []string {

	file, err := os.Open(getSteamDir() + "\\steamapps\\libraryfolders.vdf")
	if err != nil {
		log.Panic(err)
	}

	// Instantiates a new steam VDF parser and reads the VDF file.
	vdfP := vdf.NewParser(file)
	vdfInf, err := vdfP.Parse()

	if err != nil {
		log.Panic(err)
	}

	var libs []vdflib

	// Saves the VDF file contents to a golang struct.
	for _, x := range vdfInf["libraryfolders"].(map[string]interface{}) {
		var newLib vdflib
		for f, y := range x.(map[string]interface{}) {

			switch f {
			case "path":
				newLib.Path = y.(string)
			case "contentid":
				newLib.ContentID = y.(string)
			case "label":
				newLib.Label = y.(string)
			case "totalsize":
				newLib.TotalSize = y.(string)
			case "update_clean_bytes_tally":
				newLib.UpdateCleanBytesTally = y.(string)
			case "time_last_update_corruption":
				newLib.TimeLastUpdateCorruption = y.(string)
			case "apps":

				for z, u := range y.(map[string]interface{}) {
					var newapp libapp
					newapp.AppID = z
					newapp.BuildID = u.(string)
					newLib.Apps = append(newLib.Apps, newapp)
				}

			}

		}
		libs = append(libs, newLib)
	}

	var possiblePaths []string

	// Saves all possible VTOL VR paths to a variable
	for _, x := range libs {
		for _, y := range x.Apps {
			if y.AppID == "667970" {
				currentPath := x.Path + "\\steamapps\\common\\VTOL VR\\"
				possiblePaths = append(possiblePaths, currentPath)
			}
		}
	}

	return possiblePaths
}

// Finds where steam is installed based on the steam.lnk shortcut in the Windows Start Menu.
func getSteamDir() string {

	home, err := os.UserHomeDir()
	if err != nil {
		log.Panic(err)
	}
	lnk, err := lnk.File(home + "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Steam\\Steam.lnk")
	if err != nil {
		log.Panic(err)
	}

	return lnk.StringData.WorkingDir

}
