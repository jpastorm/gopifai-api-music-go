package controllers

import (
	"fmt"
	"github.com/labstack/echo"
	"os"
	"wopifai/src/apihelpers"
	"wopifai/src/resources"
	"strconv"
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
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return err
	}

	c.Response().Header().Set("Content-type","audio/mpeg")

	c.Response().Header().Set("Content-length", strconv.FormatInt(fi.Size(),10))

	return c.Stream(200,"audio/mpeg", f)
}