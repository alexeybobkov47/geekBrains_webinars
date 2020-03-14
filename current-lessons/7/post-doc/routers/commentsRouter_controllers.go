package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers:ListController"] = append(beego.GlobalControllerRouter["not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers:ListController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers:ListController"] = append(beego.GlobalControllerRouter["not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers:ListController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers:ListController"] = append(beego.GlobalControllerRouter["not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers:ListController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers:ListController"] = append(beego.GlobalControllerRouter["not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers:ListController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers:StatisticController"] = append(beego.GlobalControllerRouter["not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers:StatisticController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers:TasksController"] = append(beego.GlobalControllerRouter["not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers:TasksController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers:TasksController"] = append(beego.GlobalControllerRouter["not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers:TasksController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers:TasksController"] = append(beego.GlobalControllerRouter["not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers:TasksController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers:TasksController"] = append(beego.GlobalControllerRouter["not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/controllers:TasksController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
