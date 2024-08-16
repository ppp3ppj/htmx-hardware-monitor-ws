package dashboard

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/htmx-hardware-monitor-ws/internal/hardware"
	"github.com/ppp3ppj/htmx-hardware-monitor-ws/template"
	views_pages "github.com/ppp3ppj/htmx-hardware-monitor-ws/views/pages/index"
	views_variables "github.com/ppp3ppj/htmx-hardware-monitor-ws/views/variables"
)

type DashboardFrontend struct {}

func NewDashboardFrontend(
    g *echo.Group,
) {
    fe := &DashboardFrontend{}

    g.GET("", fe.Index)
}

func (fe *DashboardFrontend) Index(c echo.Context) error {
    bodyOpts := views_variables.BodyOpts{
        ExtraHeaders: nil,
        Component: nil,
    }
    fmt.Println("******* Test *********")
    fmt.Println(hardware.GetSystemSection())
    fmt.Println("******* End *********")

    index := views_pages.Index(bodyOpts)
    return template.AssertRender(c, http.StatusOK, index)
}
