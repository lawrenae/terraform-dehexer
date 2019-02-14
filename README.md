# terraform-dehexer
converts a terraform "hex" output from random.id back into it's native value, for cases of emergency terraform state surgery

# usage
```sh
go build -o dehex dehex.go && ./dehex ec28
```

# example output
```json
"attributes": {
		"b64": "7Cg",
		"b64_std": "7Cg=",
		"b64_url": "7Cg",
		"byte_length": "2",
		"dec": "60456",
		"hex": "ec28",
		"id": "7Cg"
	},
```

# notes
- totally quick and dirty :)