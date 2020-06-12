obj-m += cpu_201709362_201700465.o
all:
	make -C /lib/modules/$(shell uname -r)/build M=${PWD} modules
clean:
	make -C /lib/modules/$(shell uname -r)/build M=${PWD} clean