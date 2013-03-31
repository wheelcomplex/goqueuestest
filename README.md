Go Queues Test
==

Performance test of different queue implementations

--

Warning: not for production use!

-- 

Available Structure:

* CFifo : Channel based fifo
* LcLifo : list.List based lifo using chan for locking
* LcFifo : list.List based fifo using chan for locking
* LmLifo : list.List based lifo using mutex for locking
* LmFifo : list.List based fifo using mutex for locking
* ZLifo : lockfree lifo implementation (broken, ABA problem)
* ZFifo : lockfree fifo implementation (broken?)
* ZcFifo : lockfree fifo implementation using chan based freelist (broken)
* ZrFifo : lockfree fifo implementation using ring.Ring based freelist (broken)
* RmLifo : ring.Ring based lifo using mutex for locking
* RmFifo : ring.Ring based fifo using mutex for locking
* SmLifo : slice based lifo using mutex for locking
* SmFifo : slice based fifo using mutex for locking