package integration

import "os"

func gethost() map[string]string {
	host := map[string]string{}
	host["auth"] = os.Getenv("AUTH")
	return host
}
