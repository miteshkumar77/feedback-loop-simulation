package paxos

//
// Paxos library, to be included in an application.
// Multiple applications will run, each including
// a Paxos peer.
//
// Manages a sequence of agreed-on values.
// The set of peers is fixed.
// Copes with network failures (partition, msg loss, &c).
// Does not store anything persistently, so cannot handle crash+restart.
//
// The application interface:
//
// px = paxos.Make(peers []string, me string)
// px.Start(seq int, v interface{}) -- start agreement on new instance
// px.Status(seq int) (Fate, v interface{}) -- get info about an instance
// px.Done(seq int) -- ok to forget all instances <= seq
// px.Max() int -- highest instance seq known, or -1
// px.Min() int -- instances before this seq have been forgotten
//

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"net/rpc"
	"os"
	"sync"
	"sync/atomic"
	"syscall"
)

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// px.Status() return values, indicating
// whether an agreement has been decided,
// or Paxos has not yet reached agreement,
// or it was agreed but forgotten (i.e. < Min()).
type Fate int

const (
	Decided   Fate = iota + 1
	Pending        // not yet decided.
	Forgotten      // decided but forgotten.
)

// Proposer ...
type Proposer struct {
	mu     sync.Mutex
	nP     int
	me     int
	parent *Paxos
	vP     interface{}
}

// ProposerManager ...
type ProposerManager struct {
	mu        sync.Mutex
	proposers map[int]*Proposer
}

// Acceptor ...
type Acceptor struct {
	mu sync.Mutex
	nP int
	nA int
	vA interface{}
}

// AcceptorManager ...
type AcceptorManager struct {
	mu        sync.Mutex
	acceptors map[int]*Acceptor
}

// Paxos hide warning
type Paxos struct {
	mu         sync.Mutex
	l          net.Listener
	dead       int32 // for testing
	unreliable int32 // for testing
	rpcCount   int32 // for testing
	peers      []string
	me         int // index into peers[]

	doneSeq []int
	maxSeq  int

	decidedVals map[int]interface{}
	acceptorMgr *AcceptorManager
	proposerMgr *ProposerManager
}

