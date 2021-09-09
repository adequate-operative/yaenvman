## yaenvman

manage and generate enviroment var .source files for loading env vars into varios shells

Targeted at MSYS2 Bash and Cmd on Windows

### Proposed usage:

- `yaeman XYZ`

- ??

  - generate a dot file to load enviroment vars into current shell
  - `.cmd` / `.bat` file for Windows cmd
  - `.sh` for bash

- profits

## shimmy

Generate Shim Launchers for executables

- utilise a config file adjecent to shim executable

- config escapes `@` charater to allow injection of Enviroment Vars, Proccess args, ?other sources?

- usecase: reorganise proccess args, add env var into proc args,
