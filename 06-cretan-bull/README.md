# Labour 05: Cretan Bull
**Goal**  
For this labour you have to create a champion that will enter the Corewar arena and be worthy of it!

**Mandatory**  
You are going to have to create a champion and train them! It has to stay "live", have more instructions then just the "live" one, having functions is also mandatory (we are coders after all, damn it!) and finally it has to be able to beat the easiest pawns (Zork).

---
## **Commands to Run Corewar**   
**_*Note:_** *The necessary binaries are provided in `_resources/corewar.zip`. May need to `chmod` the corewar and asm binaries!*  

**Start the Corewar** *(up to 4 champions)*
```
./corewar -n <CHAMPION-01.COR> <CHAMPION-02.COR>...
```

**Compile Champion**
```
./asm <CHAMPION.S>
```

**Not sure what this does yet**
```
hexdump -C <CHAMPION.COR>
```