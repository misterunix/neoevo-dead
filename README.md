# neoevo
An experiment in machine learning and genetic evolution.

Sometime in Oct-2021 I ran across [David R. Miller's](https://github.com/davidrmiller) video on [youtube](https://www.youtube.com/watch?v=N3tRFayqVtk). _If you haven't_ watched it I highly encourage you to do so. 

Having never worked on anything related to machine learning, this really got me excited and I just had to try. 

I didn't want to copy his code but I did use it as an imperation. The only exception is the igraph python script. 

I believe the forward propagation step is finished but needs to be tested more. **So at the moment there is nothing to see**.



## Documentaion
gomarkdoc -u > DEVDOC.md

## Gene Decode

[David's](https://github.com/davidrmiller) idea of encoding the genome in a 32 bit integer was per genius. I used his example to design my own. 

Genes are 32 bit Integers stored in each Neo's struct. 
  - bits 
    - 0 - 15 Weight on input
    - 16 - 22 Output index
    - 23
      - 0 hidden
      - 1 output
    - 24-30 Input index
    - 31 
      - 0 input
      - 1 hidden




## inputs

- **AGE - Done**
  - Index 0
  - AGE 
  - Age of Neo. 0.0 - 1.0
- DIRCLOSESTFOOD
  - Index 1  
  - DCF
  - Direction to closest food 0.0 to 1.0 - 0.0 to 360.0, -1 if no food.
- DISTANCEFOOD
  - Index 2
  - RCF
  - Distace to closest food. 0.0 - 1.0, -1 if to far.
- CLOSESTNEO        = 3  // CLOSESTNEO : Direction to closest Neo 0.0 to 1.0 - 0.0 to 360.0, -1 if no Neo.
- DISTANCENEO       = 4  // DISTANCENEO : Distance to closest Neo 0.0 to 1.0, -1 if to far.
- **POSITIONNS - Done**
  - Index 5
  - PNS
  - Position between North and South wall - IE Y -1.0 North 1.0 South.
- **POSITIONWE - Done**
  - Index 6
  - PWE
  - Position between West and East wall - IE X -1.0 West 1.0 East.
- DISTANCEFORWARD   = 7  // DISTANCEFORWARD : Distance to nearest blockage forward.
- DISTANCEBACKWARDS = 8  // DISTANCEBACKWARDS : Distance to nearest blockage backwards.
- HUNGER            = 9  // HUNGER : Hunger lever. 0.0 - 1.0, 0 not hungry, 1.0 dead.
- INPUTCOUNT        = 10 // INPUTCOUNT : number of inputs.

- MOVERANDOM    = 0 // MOVERANDOM : Move in any of the 4 directions.
- MOVEFORWARD   = 1 // MOVEFORWARD : Move in the forward direction.
- MOVEBACKWARDS = 2 // MOVEBACKWARDS : Move in the backwards direction - not turning.
- TURNLEFT      = 3 // Turn 90 degrees counter clockwise.
- TURNRIGHT     = 4 // Turn 90 degress clockwise.
- MOVENORTH     = 5 // MOVENORTH : Move y-1
- MOVESOUTH     = 6 // MOVESOUTH : Move y+1
- MOVEWEST      = 7 // MOVEWEST : Move x-1
- MOVEEAST      = 8 // MOVEEAST : Move x+1
- OUTPUTCOUNT   = 9 // OUTPUTCOUNT : number of outputs  


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




