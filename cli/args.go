package cli

import (
    "strings"
)

type handlerFn func(map[string]string)

var handlers []*handler

type handler struct {
    pattern string
    fn handlerFn
    description string
}

func (self *handler) splitPattern() []string {
    return strings.Split(self.pattern, " ")
}

func (self *handler) matchArgs(args []string) bool {
    patternArgs := self.splitPattern()

    if len(args) != len(patternArgs) {
        return false;
    }

    for i, patternArg := range patternArgs {
        // Skip capture groups
        if isCaptureGroup(patternArg) {
            continue
        }

        if patternArg != args[i] {
            return false
        }
    }

    return true
}

func (self *handler) getCaptureGroups(args []string) map[string]string {
    capGroups := make(map[string]string, 0)

    for i, patternArg := range self.splitPattern() {
        if isCaptureGroup(patternArg) {
            name := captureGroupName(patternArg)
            capGroups[name] = args[i]
        }
    }

    return capGroups
}

func isCaptureGroup(arg string) bool {
    return strings.HasPrefix(arg, "<") && strings.HasSuffix(arg, ">")
}

func captureGroupName(s string) string {
    return s[1:len(s) - 1]
}

func AddHandler(pattern string, fn handlerFn, desc string) {
    handlers = append(handlers, &handler{
        pattern: pattern,
        fn: fn,
        description: desc,
    })
}


func Handle(args []string) {
    h := findHandler(args)
    if h != nil {
        capGroups := h.getCaptureGroups(args)
        h.fn(capGroups)
    }
}

func findHandler(args []string) *handler {
    for _, h := range handlers {
        if h.matchArgs(args) {
            return h
        }
    }
    return nil
}
