//+build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

// func InitializeEvent() (string, error) {
// 	wire.Build(NewVariable, NewSecrets, NewGormDB, NewPasswordHasher)
// 	return "", nil
// }
func InitializeUserRoute() (UserRoute, error) {
	wire.Build(NewVariable, NewSecrets, NewGormDB, NewPasswordHasher)
	
}
