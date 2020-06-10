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

struct sysinfo info;

static int operacionEscritura(struct seq_file *archivo, void *v){
	si_meminfo(&info);
	seq_printf(archivo, "contenido del archivo");
	return 0;
}

static int operacionAbrir(struct inode *inode, struct file *file){
	return single_open(file, operacionEscritura, NULL);
}

static struct file_operations op = 
{
	.open = operacionAbrir,
	.read = seq_read
};

static int on_init(void){
	proc_create("memo_201709362_201700465", 0, NULL, &op);
	
	printk(KERN_INFO "201709362 - 201700465\n");

	return 0;
}

static void on_exit(void){
	remove_proc_entry("memo_201709362_201700465", NULL);

	printk(KERN_INFO "Sistemas Operativos 1 - USAC\n");
}

module_init(on_init);
module_exit(on_exit);
 
