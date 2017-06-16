package singleton

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1497618625, 0),
		Content:     string("<!DOCTYPE html>\n<html>\n<head>\n    <meta http-equiv=\"content-type\" content=\"text/html;charset=utf-8\">\n    <title>hello world</title>\n</head>\n<body>\n    <h1>hello world</h1>\n</body>\n</html>"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1497610342, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "index.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`templates`, &embedded.EmbeddedBox{
		Name: `templates`,
		Time: time.Unix(1497610342, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"index.html": file2,
		},
	})
}
