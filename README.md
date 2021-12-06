# neoevo
An experiment in machine learning and genetic evolution.

Sometime in Oct-2021 I ran across [David R. Miller's](https://github.com/davidrmiller) video on [youtube](https://www.youtube.com/watch?v=N3tRFayqVtk). _If you haven't_ watched it I highly encourage you to do so. 

Having never worked on anything related to machine learning, this really got me excited and I just had to try. 

I didn't want to copy his code but I did use his comments as an imperation. The only exception is the igraph python script. That I did copy.

**So at the moment there is nothing to see**.


## Documentaion
gomarkdoc -u > DEVDOC.md

## Gene Decode

????|PPER|TTAA|QQQQ

E Destination Layer 0-F
R Source Layer 0-F
TT Destination ID 00-FF
AA Source ID 00-FF
QQQQ Weight 65535 : (X / 65535)*8-4 : -4.0 to +4.0





## inputs



### Movement rule
Working off the back of the genius that is [Steve Miller](https://github.com/davidrmiller). His idea of movement far exceeds my own. I humbly adapt his movement idea to my sim. This is his concept and I state his copyright here. The below is copyright by [Steve Miller](https://github.com/davidrmiller) under the MIT License. 

    // There are multiple action neurons for movement. Each type of movement neuron
    // urges the individual to move in some specific direction. We sum up all the
    // X and Y components of all the movement urges, then pass the X and Y sums through
    // a transfer function (tanh()) to get a range -1.0..1.0. The absolute values of the
    // X and Y values are passed through prob2bool() to convert to -1, 0, or 1, then
    // multiplied by the component's signum. This results in the x and y components of
    // a normalized movement offset. I.e., the probability of movement in either
    // dimension is the absolute value of tanh of the action level X,Y components and
    // the direction is the sign of the X, Y components. For example, for a particular
    // action neuron:
    //     X, Y == -5.9, +0.3 as raw action levels received here
    //     X, Y == -0.999, +0.29 after passing raw values through tanh()
    //     Xprob, Yprob == 99.9%, 29% probability of X and Y becoming 1 (or -1)
    //     X, Y == -1, 0 after applying the sign and probability
    //     The agent will then be moved West (an offset of -1, 0) if it's a legal move.

### Cardinal directions

|     |     |     |   |       |       |       |
|:----|:---:| ---:|:-:|:------|:-----:|------:|
| 225 | 270 | 315 |   | 0.625 | 0.750 | 0.875 |
| 180 |  *  | 000 |   | 0.500 |   *   | 0.000 |
| 135 | 090 | 045 |   | 0.375 | 0.250 | 0.125 |
|     |     |     |   |       |       |       |




