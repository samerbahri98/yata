package server

import apiV1 "github.com/samerbahri98/yata/internal/router/api/v1"

func (b *ServerBuilder) buildAPIV1() {
	apiBuilder := apiV1.New()
	apiBuilder.AddRepository(b.repository)
	apiBuilder.AddContext(b.context)
	apiBuilder.Build()
	b.app.Mount("/api/v1", apiBuilder.GetAPI())
}
