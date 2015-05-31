package language

const assemblyContent = `section .data
    msg db "Hello World!", 0ah

section .text
    global _start
_start:
    mov rax, 1
    mov rdi, 1
    mov rsi, msg
    mov rdx, 13
    syscall
    mov rax, 60
    mov rdi, 0
    syscall`

const bashContent = `echo Hello World`

const cContent = `#include <stdio.h>

int main() {
    printf("Hello World!\n");
    return 0;
}`

const clojureContent = `(println "Hello World!")`

const cppContent = `#include <iostream>
using namespace std;

int main() {
    cout << "Hello World!";
    return 0;
}`

const csharpContent = `using System;

class MainClass {
    static void Main() {
        Console.WriteLine("Hello World!");
    }
}`

const elixirContent = `IO.puts "Hello World!"`

const erlangContent = `% escript will ignore the first line

main(_) ->
    io:format("Hello World!~n").`

const fsharpContent = `printfn "Hello World!"`

const goContent = `package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello World!")
}`

const haskellContent = `main = putStrLn "Hello World!"`

const javaContent = `class Main {
    public static void main(String[] args) {
        System.out.println("Hello World!");
    }
}`

const javascriptContent = `console.log("Hello World!");`

const luaContent = `print("Hello World!");`

const perlContent = `print "Hello World!\n";`

const phpContent = `<?php

echo "Hello World\n";`

const pythonContent = `print("Hello World!")`

const rubyContent = `puts "Hello World!"`

const rustContent = `fn main() {
    println!("Hello World!");
}`

const scalaContent = `object Main extends App {
    println("Hello World!")
}`

const plaintextContent = `Hello World!`
