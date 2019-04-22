package datatypes

// User is the defintion of that data that a user should contain
type User struct {
	Name        string
	Email       string
	Picture     string
	Id          string
	AuthId      string
	Groups      []Group
	Permissions []string
}
