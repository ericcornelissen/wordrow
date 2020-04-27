package fs


// The File type represents a file's metadata and content.
type File struct {
  // The file's content.
  Content string

  // The file's extension.
  Ext string

  // The file's absolute path.
  Path string
}
