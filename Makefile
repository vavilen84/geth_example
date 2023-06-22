abi:
	solc --abi contracts/$(name).sol -o gen/abi --overwrite

abi-gen:
	@mkdir -p gen/go/$(pkg)
	abigen --abi=./gen/abi/$(name).abi --pkg=$(pkg) --out=gen/go/$(pkg)/$(pkg).go
