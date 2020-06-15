
# Proyecto 1 del curso: Sistemas Operativos 1

**Desarrollada por Bruno Coronado [201709362] y Max Florian [201700465]**</br>
---

# Modulos de Kernel

## Módulo de Memoria

Para desarrolllar el modulo de memoria se utilizo la libreria **sysinfo** para obtener los datos de memoria del sistema operativo Ubuntu, esta libreria nos pertimite utilizar la siguiente estructura para obtener los datos de memoria:

    struct sysinfo {
		long uptime;             /* Seconds since boot */
		unsigned long loads[3];  /* 1, 5, and 15 minute load averages */
		unsigned long totalram;  /* Total usable main memory size */
		unsigned long freeram;   /* Available memory size */
		unsigned long sharedram; /* Amount of shared memory */
		unsigned long bufferram; /* Memory used by buffers */
		unsigned long totalswap; /* Total swap space size */
		unsigned long freeswap;  /* Swap space still available */
		unsigned short procs;    /* Number of current processes */
		char _f[22];             /* Pads structure to 64 bytes */
	};

Código Fuente: [Módulo Memoria](./modulos_kernel/memoria/memo_201709362_201700465.c)

### Instalación y Uso

Para ejecutar este modelo se deben de ejecutar los siguientes comandos dentro de la carpeta **modulos_kernel/memoria** 

 1. Compilamos el módulo<br>`$ Make`
 2. Instalamos el módulo<br>`$ insmod memo_201709362_201700465.ko`
 3. (Opcional) podemos ver los mensajes utilizando el comando<br>`$ dmesg`
 4. (Opcional) podemos ver el contenido del archivo creado en proc<br>`$ cat /proc/memo_201709362_201700465`
 5. (Opcional) para desinstalar el modulo utilizamos el comando<br>`$ rmmod memo_201709362_201700465`
 6. (Opcional) Para limpiar los archivos compilados podemos ejecutar el comando<br>`$ Make clean`
 
## Módulo de CPU

Para desarrolllar el modulo de CPU se utilizo la libreria **linux/sched.h** y **linux/sched/signal.h** para obtener los datos sobre los procesos que se estan ejecutando en el sistema operativo Ubuntu, pera obtener los datos hacemos uso de las estructuras [task_struct](https://docs.huihoo.com/doxygen/linux/kernel/3.7/structtask__struct.html) y [list_head](https://linux.die.net/man/3/list_head).

Código Fuente: [Módulo CPU](./modulos_kernel/cpu/cpu_201709362_201700465.c)

### Instalación y Uso

Para ejecutar este modelo se deben de ejecutar los siguientes comandos dentro de la carpeta **modulos_kernel/cpu** 

 1. Compilamos el módulo<br>`$ Make`
 2. Instalamos el módulo<br>`$ insmod cpu_201709362_201700465.ko`
 3. (Opcional) Podemos ver los mensajes utilizando el comando<br>`$ dmesg`
 4. (Opcional) Podemos ver el contenido del archivo creado en proc<br>`$ cat /proc/cpu_201709362_201700465`
 5. (Opcional) Para desinstalar el modulo utilizamos el comando<br>`$ rmmod cpu_201709362_201700465`
 6. (Opcional) Para limpiar los archivos compilados podemos ejecutar el comando<br>`$ Make clean`
