package demo

type Server struct{
	bufferSize uint
	parallel uint
	
	msgList chan interface{}
	
	run func(interface{})
	err func(interface{})
}

func New(bufferSize, parallel uint, run func(interface{}), err func(interface{})) *Server{
	msgList := make (chan interface{}, bufferSize)
	
	s := &Server{
		bufferSize: bufferSize,
		parallel:parallel,
		msgList:msgList,
		run:run,
		err:err,
	}
	s.Consume()
	return s
}

func (s *Server)Produce(msg interface{}) {
	s.msgList <- msg
}

func (s *Server)Consume() {
	ch := make(chan struct{}, s.parallel)
	for i:=uint(0); i<s.parallel; i++{
		ch <- struct{}{}
	}
	for {
		<- ch
		go func(){
			defer func() {
				if err := recover(); err != nil {
					s.err(err)
					ch <- struct{}{}
				}
			}()
			for c := range s.msgList {
				s.run(c)
			}
		}()
	}
}

