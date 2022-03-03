package Helpers

import (
	"os"
)

func GetGoRunPath() string {
	appPath, _ := os.Getwd()
	return appPath
}