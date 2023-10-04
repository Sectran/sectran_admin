#ifndef SRCTRAN_TERMINAL_H
#define SRCTRAN_TERMINAL_H

/*allocte a terminal*/
void *sectran_terminal_alloc(int width, int height);
/*wiret data to terminal*/
int sectran_terminal_write(void *terminal, const char *c, int size);
/*stop terminal*/
void sectran_terminal_stop(void *terminal);
/*get current command line*/
char *get_current_command(void *term);
#endif