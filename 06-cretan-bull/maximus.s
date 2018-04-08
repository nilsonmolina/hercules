.name "maximus"
.comment "Are you not entertained?!?!"

main:
	# pass the playername (r1) to the loop function at 1 byte away
	sti r1, %:loop, %1
	# create a new process and have it run the recruit function
	fork %:recruit
	# load 0 into r2. changes the 'carry' flag to 1 and allows zjmp
	ld %0, r2

# makes a 'live' call and then zjmp to run this over and over again 
loop:	
	live %13
	zjmp %:loop

# keeps running and eventually replaces zorks 'live' call
recruit:
	live %13
