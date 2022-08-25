package interfaces

type Repository interface {
	login() (error, *Session)
}

type GCR interface {
	createRepository() (error, *Repository)
}

type Session interface {
	getNode() (error, Node)
	getNodeAt(string) (error, Node)
	getRepository() Repository
	getRootNode() Node
	getPropertyAt(string)
	getItemAt(string) Item
	nodeExistsAt(string) bool
	propertyExistsAt(string) bool
	hasPendingChanges(string) bool
	save() error
	refresh() error
	logout()
}
