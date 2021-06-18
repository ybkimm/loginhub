package main

import (
	"net/http"

	"github.com/ybkimm/loginhub/internal/assets"
)

const (
	imagePrefix  = "/assets/images"
	imageDir     = "assets/images"
	scriptPrefix = "/assets/scripts"
	scriptDir    = "assets/scripts"
	stylePrefix  = "/assets/styles"
	styleDir     = "build/styles"
)

var (
	imageHandler  = http.StripPrefix(imagePrefix, assets.NewHandler(&realfs{imageDir}, logger))
	scriptHandler = http.StripPrefix(scriptPrefix, assets.NewHandler(&realfs{scriptDir}, logger))
	styleHandler  = http.StripPrefix(stylePrefix, assets.NewHandler(&realfs{styleDir}, logger))
)
