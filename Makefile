testAll: testSimpleDataStore

testSimpleDataStore: simple_dining_store.go dining_store_test.go dining.go dining_store.go
	go test simple_dining_store.go dining_store_test.go dining_store.go dining.go

testEvents: event.go  simple_event_store.go event_store.go event_store_test.go
	go test $?
