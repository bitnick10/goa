package pool

type InitFunc func()(interface{},error)

type Pool {
	conn chan interface{}
}

func  New(initFunc InitFunc,size int) (pool *Pool,err error){
	pool = &Pool{make(chan interface{},size)}
	for i:=0;i<size;i++{
		conn,err :=initFunc()
		if err!=nil{
			return nil,err
		}
		pool.conn<-conn
	}
	return pool,nil
}

func (pool* Pool) GetConnection() interface{}{
	return <-pool.conn
}

func (pool* Pool)ReleaseConnection(conn interface{}){
	pool.conn<-conn
}