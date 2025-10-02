package handlers

import (
	cloudflareconnection "api-gateway/internal/pkg/CloudFlareConnection"
	"api-gateway/internal/service"
)

type HandlerSt struct {
	Service    *service.ServiceRepositoryClient
	CloudFlare *cloudflareconnection.CloudflareStorage
}

func NewHandlerSt(
	srv *service.ServiceRepositoryClient,
	CloudFlare *cloudflareconnection.CloudflareStorage,
) *HandlerSt {
	return &HandlerSt{
		Service:    srv,
		CloudFlare: CloudFlare,
	}
}
