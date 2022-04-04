package controllers

import (
	"log"

	"github.com/astaxie/beego"
)

type FileUploadController struct {
	beego.Controller
}

func (c *FileUploadController) Get() {
	c.TplName = "fileupload.html"
}

func (c *FileUploadController) Post() {
	f, h, err := c.GetFile("uploadfilename")
	if err != nil {
		log.Fatal("getfile err", err)
	}
	defer f.Close()
	path := "upload/" + h.Filename
	c.SaveToFile("uploadfilename", path)
	c.TplName = "success.html"
}
