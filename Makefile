testAll: testSimpleDataStore

testSimpleDataStore: simple_dining_store.go dining_store_test.go dining.go dining_store.go
	go test simple_dining_store.go dining_store_test.go dining_store.go dining.go
