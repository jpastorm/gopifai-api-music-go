package controllers

import (
	"fmt"
	"github.com/labstack/echo"
	"wopifai/src/apihelpers"
	"wopifai/src/resources"
)

func GetDir(c echo.Context) error{
	var data []resources.Explorer
	var response apihelpers.ResponseLibraries
	if	c.QueryParam("path")=="" {
		return c.JSON(404, "Query incorrecto")
	}
	path := c.QueryParam("path")
	data = resources.List_Dir(path)
	fmt.Println("Respuesta")
	fmt.Println(data)
	response = apihelpers.ResponseLibraries{Status: 200,Message: "Con exito",Data: data}
	return c.JSON(200, response)
}

func GetCover(c echo.Context) error {

	if	c.QueryParam("path")=="" {
		return c.JSON(404, "Query incorrecto")
	}
	path := c.QueryParam("path")

	response,mime := resources.Get_Cover(path)
	return c.Blob(200,mime, response)
}

func GetTrack(c echo.Context) error {
	if	c.QueryParam("path")=="" {
		return c.JSON(404, "Query incorrecto")
	}
	path := c.QueryParam("path")

	response := resources.Get_Track(path)
	return c.JSON(200,response)
}