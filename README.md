[![Go Reference](https://pkg.go.dev/badge/github.com/tommzn/recipeboard-core.svg)](https://pkg.go.dev/github.com/tommzn/recipeboard-core)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tommzn/recipeboard-core)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/tommzn/recipeboard-core)

# Recipe Board Core
Core components of the recipe board project.

Includes the receipe manager to manager life circle of receips, together with a persistence layer
and a publisher to send notifications about actions performed for recipe.

## Sun Modules
### Mock
[![Go Reference](https://pkg.go.dev/badge/github.com/tommzn/recipeboard-core.svg)](https://pkg.go.dev/github.com/tommzn/recipeboard-core/mock)
Provices a persitsence and a publisher mock. Both implementes inferface given by core module can 
i.e. be used for testing.

### Model
[![Go Reference](https://pkg.go.dev/badge/github.com/tommzn/recipeboard-core.svg)](https://pkg.go.dev/github.com/tommzn/recipeboard-core/model)
Contains the core model fo recipes and messaging and interfaces for persistence layer and publishers.
