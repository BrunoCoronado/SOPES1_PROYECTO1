# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.15

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Remove some rules from gmake that .SUFFIXES does not remove.
SUFFIXES =

.SUFFIXES: .hpux_make_needs_suffix_list


# Suppress display of executed commands.
$(VERBOSE).SILENT:


# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /snap/clion/99/bin/cmake/linux/bin/cmake

# The command to remove a file.
RM = /snap/clion/99/bin/cmake/linux/bin/cmake -E remove -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/trext/Documentos/Sopes/SOPES1_PROYECTO1/Pruebas

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/trext/Documentos/Sopes/SOPES1_PROYECTO1/Pruebas/cmake-build-debug

# Include any dependencies generated for this target.
include CMakeFiles/Pruebas.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/Pruebas.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/Pruebas.dir/flags.make

CMakeFiles/Pruebas.dir/main.c.o: CMakeFiles/Pruebas.dir/flags.make
CMakeFiles/Pruebas.dir/main.c.o: ../main.c
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/trext/Documentos/Sopes/SOPES1_PROYECTO1/Pruebas/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building C object CMakeFiles/Pruebas.dir/main.c.o"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -o CMakeFiles/Pruebas.dir/main.c.o   -c /home/trext/Documentos/Sopes/SOPES1_PROYECTO1/Pruebas/main.c

CMakeFiles/Pruebas.dir/main.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/Pruebas.dir/main.c.i"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/trext/Documentos/Sopes/SOPES1_PROYECTO1/Pruebas/main.c > CMakeFiles/Pruebas.dir/main.c.i

CMakeFiles/Pruebas.dir/main.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/Pruebas.dir/main.c.s"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/trext/Documentos/Sopes/SOPES1_PROYECTO1/Pruebas/main.c -o CMakeFiles/Pruebas.dir/main.c.s

# Object files for target Pruebas
Pruebas_OBJECTS = \
"CMakeFiles/Pruebas.dir/main.c.o"

# External object files for target Pruebas
Pruebas_EXTERNAL_OBJECTS =

Pruebas: CMakeFiles/Pruebas.dir/main.c.o
Pruebas: CMakeFiles/Pruebas.dir/build.make
Pruebas: CMakeFiles/Pruebas.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/trext/Documentos/Sopes/SOPES1_PROYECTO1/Pruebas/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking C executable Pruebas"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/Pruebas.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/Pruebas.dir/build: Pruebas

.PHONY : CMakeFiles/Pruebas.dir/build

CMakeFiles/Pruebas.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/Pruebas.dir/cmake_clean.cmake
.PHONY : CMakeFiles/Pruebas.dir/clean

CMakeFiles/Pruebas.dir/depend:
	cd /home/trext/Documentos/Sopes/SOPES1_PROYECTO1/Pruebas/cmake-build-debug && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/trext/Documentos/Sopes/SOPES1_PROYECTO1/Pruebas /home/trext/Documentos/Sopes/SOPES1_PROYECTO1/Pruebas /home/trext/Documentos/Sopes/SOPES1_PROYECTO1/Pruebas/cmake-build-debug /home/trext/Documentos/Sopes/SOPES1_PROYECTO1/Pruebas/cmake-build-debug /home/trext/Documentos/Sopes/SOPES1_PROYECTO1/Pruebas/cmake-build-debug/CMakeFiles/Pruebas.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/Pruebas.dir/depend

