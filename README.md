# vixenEffectFix
A tweak to FPP 2.0 effect files produced by vixen lights (as of build 819) to fix an off by one 
discrepancy with Falcon Player 2.6

As of Vixen 3.5u4 (build 819) sequences exported in the Falcon Pi Player 2.x format produces a file that is off by one with 
what FPP is expecting. The result the sequence being noticably off-color from when the sequence itself is played. 

It is very likely that future builds of FPP or Vixen will address this issue. In the meantime, I built this nifty golan app to 
fix the file.

Usage: 
vixfix -infile <infile name> -outfile <outfile name>
