fib:
	addi	sp,sp,-16
	sw	ra,12(sp)
	sw	s0,8(sp)
	sw	s1,4(sp)
	mv	s0,a0
	addi	a5,x0,1
	ble	a0,a5,.L1
	addi	a0,a0,-1
	call	fib
	mv	s1,a0
	addi	a0,s0,-2
	call	fib
	add	a0,s1,a0
.L1:
	lw	ra,12(sp)
	lw	s0,8(sp)
	lw	s1,4(sp)
	addi	sp,sp,16
	jr	ra
