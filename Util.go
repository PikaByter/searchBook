package main

type Queue struct {
	data []DirInfoAndPath
	head int
	tail int
}

func CreateQueue(size int)Queue{
	data:=make([]DirInfoAndPath,size)
	q:=Queue{}
	q.data=data
	return q
}

func (q *Queue)Empty()bool{
	if q.head==q.tail{
		return true
	}
	return false
}
//出队
func (q *Queue)Pop()DirInfoAndPath{
	res:=q.data[q.head]
	q.head++
	q.resize()
	return res
}

//入队
func (q *Queue)Push(d DirInfoAndPath){
	q.data[q.tail]=d
	q.tail++
	q.resize()
}

func (q *Queue)resize(){
	var newSize int
	qCap:=cap(q.data)
	if q.tail-q.head>qCap*4/5{
		newSize=2*qCap
	}else{
		newSize=qCap
	}
	//容量受限,或者浪费太多空间，就重新分配
	if newSize==2*qCap||q.tail>=qCap*4/5||q.head>=qCap/2{
		newData:=make([]DirInfoAndPath,newSize)
		for i:=q.head;i<q.tail;i++{
			newData[i-q.head]=q.data[i]
		}
		q.data=newData
		q.tail=q.tail-q.head
		q.head=0
	}
}

