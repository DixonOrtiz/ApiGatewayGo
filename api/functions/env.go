package functions

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

//GetEnv function
//Function that return indicated enviorement variable
func GetEnv(envInput string) string {
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
		return ""
	}

	envOutput := os.Getenv(envInput)

	return envOutput
}
