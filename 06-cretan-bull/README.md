# Labour 06: Cretan Bull
**Goal**  
For this labour you have to create a champion that will enter the Corewar arena and be worthy of it!

**Mandatory**  
You are going to have to create a champion and train them! It has to stay "live", have more instructions then just the "live" one, having functions is also mandatory (we are coders after all, damn it!) and finally it has to be able to beat the easiest pawns (Zork).

## **Commands to Run Corewar**   
**_*Note:_** *The original corewar and asm binaries are provided in `_resources/corewar.zip`. May need to `chmod` the corewar and asm binaries!*  

**Start the Corewar** *(up to 4 champions)*
```
./corewar <CHAMPION-01.COR> <CHAMPION-02.COR>...
```

**Visualize the Corewar**
```
./corewar -n <CHAMPION-01.COR> <CHAMPION-02.COR>...
```
 - press `"spacebar"` to play/pause the battle.
 - press `"q"` to slow down by 10 cycles.
 - press `"w"` to slow down by 1 cycle.
 - press `"e"` to speed up by 1 cycles.
 - press `"r"` to speed up by 10 cycles.

**Assemble Champion**
```
./asm <CHAMPION.S>
```

**Not sure what this does yet**
```
hexdump -C <CHAMPION.COR>
```