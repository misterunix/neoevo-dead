# neoevo
Experiment in machine learning and genetic evolution.



# neoevo


## Gene Decode
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

- **00 Complete**
  - age 
  - age
- **01 Complete**
  - direction to closest food 0.0 to 1.0 - 0.0 to 360.0 , -1 if distance to far
  - clF
- **02 Complete**
  - direction to closest Neo 0.0 to 1.0 - 0.0 to 360.0
  - clN
- **03 Complete**
  - Position N to S wall 1.0 to -1.0 - middle is 0 / +1.0 = North -1.0 south
  - pNS
- **04 Complete**
  - Position W to E wall 1.0 to -1.0 - middle is 0 / +1.0 = West -1.0 East
  - pWE
- **05 Complete**
  - distance to closest food 0.0 to 1.0 max distance programmed. 30? -1.0 if to far
  - dsF
- **06 Complete**
  - distance to closest Neo 0.0 to 1.0 max distance programmed. 30? -1.0 if to far
  - dsN
- **07 Complete**
  - Hunger 0.0 to 1.0 set to 100 steps?
  - Hgr
- **08 Complete**
  - distance to blockage forward 0.0 to 1.0 max distance programmed. 30? -1.0 if to far
  - dFB
- **09 Complete**
  - distance to blockage backward 0.0 to 1.0 max distance programmed. 30? -1.0 if to far
  - dBB


## Outputs

- 00 
  - Move in a random direction
  - mRD
- 01
  - Move forward
  - mFW
- 02
  - Move backwards
  - mBK
- 03 
  - Turn Left
  - tLF
- 04 
  - Turn Right
  - tRT
- 05
  - Move North
  - mNT
- 06 
  - Move South
  - mST
- 07 
  - Move West
  - mWS
- 08 
  - Move East
- 09
  - Do Nothing
  - NOP






### Movement rule
Working off the back of the the genius that is [Steve Miller](https://github.com/davidrmiller). His idea of movement far exceeds my over my own. I humbley adapt his movement idea to my sim. This is his concept and I state his copyright here. The below is copyright by [Steve Miller](https://github.com/davidrmiller) under the MIT License. 

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
