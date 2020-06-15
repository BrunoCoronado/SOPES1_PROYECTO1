#include <linux/fs.h>
#include <linux/kernel.h>   
#include <linux/init.h>
#include <linux/hugetlb.h>
#include <linux/module.h>
#include <asm/uaccess.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>

#include <linux/sched/signal.h>
#include <linux/sched.h>
 
 
struct task_struct *task;        //Estructura definida en sched.h con la informacion del proceso/tarea
struct task_struct *task_child;  //Estructura para iterar los hijos de un proceso
struct list_head *list;          //Estructura para iterar la lista de task->children

#define BUFSIZE  150
 

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Bruno Coronado - 201709362 | Max Florian - 201700465");

static int operacionEscritura(struct seq_file *archivo, void *v){

	seq_printf(archivo, "*****************************************\n");
	seq_printf(archivo, "* CARNET:   201709362    |  201700465   *\n");
	seq_printf(archivo, "* NOMBRE: Bruno Coronado | Max Florian  *\n");
	seq_printf(archivo, "*****************************************\n");	
    
    for_each_process( task ){ //MACRO para iterar los procesos alojados en linux\sched\signal.h
        seq_printf(archivo, "________________PROCESO__________________\n");	
        seq_printf(archivo, "PID:\t %d\n", task->pid);	
        seq_printf(archivo, "NOMBRE:\t %s\n", task->comm);	
        seq_printf(archivo, "ESTADO:\t %ld\n", task->state);	

        list_for_each(list, &task->children){ //MACRO para iterar entre los hijos de un proceso
            seq_printf(archivo, "\t\t<<<<<<<<<<<<<<<<<<<HIJO>>>>>>>>>>>>>>>>>>>\n");	

            task_child = list_entry( list, struct task_struct, sibling ); //se usa list_for_each para obtener todos los valores a task_child
     
            seq_printf(archivo, "\t\tPID:\t %d\n", task_child->pid);	
            seq_printf(archivo, "\t\tNOMBRE:\t %s\n", task_child->comm);	
            seq_printf(archivo, "\t\tESTADO:\t %ld\n", task_child->state);	
        }
    }    
    
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
	proc_create("cpu_201709362_201700465", 0, NULL, &op); //crea el archivo en proc

	printk(KERN_INFO "201709362 - 201700465\n");

	return 0;
}
 
static void on_exit(void){ //al desinstalar el modulo se ejecuta
	remove_proc_entry("cpu_201709362_201700465", NULL); //remueve el archivo en proc 

	printk(KERN_INFO "Sistemas Operativos 1 - USAC\n");
}

module_init(on_init);
module_exit(on_exit);
 
