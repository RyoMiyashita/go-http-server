package extension

import(
	"regexp"
	"path/filepath"
)

func GetFileNameFromPath(path string) string {
	r, _ := regexp.Compile(`([^/]+?)?$`)
	return string(r.Find([]byte(path)))
}

func GetContentType(fileName string) string {
	switch string(filepath.Ext(fileName)) {
		case ".html": return "text/html"
		case ".txt": return "text/plain"
		case ".css": return "text/css"
		case ".js": return "text/javascript"
		case ".json": return "application/json"
		case ".jpeg": return "image/jpeg"
		case ".jpg": return "image/jpeg"
		case ".png": return "image/png"
	default: return "application/force-download"
	}
}
