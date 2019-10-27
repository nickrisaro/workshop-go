package tweet

// Tweet estructura para guardar la información de un tweet
type Tweet struct {
	Username string
	Text     string
}

// New construye un nuevo tweet
func New(username, text string) *Tweet {
	return &Tweet{username, text}
}
