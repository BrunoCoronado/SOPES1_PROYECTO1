#include <stdio.h>

#include <sys/sysinfo.h>

int main() {

    /*
     struct sysinfo {
        long uptime;              // Seconds since boot //
        unsigned long loads[3];  // 1, 5, and 15 minute load averages //
        unsigned long totalram;  // Total usable main memory size //
        unsigned long freeram;   // Available memory size //
        unsigned long sharedram; // Amount of shared memory //
        unsigned long bufferram; // Memory used by buffers //
        unsigned long totalswap; // Total swap space size //
        unsigned long freeswap;  // Swap space still available //
        unsigned short procs;    // Number of current processes //
        unsigned long totalhigh; // Total high memory size //
        unsigned long freehigh;  // Available high memory size //
        unsigned int mem_unit;   // Memory unit size in bytes //
        char _f[20-2*sizeof(long)-sizeof(int)];
        //Padding to 64 bytes
    };
    */
    struct sysinfo info;

    /*
     *sysinfo() retorna 0 si se ejecuto correctamente y -1 si existio un error
     * */

    if(sysinfo(&info) != 0)
        printf("\n%s\n", "Existio un error al obtener la información de /proc");

    printf("Tiempo corriendo desde el arranque: %ld segundos\n", info.uptime);
    printf("Total memoria RAM: \n->%ld KB \n->%ld MB", info.totalram / 1024, (info.totalram / 1024) / 1024);

    return 0;
}
