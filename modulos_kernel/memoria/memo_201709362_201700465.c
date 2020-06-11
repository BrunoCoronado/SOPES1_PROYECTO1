#include <linux/fs.h>
#include <linux/kernel.h>   
#include <linux/init.h>
#include <linux/hugetlb.h>
#include <linux/module.h>
#include <asm/uaccess.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>

#define BUFSIZE  150
 

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Bruno Coronado - 201709362 | Max Florian - 201700465");

struct sysinfo info; //estructura donde esta la informacion de la memoria

static int operacionEscritura(struct seq_file *archivo, void *v){
	si_meminfo(&info); //provee la informacion de /proc/meminfo hacia la estructura info

	unsigned long mem_total = ((info.totalram*info.mem_unit) / 1024) / 1024; //convertimos totalram a y usamos mem_unit
	unsigned long mem_libre = ((info.freeram*info.mem_unit) / 1024) / 1024; //convertimos freeram a y usamos mem_unit
	int utilizado = ((mem_total - mem_libre) * 100) / mem_total; //sacamos el porcentaje utilizado

	seq_printf(archivo, "*****************************************\n");
	seq_printf(archivo, "* CARNET:   201709362    |  201700465   *\n");
	seq_printf(archivo, "* NOMBRE: Bruno Coronado | Max Florian  *\n");
	seq_printf(archivo, "*****************************************\n");	
	seq_printf(archivo, "MEMORIA TOTAL: \t %lu MB \n", mem_total);
	seq_printf(archivo, "MEMORIA LIBRE: \t %lu MB \n", mem_libre);
	seq_printf(archivo, "%% UTILIZADO: \t %i%% \n", utilizado);
	seq_printf(archivo, "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n");
	return 0;
}

static int operacionAbrir(struct inode *inode, struct file *file){
	return single_open(file, operacionEscritura, NULL); //al abrir el archivo siempre ejecuta "operacionEscritura"
}

static struct file_operations op = //definimos la operaciones que va a tener el archivo
{
	.open = operacionAbrir,
	.read = seq_read
};

static int on_init(void){ //al iniciar o instalar el modulo se ejecuta
	proc_create("memo_201709362_201700465", 0, NULL, &op); //crea el archivo en proc

	printk(KERN_INFO "201709362 - 201700465\n");

	return 0;
}
 
static void on_exit(void){ //al desinstalar el modulo se ejecuta
	remove_proc_entry("memo_201709362_201700465", NULL); //remueve el archivo en proc 

	printk(KERN_INFO "Sistemas Operativos 1 - USAC\n");
}

module_init(on_init);
module_exit(on_exit);
 
