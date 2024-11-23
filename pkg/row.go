package row

import (
	"fmt"
	"io/fs"
	"regexp"
	"strings"
)

type FileRow struct {
	Name string
	Type string
	Icon string
	DirFlag string
	Perms string
}

func NewFileRow(file fs.DirEntry) *FileRow {
	t := getTypeName(file)
	d := "."

	if file.IsDir() {
		d = "d"
	}

	return &FileRow{
		Name: file.Name(),
		Type: t,
		Icon: getNerdIcon(t),
		DirFlag: d,
		Perms: file.Type().Perm().String(),
	}
}

func getTypeName(file fs.DirEntry) string {
	if file.IsDir() {
		return "/dir"
	}

	special, name := isSpecialFileName(file.Name())
	if special {
		return name
	}

	ext := getExtension(file.Name())
	if ext != "" {
		return standardizeExtension(ext)
	}

	return "/unknown"
}

func isSpecialFileName(name string) (bool, string) {
	matched, _ := regexp.MatchString(`(?i)^readme(\..*)?$`, name)
	if matched {
		return true, "/readme"
	}

	matched, _ = regexp.MatchString(`(?i)^license(\..*)?$`, name)
	if matched {
		return true, "/license"
	}

	if name == "go.mod" || name == "go.sum" {
		return true, "/go"
	}

	return false, ""
}

func getExtension(filename string) string {
	if strings.Contains(filename, ".") {
		parts := strings.Split(filename, ".")
		return parts[len(parts)-1]
	}

	return ""
}

func standardizeExtension(ext string) string {
	ext = strings.ToLower(ext)

	if ext == "jpeg" || ext == "jpg" || ext == "png" || ext == "gif" || ext == "bmp" {
		return "/image"
	}

	return ext
}

func getNerdIcon(typename string) string {
	switch typename {
	case "/dir":
		return "📁"
	case "/readme":
		return "📖"
	case "/license":
		return "📜"
	case "/go":
		return "📦"
	case "/image":
		return "🖼"
	default:
		return "📄"
	}
}

func (r *FileRow) String() string {
	return fmt.Sprintf("%s%s %s %s", r.DirFlag, r.Perms, r.Icon, r.Name)
}