package migrations

// Генерация bindata для миграций

// install go-bindata
// go get -u github.com/kevinburke/go-bindata/...

//go:generate go-bindata -ignore=.*\.go -pkg migrations .
