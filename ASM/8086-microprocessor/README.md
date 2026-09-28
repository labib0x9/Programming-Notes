<!-- AUTO-GENERATED FROM SOURCE CODE -->
# ASM / 8086-microprocessor

## Overview

This directory explores **ASM / 8086-microprocessor** in **Assembly**.
It contains hands-on code examples and experimental scripts demonstrating core concepts, memory semantics, execution flows, and practical programming patterns.

## Learning Objectives

- Understand the implementation and mechanics of ASM / 8086-microprocessor in Assembly.
- Inspect how data structures, memory layouts, and runtime operations interact under the hood.
- Analyze execution flows, edge cases, and best practices across the provided source files.

## Files

| File | Language | Purpose |
|------|----------|---------|
| [`function.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/function.asm) | Assembly | load ds print msg1 |
| [`hello_world.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/hello_world.asm) | Assembly | Set Data Segement Copy TEXT into DX, as ah = 09h reads from DS:DX |
| [`loop.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/loop.asm) | Assembly | num1 db ? buffer db 5, 0, 5 dup('$') |
| [`print_new_line.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/print_new_line.asm) | Assembly | load ds print msg1 |
| [`print_string_loop.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/print_string_loop.asm) | Assembly | load ds print msg1 |
| [`read_one_char.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/read_one_char.asm) | Assembly | load ds print msg1 |
| [`string_input.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/string_input.asm) | Assembly | load ds print msg1 |
| [`string_input_s.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/string_input_s.asm) | Assembly | load ds print msg1 |
| [`string_to_int.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/string_to_int.asm) | Assembly | load ds print msg1 |
| [`sum.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/sum.asm) | Assembly | summation and subtraction of two number first int |

## Concepts

### 1. Function (`function.asm`)

File: [`function.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/function.asm)  
load ds print msg1.

Relevant code excerpt:

```x86asm
main proc
    ; load ds
    mov ax, @data
    mov ds, ax

    ; print msg1
    mov dx, offset num1Msg
    call print_msg

    ; input one char
    mov ah, 1
    int 21h
    mov num1, al

    ; print '\n'
    mov dx, offset newLine
    call print_msg

    ; print num1
    mov dx, offset num1
    call print_msg
```

**Explanation**:
- Assembly procedure(s): `main, print_msg` with DOS interrupt interactions.
- Demonstrates step-by-step logic and runtime behavior.

### 2. Hello World (`hello_world.asm`)

File: [`hello_world.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/hello_world.asm)  
Set Data Segement Copy TEXT into DX, as ah = 09h reads from DS:DX.

Relevant code excerpt:

```x86asm
.MODEL SMALL
.STACK 100H
.DATA
    TEXT DB "Hello, world!$"
.CODE
MAIN PROC
    ; Set Data Segement
    MOV AX, @DATA
    MOV DS, AX

    ; Copy TEXT into DX, as ah = 09h reads from DS:DX
    MOV DX, OFFSET TEXT
    MOV AH, 9
    INT 21H

    ; tell program to exit, ah = 4ch meaning terminate the program.
    MOV AH, 4CH
    INT 21H
MAIN ENDP
END MAIN
```

**Explanation**:
- Assembly procedure(s): `MAIN` with DOS interrupt interactions.
- Demonstrates step-by-step logic and runtime behavior.

### 3. Loop (`loop.asm`)

File: [`loop.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/loop.asm)  
num1 db ? buffer db 5, 0, 5 dup('$').

Relevant code excerpt:

```x86asm
main proc
    ; load ds
    mov ax, @data
    mov ds, ax

    ; loop range si = 0 to si = cl
    mov cl, 3   ; count
    mov si, 0   ; source index

    print_loop:
        ; print msg1
        mov dx, offset num1Msg
        call print_msg
        
        ; print '\n'
        mov dx, offset newLine
        call print_msg

        inc si  ; si = si + 1
        loop print_loop ; next iteration

    ; exit
```

**Explanation**:
- Assembly procedure(s): `main, print_msg` with DOS interrupt interactions.
- Demonstrates step-by-step logic and runtime behavior.

### 4. Print New Line (`print_new_line.asm`)

File: [`print_new_line.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/print_new_line.asm)  
load ds print msg1.

Relevant code excerpt:

```x86asm
main proc
    ; load ds
    mov ax, @data
    mov ds, ax

    ; print msg1
    mov dx, offset num1
    mov ah, 9
    int 21h

    ; print '\n'
    mov dx, offset newLine
    mov ah, 9
    int 21h

    ; print msg2
    mov dx, offset num2
    mov ah, 9
    int 21h

    ; exit
    mov ah, 4ch
```

**Explanation**:
- Assembly procedure(s): `main` with DOS interrupt interactions.
- Demonstrates step-by-step logic and runtime behavior.

### 5. Print String Loop (`print_string_loop.asm`)

