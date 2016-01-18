package generation

import (
	"bufio"
	"math/rand"
	"os"
)

type domains []string

func newDomains(filename string) (domains, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	domains := domains{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domains = append(domains, scanner.Text())
	}
	return domains, scanner.Err()
}

func (ds domains) random() string {
	return ds[rand.Intn(len(ds))]
}
