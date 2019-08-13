// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package include

import (
	"fmt"

	"github.com/elastic/beats/libbeat/feature"
	"github.com/elastic/beats/x-pack/functionbeat/function/provider"
	"github.com/elastic/beats/x-pack/functionbeat/provider/gcp/gcp"
)

// Bundle exposes the trigger supported by the GCP provider.
var bundle = provider.MustCreate(
	"gcp",
	provider.NewDefaultProvider("gcp", provider.NewNullCli, provider.NewNullTemplateBuilder),
	feature.NewDetails("Google Cloud Platform", "listen to events from Google Cloud Platform", feature.Stable),
).MustAddFunction("pubsub",
	gcp.NewPubSub,
	gcp.PubSubDetails(),
).Bundle()

func init() {
	fmt.Println("registered")
	feature.MustRegisterBundle(bundle)
}
