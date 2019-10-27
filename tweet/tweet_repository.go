package tweet

// Repository almacena tweets
type Repository struct {
	tweets []*Tweet
}

// NewRepository construye un nuevo repositorio
func NewRepository() *Repository {
	return &Repository{}
}

// GetAllTweets obtiene todos los tweets del repositorio
func (r *Repository) GetAllTweets() []*Tweet {
	return r.tweets
}

// Save almacena el tweet en el repositorio
func (r *Repository) Save(tweet *Tweet) {
	// TODO validar los datos del tweet, por ejemplo la longitud del texto
	// o que el usuario esté en alguna lista de usuario
	// construir un error con errors.New("El texto excede el límite de caracteres")
	r.tweets = append(r.tweets, tweet)
}

// TODO Implementar una búsqueda de tweets
// func (r *Repository) Search(patter string) []*Tweet {}

// TODO Agregar un método para recuperar todos los tweets de un usuario
// Recomiendo que lo hagas usando un mapa con clave nombre de usuario y valor los tweets de ese usuario
// para practicar eso hint: tweetsByUser map[string][]*Tweet
// func (r *Repository) GetTweetsByUser(user string) []*Tweet {}
