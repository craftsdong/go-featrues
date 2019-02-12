package demo

type Server struct{
	bufferSize uint
	parallel uint
	
	msgList chan interface{}
	
	run func(interface{})
}

func New(bufferSize, parallel uint, run func(interface{})) *Server{
	msgList := make (chan interface{}, bufferSize)
	
	s := &Server{
		bufferSize: bufferSize,
		parallel:parallel,
		msgList:msgList,
		run:run,
	}
	s.Consume()
	return s
}

func (s *Server)Produce(msg interface{}) {
	s.msgList <- msg
}

func (s *Server)Consume() {
	for i :=uint(0); i < s.parallel; i++ {
		go func(){
			for c := range s.msgList {
				s.run(c)
			}
		}()
	}
}

