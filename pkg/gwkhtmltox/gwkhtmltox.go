package gwkhtmltox

var Status bool = false

func init() {
	Destroy()
	Status = ImageInit() == nil
}

func IsAvailable() bool {
	return Status
}
func Destroy() {
	ImageDestroy()
}
