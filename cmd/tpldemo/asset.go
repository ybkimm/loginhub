package main

import "github.com/ybkimm/loginhub/internal/assets"

const assetdir = "internal/assets"

var assetHandler = assets.NewHandler(&realfs{assetdir})
