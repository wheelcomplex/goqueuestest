Go Queues Test
==

Performance test of different queue implementations

--

Warning the implementations here will change and changes probably will break API!

-- 

Available Structure:

* Channel based fifo (NewChanFifo(size))
* list.List based fifo using chan for locking (NewListCFifo)
* list.List based lifo using chan for locking  (NewListCLifo)
* list.List based fifo using mutex for locking (NewListFifo)
* list.List based lifo using mutex for locking (NewListLifo)
* lockfree lifo implementation  (NewZLifo) (broken, ABA problem)
* lockfree fifo implementation (NewZFifo) (broken?)
* lockfree fifo implementation using chan based freelist (NewZFifoFreechan)
* lockfree fifo implementation using ring.Ring based freelist (NewZFifoFreering)
* ring.Ring based lifo using mutex for locking (NewRLifo)
* ring.Ring based fifo using mutex for locking (NewRFifo)
* slice based lifo using mutex for locking (NewSLifo)
* slice based fifo using mutex for locking (NewSFifo)