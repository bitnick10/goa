package pool

type InitFunc func() (interface{}, error)

type P struct {
	conn chan interface{}
}

func New(initFunc InitFunc, size int) (pool *P, err error) {
	pool = &P{make(chan interface{}, size)}
	for i := 0; i < size; i++ {
		conn, err := initFunc()
		if err != nil {
			return nil, err
		}
		pool.conn <- conn
	}
	return pool, nil
}

func (pool *P) GetConnection() interface{} {
	return <-pool.conn
}

func (pool *P) ReleaseConnection(conn interface{}) {
	pool.conn <- conn
}
