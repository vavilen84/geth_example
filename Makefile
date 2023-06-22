abi:
	solc --abi contracts/Contract1.sol -o build --overwrite

abigen:
	abigen --abi=./build/Contract1.abi --pkg=abigen --out=abigen/Contract1.go
