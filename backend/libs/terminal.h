#ifndef SRCTRAN_TERMINAL_H
#define SRCTRAN_TERMINAL_H

typedef void *sectran_terminal_handle;
sectran_terminal_handle *sectran_terminal_alloc(int width, int height);
int sectran_terminal_write(sectran_terminal_handle *terminal, const char *c,
                           int size);
void sectran_terminal_stop(sectran_terminal_handle *terminal);
char *get_current_command(sectran_terminal_handle *term);
#endif