File: [`print_string_loop.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/print_string_loop.asm)  
load ds print msg1.

Relevant code excerpt:

```x86asm
main proc
    ; load ds
    mov ax, @data
    mov ds, ax

    ; print msg1
    mov dx, offset num1Msg
    call print_msg

    ; input string
    mov dx, offset buffer
    call input_string

    ; exit
    mov ah, 4ch
    int 21h
main endp
input_string proc
    ; input string 
    mov ah, 0Ah
    int 21h
```

**Explanation**:
- Assembly procedure(s): `main, input_string, print_msg` with DOS interrupt interactions.
- Demonstrates step-by-step logic and runtime behavior.

### 6. Read One Char (`read_one_char.asm`)

File: [`read_one_char.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/read_one_char.asm)  
load ds print msg1.

Relevant code excerpt:

```x86asm
main proc
    ; load ds
    mov ax, @data
    mov ds, ax

    ; print msg1
    mov dx, offset num1Msg
    mov ah, 9
    int 21h

    ; input one char
    mov ah, 1
    int 21h
    mov num1, al

    ; print '\n'
    mov dx, offset newLine
    mov ah, 9
    int 21h

    ; print char
    mov dl, num1
```

**Explanation**:
- Assembly procedure(s): `main` with DOS interrupt interactions.
- Demonstrates step-by-step logic and runtime behavior.
## How the Code Works

Execution flow across files in this folder:

1. [`function.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/function.asm) provides or tests `function`.
2. [`hello_world.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/hello_world.asm) provides or tests `hello_world`.
3. [`loop.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/loop.asm) provides or tests `loop`.
4. [`print_new_line.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/print_new_line.asm) provides or tests `print_new_line`.

```text
Caller / Test Runner
  ├──> [function.asm] (load ds print msg1...)
  ├──> [hello_world.asm] (Set Data Segement Copy TEXT into DX...)
  ├──> [loop.asm] (num1 db ? buffer db 5, 0, 5 dup('$'...)
  ├──> [print_new_line.asm] (load ds print msg1...)
  └──> Execution Completion / Assertion
```

## System Interaction

```text
Source Code (.asm)
       │
       ▼ [MASM / Assembler]
Object File (.obj)
       │
       ▼ [LINK / Linker]
DOS Executable (.exe)
       │
       ▼ [DOSBox / Real-Mode CPU]
Segment Registers (CS, DS, SS, ES) + General Registers (AX, BX, CX, DX) -> Interrupt Handler (INT 21h)
```

## Compilation & Execution

```text
# In DOSBox (16-bit real mode):
masm function.asm;
link function;
function.exe
```

## Important Details

- **Memory & Scope**: Notice variable allocation, pointer indirection, and lifetime across function boundaries.
- **Error & Return Handling**: Always verify return values and errors before proceeding to prevent panics or undefined behavior.
- **Resource Management**: Check that open files, network sockets, database handles, and heap allocations are properly closed or freed.
- **16-bit Real Mode Addressing**: Memory is accessed via segment:offset pairs (`DS:DX`, `CS:IP`, `SS:SP`). The data segment base register `DS` must be initialized before addressing variables.
- **DOS Interrupt 21h Servicing**: `AH=01h` reads a single character into `AL`, `AH=02h` displays `DL`, `AH=09h` outputs `$`-terminated string at `DX`, and `AH=0Ah` performs buffered console input.

## Common Mistakes

- Assuming implicit synchronization or thread safety where none is provided.
- Ignoring error return values or missing zero-value edge cases.
- Misunderstanding pass-by-value versus pointer/reference semantics for complex types.
- Forgetting to terminate strings with `$` when invoking `INT 21h, AH=09h`, causing memory leakage until a random `$` is found in RAM.
- Reading ASCII character digits (`'0'` to `'9'`) and treating them directly as numerical integers without subtracting `30h` (or `'0'`).

## Language Notes

- **Language Features**: Written using idiomatic Assembly paradigms.
- **Assembly Execution**: Assembled via MASM/TASM and executed within 16-bit DOS or emulators like DOSBox.

## Related Concepts

- Data Structures & Algorithms in Assembly
- Memory Layout & Execution Lifecycles
- Systems Programming & Standard Libraries

## Questions to Test Myself

1. What is the primary role of the functions/routines demonstrated in `function.asm`?
2. How is memory allocated, managed, and freed during the execution of this code?
3. What edge cases (empty inputs, concurrency races, boundary values) could cause this code to fail?
4. How would you refactor this implementation to improve efficiency, safety, or readability?

## Further Experiments

- Add unit tests or benchmark functions to measure performance and correctness under varying input sizes.
- Modify `function.asm` to handle error conditions or invalid inputs gracefully.
- Introduce concurrency or parallel execution where applicable and inspect the synchronization behavior.

## Source Files

- [`function.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/function.asm)
- [`hello_world.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/hello_world.asm)
- [`loop.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/loop.asm)
- [`print_new_line.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/print_new_line.asm)
- [`print_string_loop.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/print_string_loop.asm)
- [`read_one_char.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/read_one_char.asm)
- [`string_input.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/string_input.asm)
- [`string_input_s.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/string_input_s.asm)
- [`string_to_int.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/string_to_int.asm)
- [`sum.asm`](file:///Users/labib0x9/Desktop/Programming-Notes/ASM/8086-microprocessor/sum.asm)
