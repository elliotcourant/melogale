.PHONY: generate

generate:
	cd pkg/ast && make generate
	cd pkg/base && make generate