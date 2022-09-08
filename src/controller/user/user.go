package user
import(
	user "echo/src/orm/user"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func AddUser(c echo.Context) error {
	name := c.FormValue("name")
	id := user.Create(name)
	return c.String(http.StatusOK, strconv.Itoa(id))
}

func DelUser(c echo.Context) error {
	id := c.FormValue("id")
	ok := user.Del(parseint(id))
	return c.String(http.StatusOK, boolToString(ok))
}

func UpdateUser(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	ok := user.Update(parseint(id),name)
	return c.String(http.StatusOK, boolToString(ok))
}

func GetUser(c echo.Context) error {
	id := c.QueryParam("id")
	data := user.Get(parseint(id))
	return c.String(http.StatusOK, data)
}

func parseint(str string) int{
	data,_ := strconv.ParseInt(str,10,32)
	return int(data)
}

func boolToString(b bool) string{
	if b {
		return "true"
	}else{
		return "false"
	}
}

func Init(server *echo.Echo)  {
	server.POST("/user",AddUser)
	server.DELETE("/user",DelUser)
	server.PUT("/user",UpdateUser)
	server.GET("/user",GetUser)
}


