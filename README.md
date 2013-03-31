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
* ZFifo : lockfree fifo implementation (probably broken)
* ZcFifo : lockfree fifo implementation using chan based freelist (crashes)
* ZrFifo : lockfree fifo implementation using ring.Ring based freelist (probably broken)
* RmLifo : ring.Ring based lifo using mutex for locking
* RmFifo : ring.Ring based fifo using mutex for locking
* SmLifo : slice based lifo using mutex for locking
* SmFifo : slice based fifo using mutex for locking

Current performance:

                1.      2.      3.      4.      5.
    CFifo       97      93      88      92      89
    LcLifo      266     325     377     339     334
    LcFifo      267     327     401     350     335
    LmLifo      152     475     517     516     477
    LmFifo      151     498     492     509     499
    ZLifo       102     154     106     132     112
    ZFifo       126     144     128     137     134
    ZrFifo      169     358     332     375     334
    RmLifo      145     397     293     336     353
    RmFifo      146     392     297     329     343
    SmLifo      99      307     313     309     311
    SmFifo      107     327     312     331     324

Tests:

1. Single threaded
2. Add N, Remove N
3. N times (Add 1, Remove 1)
4. N/2 times (Add 2, Remove 1), Remove N/2
5. Add N/2, N/2 times (Add 1, Remove 2)

ZcFifo was excluded due to crashes and infinite loops.