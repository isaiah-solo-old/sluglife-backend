testAll: testSimpleDataStore

testSimpleDataStore: simple_datastore.go simple_datastore_test.go datastore.go
	go test simple_datastore.go simple_datastore_test.go datastore.go
