package main

// main.go
import (
	"fmt"

	"github.com/codyveladev/day-five/models"
	"github.com/fatih/color"
)

func main() {
	b := models.Book{Title: "Go", Author: "Cody Vela"}
	u := models.User{Username: "c_v204", Email: "c_v203@txst.edu", FirstName: "Cody", LastName: "Vela"}
	fmt.Println(b, u)
	color.Green("SUCCESS: USER CREATED")
	color.HiBlue("INFO: ISSUE WAS MUTED")
}
