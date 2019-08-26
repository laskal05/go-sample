package handler

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type data struct {
	Data string `json:"data"`
}

// Post is ...
func Post(c echo.Context) error {
	post := new(data)
	if err := c.Bind(post); err != nil {
		return err
	}
	data, err := base64.StdEncoding.DecodeString(post.Data)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	f, _ := os.Create("/tmp/hoge.py")
	_, _ := f.Write(data)
	os.Exec("docker exec -it 2b7b4174a62f python /tmp/hoge.py")
	return c.String(http.StatusOK, "Hello World")
}
