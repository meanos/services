package apps

type Application struct {
	Name              string
	DeveloperName     string
	DeveloperID       string
	Description       string
	DeveloperContacts string
	Price             float64
	IconURL           string
	Rating            float64
	DownloadsCounter  int
	AppID             string
	ScreenshotURLs    []string
	IsPaid            bool
}

type ApplicationInfo struct {
	Name              string
	DeveloperName     string
	DeveloperID       string
	Description       string
	DeveloperContacts string
	Price             string
	IconURL           string
	Rating            string
	DownloadsCounter  string
	AppID             string
	ScreenshotURLs    []string
	IsPaid            bool
	isSnap            bool
	snapName          string
}
