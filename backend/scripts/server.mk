run_app:
	go run ./core/main.go -migrate=${MIGRATE_FLAG} -redis=${REDIS_MODE}

run_tree:
	go run ./scripts/src/tree.go
