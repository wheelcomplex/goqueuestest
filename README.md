Go Queues Test
==

Performance test of different concurrent queue implementations

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
    CFifo       92      92      87      88      88
    ZLifo       104     179     105     119     110
    ZFifo       129     160     132     147     137
    ScLifo      198     216     225     202     201
    ScFifo      206     223     195     238     207
    ZrFifo      171     334     353     310     377
    SmLifo      91      311     315     317     325
    SmFifo      99      322     313     330     325
    RcLifo      250     263     171     220     220
    RcFifo      253     260     174     221     221
    RmLifo      148     386     291     349     345
    RmFifo      147     388     303     334     348
    LcLifo      267     321     412     338     330
    LcFifo      267     323     368     355     336
    LmLifo      153     634     451     486     483
    LmFifo      153     636     481     469     537

Tests:

1. Single threaded
2. Add N, Remove N
3. N times (Add 1, Remove 1)
4. N/2 times (Add 2, Remove 1), Remove N/2
5. Add N/2, N/2 times (Add 1, Remove 2)

ZcFifo was excluded due to crash.

General recommendation use chan if it suits otherwise use slice + "chan as a lock".