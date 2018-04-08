.name "maximus"
.comment "Are you not entertained?!?!"

main:
	sti r1, %:loop, %1
	fork %:recruit
	ld %0, r2
	
loop:	
	live %13
	zjmp %:loop

recruit:
	live %13
