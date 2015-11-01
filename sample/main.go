package main

import (
	"fmt"

	"github.com/albertoleal/features"
	"github.com/albertoleal/features/engine"
	"github.com/albertoleal/features/engine/memory"
)

func main() {
	Features := features.New(memory.New())

	feature := engine.FeatureFlag{
		Key:     "Feature X",
		Enabled: true,
	}
	Features.AddFeature(feature)

	fmt.Printf("Is `Feature X` enabled? %t \n", Features.IsActive("Feature X"))
	fmt.Printf("Is `Feature X` disabled? %t \n", Features.IsInactive("Feature X"))

	Features.With("Feature X", func() {
		fmt.Println("`Feature X` is enabled!")
	})

	Features.Without("Feature X", func() {
		fmt.Println("`Feature X` is disabled!")
	})

}
