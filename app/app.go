package app

import (
	"fmt"
	"net/http"

	controllers "../controllers"
)

func StartApp() {
	http.HandleFunc("/user", controllers.GetUser)

	if err := http.ListenAndServe(":1234", nil); err != nil {
		fmt.Println(err)
	}

}
