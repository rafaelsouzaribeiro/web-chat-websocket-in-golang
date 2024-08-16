package factory

import "fmt"

func (f *Factory) GetConnection() (*Iconnection, error) {

	switch f.types {
	case "redis":
		con, err := f.GetConRedis()

		if err != nil {
			panic(err)
		}

		c := &Iconnection{Redis: con}

		return c, nil
	case "cassandra":
		con, err := f.GetConCassandra()

		if err != nil {
			panic(err)
		}

		c := &Iconnection{Gocql: con}

		return c, nil
	}

	return nil, fmt.Errorf("invalid driver")
}
