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


### Cardinal directions

|     |     |     |   |       |       |       |
|:----|:---:| ---:|:-:|:------|:-----:|------:|
| 225 | 270 | 315 |   | 0.625 | 0.750 | 0.875 |
| 180 |  *  | 000 |   | 0.500 |   *   | 0.000 |
| 135 | 090 | 045 |   | 0.375 | 0.250 | 0.125 |
|     |     |     |   |       |       |       |
