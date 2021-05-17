package core

type Repository interface {
	Add(Recipe) error
	Update(Recipe) error
	Get(string) (*Recipe, error)
	List(RecipeType) ([]*Recipe, error)
	Delete(string) error
}
