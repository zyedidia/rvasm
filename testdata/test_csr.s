csrrs t0, instreth, a0
rdinstreth t0
rdcycle t0
rdtime t0
csrr t0, instret
csrw cycleh, t1
csrs uip, a0
mret
sret
uret
wfi
