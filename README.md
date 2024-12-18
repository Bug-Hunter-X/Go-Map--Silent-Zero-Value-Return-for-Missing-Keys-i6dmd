# Go Map: Silent Zero Value Return for Missing Keys

This code demonstrates a common issue with Go maps: accessing a non-existent key returns the zero value for the key's type without raising an error.  This silent behavior can lead to hard-to-debug issues.

**Problem:**
The `main` function attempts to access a non-existent key (key `2`) in the map `m`. Instead of throwing an error, Go returns the zero value for `int`, which is `0`. This might be unexpected and cause logical errors in larger programs.

**Solution:**
Use the comma ok idiom to check if a key exists in a map before using its value.

**How to reproduce:**
1. Save the code in `bug.go`.
2. Run `go run bug.go`.
3. Observe that the output is `0`, not an error.
