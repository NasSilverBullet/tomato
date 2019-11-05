# tomato

Pagination package for GO

## How to use

```go
package main

import (
	"fmt"
	"os"

	"github.com/NasSilverBullet/tomato"
)

type User struct {
	ID   int
	Name string
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	us := []*User{&User{1, "nas"}, &User{2, "tomato"}, &User{3, "cucumber"}, &User{4, "green pepper"}, &User{5, "potato"}}

	p, err := tomato.New( /* current page number: */ 2 /* amount of records for one page: */, 2 /* amount of records: */, len(us))
	if err != nil {
		return err
	}

	currntPageUsInterface, err := p.FilterCurrent(us)
	if err != nil {
		return err
	}

	for _, u := range currntPageUsInterface.([]*User) {
		fmt.Printf("ID: %d => Name: %s\n", u.ID, u.Name)
		//ID: 3 => Name: cucumber
		//ID: 4 => Name: green pepper
	}

	return nil
}
```

## License

MIT License. See LICENSE.txt for more information.