//
// call() sends an RPC to the rpcname handler on server srv
// with arguments args, waits for the reply, and leaves the
// reply in reply. the reply argument should be a pointer
// to a reply structure.
//
// the return value is true if the server responded, and false
// if call() was not able to contact the server. in particular,
// the replys contents are only valid if call() returned true.
//
// you should assume that call() will time out and return an
// error after a while if it does not get a reply from the server.
//
// please use call() to send all RPCs, in client.go and server.go.
// please do not change this function.
//
func call(srv string, name string, args interface{}, reply interface{}) bool {
	c, err := rpc.Dial("unix", srv)
	if err != nil {
		err1 := err.(*net.OpError)
		if err1.Err != syscall.ENOENT && err1.Err != syscall.ECONNREFUSED {
			fmt.Printf("paxos Dial() failed: %v\n", err1)
		}
		return false
	}
	defer c.Close()

	err = c.Call(name, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}

// PrepareArgs hide warning
type PrepareArgs struct {
	Seq int
	N   int
}

// PrepareReply hide warning
type PrepareReply struct {
	Seq  int
	N    int
	Na   int
	Va   interface{}
	Fail bool
}

// AcceptArgs hide warning
type AcceptArgs struct {
	Seq int
	N   int
	V   interface{}
}

// AcceptReply hide warning
type AcceptReply struct {
	Seq  int
	N    int
	Fail bool
}

// DecideArgs hide warning
type DecideArgs struct {
	Seq int
	V   interface{}
}

// DecideReply hide warning
type DecideReply struct {
	Seq  int
	Fail bool
}

// Prepare Acceptor's prepare handler called via RPC
func (px *Paxos) Prepare(args *PrepareArgs, reply *PrepareReply) error {
	px.mu.Lock()
	px.maxSeq = max(px.maxSeq, args.Seq)
	reply.Seq = px.doneSeq[px.me]
	px.mu.Unlock()

	px.acceptorMgr.mu.Lock()
	acceptor, ok := px.acceptorMgr.acceptors[args.Seq]
	if !ok {
		acceptor = &Acceptor{mu: sync.Mutex{}, nP: 0, nA: 0, vA: nil}
		px.acceptorMgr.acceptors[args.Seq] = acceptor
	}
	px.acceptorMgr.mu.Unlock()

	acceptor.mu.Lock()
	defer acceptor.mu.Unlock()
	if args.N > acceptor.nP {
		acceptor.nP = args.N
		reply.N = args.N
		reply.Na = acceptor.nA
		reply.Va = acceptor.vA
		reply.Fail = false
	} else {
		reply.Fail = true
		reply.N = acceptor.nP
	}
	return nil
}

// Accept Acceptor's acceptor handler called via RPC
func (px *Paxos) Accept(args *AcceptArgs, reply *AcceptReply) error {

	fmt.Printf("-- %d -- Calling accept for Seq: %d on peer: %d with N: %d and V: %v\n", px.me, args.Seq, px.me, args.N, args.V)
	px.mu.Lock()
	px.maxSeq = max(px.maxSeq, args.Seq)
	reply.Seq = px.doneSeq[px.me]
	px.mu.Unlock()
	px.acceptorMgr.mu.Lock()
	px.maxSeq = max(px.maxSeq, args.Seq)
	acceptor, ok := px.acceptorMgr.acceptors[args.Seq]
	if !ok {
		// log.Fatalf("-- %d -- Acceptor data for Seq: %d was not found!!!\n", px.me, args.Seq)
		reply.Fail = true
		px.acceptorMgr.mu.Unlock()

		return nil
	}
	px.acceptorMgr.mu.Unlock()

	acceptor.mu.Lock()
	defer acceptor.mu.Unlock()

	if args.N >= acceptor.nP {
		acceptor.nP = args.N
		acceptor.nA = args.N
		acceptor.vA = args.V
		reply.N = args.N
		reply.Fail = false
	} else {
		reply.N = args.N
		reply.Fail = true
	}

	fmt.Printf("-- %d -- Finished calling acceptor on peer: %d. Reply Fail = %v\n", px.me, px.me, reply.Fail)
	return nil
}

// Decide Acceptor's decided handler called via RPC
func (px *Paxos) Decide(args *DecideArgs, reply *DecideReply) error {

	px.mu.Lock()
	_, ok := px.decidedVals[args.Seq]
	if !ok {
		px.decidedVals[args.Seq] = args.V
	}
	reply.Seq = px.doneSeq[px.me]
	px.mu.Unlock()
	reply.Fail = false
	return nil

}

// ProposerLoop ...
func (pr *Proposer) ProposerLoop(seq int) {
	totalPeers := len(pr.parent.peers)
	majorityCount := (totalPeers / 2) + 1
	shifts := int(math.Log2(float64(totalPeers))) + 1

	fmt.Printf("-- %d -- Starting proposal loop for Seq: %d with majorityCount: %d\n", pr.me, seq, majorityCount)

	for !pr.parent.isdead() {
		s, _ := pr.parent.Status(seq)
		if s == Forgotten {
			return
		}
		if s == Decided {
			return
		}

		pr.parent.acceptorMgr.mu.Lock()
		val, ok := pr.parent.acceptorMgr.acceptors[seq]
		pr.parent.acceptorMgr.mu.Unlock()

		pr.mu.Lock()

		if ok {
			val.mu.Lock()
			pr.nP = max(pr.nP, max(val.nP, val.nA))
			val.mu.Unlock()
		}

		pr.nP = (((pr.nP >> shifts) + 1) << shifts) + pr.me
		args := PrepareArgs{Seq: seq, N: pr.nP}

		numPrepareOk := 0
		maxPrepareNa := pr.nP

		pr.mu.Unlock()

		fmt.Printf("-- %d -- Sending propose %v to %v peers\n", pr.me, args.N, totalPeers)
		var prepareWg sync.WaitGroup
		prepareWg.Add(totalPeers)
		for peerIdx := 0; peerIdx < totalPeers; peerIdx++ {
			go func(peerIdx_ int) {
				defer prepareWg.Done()
				reply := PrepareReply{}

				if peerIdx_ == pr.me {
					pr.parent.Prepare(&args, &reply)
				} else {
					if call(pr.parent.peers[peerIdx_], "Paxos.Prepare", &args, &reply) == false {
						return
					}
				}

				pr.parent.mu.Lock()
				pr.parent.doneSeq[peerIdx_] = reply.Seq
				pr.parent.mu.Unlock()
				if reply.Fail == true {
					return
				}

				pr.mu.Lock()
				defer pr.mu.Unlock()

				numPrepareOk++
				if reply.Na > maxPrepareNa {
					maxPrepareNa = reply.Na
				}

				if reply.Va != nil {

					// If an acceptor already accepted a value
					// then a majority already decided on a value
					// so we must continue Phase 1 as if we are
					// proposing the value that was already decided upon
					// by a previous majority!
					pr.vP = reply.Va
				}

			}(peerIdx)
		}

		prepareWg.Wait()
		fmt.Printf("-- %d -- Sent propose. Got back %d Oks\n", pr.me, numPrepareOk)
		if numPrepareOk < majorityCount {
			continue
		}

		pr.mu.Lock()

		if maxPrepareNa > pr.nP {
			pr.nP = maxPrepareNa
		}

		pr.mu.Unlock()

		var acceptWg sync.WaitGroup
		acceptWg.Add(totalPeers)

		fmt.Printf("-- %d -- Sending accept n: %d, v: %v to %d peers\n", pr.me, pr.nP, pr.vP, totalPeers)
		numAccepted := 0
		acceptArgs := AcceptArgs{Seq: seq, N: pr.nP, V: pr.vP}
		for peerIdx := 0; peerIdx < totalPeers; peerIdx++ {
			go func(peerIdx_ int) {
				defer acceptWg.Done()
				reply := AcceptReply{}
				if peerIdx_ != pr.me {
					if call(pr.parent.peers[peerIdx_], "Paxos.Accept", &acceptArgs, &reply) == false {
						return
					}
				} else {
					pr.parent.Accept(&acceptArgs, &reply)
				}

				pr.parent.mu.Lock()
				pr.parent.doneSeq[peerIdx_] = reply.Seq
				pr.parent.mu.Unlock()

				if reply.Fail == true {
					return
				}
				pr.mu.Lock()
				numAccepted++
				pr.mu.Unlock()

			}(peerIdx)
		}

		acceptWg.Wait()

		fmt.Printf("-- %d -- Sent accept. Got back %d Oks\n", pr.me, numAccepted)
		if numAccepted < majorityCount {
			continue
		}

		var decideWg sync.WaitGroup
		decideWg.Add(totalPeers)

		decideArgs := DecideArgs{Seq: seq, V: pr.vP}
		fmt.Printf("-- %d -- Sending decide Seq: %d, V: %v to %d peers\n", pr.me, seq, pr.vP, totalPeers)

		for peerIdx := 0; peerIdx < totalPeers; peerIdx++ {
			go func(peerIdx_ int) {
				defer decideWg.Done()
				reply := DecideReply{}
				if peerIdx_ != pr.me {
					if !call(pr.parent.peers[peerIdx_], "Paxos.Decide", &decideArgs, &reply) {
						return
					}
				} else {
					pr.parent.Decide(&decideArgs, &reply)
				}

				pr.parent.mu.Lock()
				pr.parent.doneSeq[peerIdx_] = reply.Seq
				pr.parent.mu.Unlock()

			}(peerIdx)
		}
		decideWg.Wait()
	}
}

//
// the application wants paxos to start agreement on
// instance seq, with proposed value v.
// Start() returns right away; the application will
// call Status() to find out if/when agreement
// is reached.
//
func (px *Paxos) Start(seq int, v interface{}) {
	status, _ := px.Status(seq)
	if status == Decided {
		return
	}
	if status == Forgotten {
		return
	}

	px.proposerMgr.mu.Lock()
	defer px.proposerMgr.mu.Unlock()
	if _, ok := px.proposerMgr.proposers[seq]; ok {
		return
	}

	newProp := Proposer{}
	newProp.me = px.me
	newProp.nP = 0
	newProp.vP = v
	newProp.parent = px
	px.proposerMgr.proposers[seq] = &newProp
	go newProp.ProposerLoop(seq)
}

func (px *Paxos) cleanUp(seq int) {

}

//
// the application on this machine is done with
// all instances <= seq.
//
// see the comments for Min() for more explanation.
//

func (px *Paxos) Done(seq int) {

	fmt.Printf("-- %d -- 1: Done seq: %d\n", px.me, seq)
	px.mu.Lock()
	fmt.Printf("-- %d -- 2: Done seq: %d\n", px.me, seq)

	px.acceptorMgr.mu.Lock()
	fmt.Printf("-- %d -- 3: Done seq: %d\n", px.me, seq)

	px.proposerMgr.mu.Lock()
	fmt.Printf("-- %d -- 4: Done seq: %d\n", px.me, seq)

	px.doneSeq[px.me] = max(px.doneSeq[px.me], seq)

	gMinSeq := math.MaxInt32
	for _, val := range px.doneSeq {
		gMinSeq = min(gMinSeq, val)
	}
	gMinSeq++

	deleteSeqs := make(map[int]bool)
	for key := range px.decidedVals {
		if key < gMinSeq {
			deleteSeqs[key] = true
		}
	}

	for key := range px.acceptorMgr.acceptors {
		if key < gMinSeq {
			deleteSeqs[key] = true
		}
	}

	for key := range px.proposerMgr.proposers {
		if key < gMinSeq {
			deleteSeqs[key] = true
		}
	}

	for key := range deleteSeqs {

		delete(px.decidedVals, key)
		_, ok := px.acceptorMgr.acceptors[key]
		if ok {
			px.acceptorMgr.acceptors[key] = nil
		}
		delete(px.acceptorMgr.acceptors, key)

		_, ok = px.proposerMgr.proposers[key]
		if ok {
			px.proposerMgr.proposers[key] = nil
		}
		delete(px.proposerMgr.proposers, key)
	}
	px.proposerMgr.mu.Unlock()
	px.acceptorMgr.mu.Unlock()
	px.mu.Unlock()

	fmt.Printf("-- %d -- 5.Done seq: %d\n", px.me, seq)

}

//
// the application wants to know the
// highest instance sequence known to
// this peer.
//
func (px *Paxos) Max() int {
	return px.maxSeq
}

// Min should return one more than the minimum among z_i,
// where z_i is the highest number ever passed
// to Done() on peer i. A peers z_i is -1 if it has
// never called Done().
//
// Paxos is required to have forgotten all information
// about any instances it knows that are < Min().
// The point is to free up memory in long-running
// Paxos-based servers.
//
// Paxos peers need to exchange their highest Done()
// arguments in order to implement Min(). These
// exchanges can be piggybacked on ordinary Paxos
// agreement protocol messages, so it is OK if one
// peers Min does not reflect another Peers Done()
// until after the next instance is agreed to.
//
// The fact that Min() is defined as a minimum over
// *all* Paxos peers means that Min() cannot increase until
// all peers have been heard from. So if a peer is dead
// or unreachable, other peers Min()s will not increase
// even if all reachable peers call Done. The reason for
// this is that when the unreachable peer comes back to
// life, it will need to catch up on instances that it
// missed -- the other peers therefor cannot forget these
// instances.
//
func (px *Paxos) Min() int {
	px.mu.Lock()
	defer px.mu.Unlock()
	mval := math.MaxInt32
	for _, val := range px.doneSeq {
		mval = min(mval, val)
	}
	return mval + 1
}

// Status the application wants to know whether this
// peer thinks an instance has been decided,
// and if so what the agreed value is. Status()
// should just inspect the local peer state;
// it should not contact other Paxos peers.
//
func (px *Paxos) Status(seq int) (Fate, interface{}) {

	if seq < px.Min() {
		return Forgotten, nil
	}

	px.mu.Lock()
	defer px.mu.Unlock()

	if val, ok := px.decidedVals[seq]; ok {
		return Decided, val
	}
	return Pending, nil
}

//
// tell the peer to shut itself down.
// for testing.
// please do not change these two functions.
//
func (px *Paxos) Kill() {
	atomic.StoreInt32(&px.dead, 1)
	if px.l != nil {
		px.l.Close()
	}
}

//
// has this peer been asked to shut down?
//
func (px *Paxos) isdead() bool {
	return atomic.LoadInt32(&px.dead) != 0
}

func (px *Paxos) setunreliable(what bool) {
	if what {
		atomic.StoreInt32(&px.unreliable, 1)
	} else {
		atomic.StoreInt32(&px.unreliable, 0)
	}
}

func (px *Paxos) isunreliable() bool {
	return atomic.LoadInt32(&px.unreliable) != 0
}

//
// the application wants to create a paxos peer.
// the ports of all the paxos peers (including this one)
// are in peers[]. this servers port is peers[me].
//
func Make(peers []string, me int, rpcs *rpc.Server) *Paxos {
	px := &Paxos{}
	px.peers = peers
	px.me = me

	// Your initialization code here.
	px.doneSeq = make([]int, len(px.peers))
	for i := 0; i < len(px.peers); i++ {
		px.doneSeq[i] = -1
	}
	px.maxSeq = -1
	px.decidedVals = make(map[int]interface{})
	aMgr := AcceptorManager{}
	aMgr.acceptors = make(map[int]*Acceptor)
	pMgr := ProposerManager{}
	pMgr.proposers = make(map[int]*Proposer)
	px.acceptorMgr = &aMgr
	px.proposerMgr = &pMgr

	if rpcs != nil {
		// caller will create socket &c
		rpcs.Register(px)
	} else {
		rpcs = rpc.NewServer()
		rpcs.Register(px)

		// prepare to receive connections from clients.
		// change "unix" to "tcp" to use over a network.
		os.Remove(peers[me]) // only needed for "unix"
		l, e := net.Listen("unix", peers[me])
		if e != nil {
			log.Fatal("listen error: ", e)
		}
		px.l = l

		// create a thread to accept RPC connections
		go func() {
			for px.isdead() == false {
				conn, err := px.l.Accept()
				if err == nil && px.isdead() == false {
					if px.isunreliable() && (rand.Int63()%1000) < 100 {
						// discard the request.
						conn.Close()
					} else if px.isunreliable() && (rand.Int63()%1000) < 200 {
						// process the request but force discard of reply.
						c1 := conn.(*net.UnixConn)
						f, _ := c1.File()
						err := syscall.Shutdown(int(f.Fd()), syscall.SHUT_WR)
						if err != nil {
							fmt.Printf("shutdown: %v\n", err)
						}
						atomic.AddInt32(&px.rpcCount, 1)
						go rpcs.ServeConn(conn)
					} else {
						atomic.AddInt32(&px.rpcCount, 1)
						go rpcs.ServeConn(conn)
					}
				} else if err == nil {
					conn.Close()
				}
				if err != nil && px.isdead() == false {
					fmt.Printf("Paxos(%v) accept: %v\n", me, err.Error())
				}
			}
		}()
	}

	return px
}
