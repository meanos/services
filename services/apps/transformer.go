package apps

import (
	"fmt"
	"strconv"
)

func (a *Application) toInfo() ApplicationInfo {
	return ApplicationInfo{
		Name:              a.Name,
		DeveloperName:     a.DeveloperName,
		DeveloperID:       a.DeveloperID,
		Description:       a.Description,
		DeveloperContacts: a.DeveloperContacts,
		Price:             fmt.Sprintf("%f", a.Price),
		IconURL:           a.IconURL,
		Rating:            fmt.Sprintf("%f", a.Rating),
		DownloadsCounter:  strconv.Itoa(a.DownloadsCounter),
		AppID:             a.AppID,
		ScreenshotURLs:    a.ScreenshotURLs,
		IsPaid:            a.IsPaid,
		isSnap:            false,
		snapName:          "",
	}
}
