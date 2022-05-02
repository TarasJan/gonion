package gonion

import (
    "errors"
    "fmt"
)

type gontext map[string]interface{}

type GonionBody func(gontext) (gontext, error)

type Gonion struct {
    requiredFields []string
    nextGonion *Gonion
    body GonionBody
}

func (g *Gonion) Wrap(newGonion *Gonion) *Gonion {
    g.nextGonion = newGonion
    return g
}

func (g *Gonion) Call(context gontext) (gontext, error) {
    for _, key := range g.requiredFields {
        if context[key] == nil {
            return context, errors.New(fmt.Sprintf("%s key is missing from context - gonion execution cannot continue", key))
        }
    }

    newContext, err := g.body(context)

    if err != nil {
        return context, err
    }

    if g.nextGonion != nil {
        return g.nextGonion.Call(context)
    } else {
        return newContext, nil
    }
}



