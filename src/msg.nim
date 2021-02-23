import strformat

const close = "\x1B[0m"
const err = "\x1B[31m\x1B[1merror:\x1B[0m\x1B[1m"

const existsError* = &"{err} frontend already exists{close}"
const argumentsError* = &"{err} too many arguments{close}"
const warningTypescript* = &"{err} typescript not installed{close}"
const initialized* = "\x1B[32minitialized\x1b[0m"