//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package ent

import (
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/entities/moduletools"
	basesettings "github.com/weaviate/weaviate/usecases/modulecomponents/settings"
)

const (
	DefaultUpstageDocumentType    = "text"
	DefaultUpstageModel           = "solar-embedding-1-large-passage"
	DefaultVectorizeClassName    = true
	DefaultPropertyIndexed       = true
	DefaultVectorizePropertyName = false
	DefaultBaseURL               = "https://api.upstage.ai/v1/solar/embeddings"
)

type classSettings struct {
	basesettings.BaseClassSettings
	cfg moduletools.ClassConfig
}

func NewClassSettings(cfg moduletools.ClassConfig) *classSettings {
	return &classSettings{cfg: cfg, BaseClassSettings: *basesettings.NewBaseClassSettings(cfg)}
}

func (cs *classSettings) Model() string {
	return cs.BaseClassSettings.GetPropertyAsString("model", DefaultUpstageModel)
}

func (cs *classSettings) BaseURL() string {
	return cs.BaseClassSettings.GetPropertyAsString("baseURL", DefaultBaseURL)
}

func (cs *classSettings) Validate(class *models.Class) error {
	return cs.BaseClassSettings.Validate(class)
}
