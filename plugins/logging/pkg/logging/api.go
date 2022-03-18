package logging

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/rancher/opni-monitoring/pkg/auth/cluster"
	"github.com/rancher/opni-monitoring/pkg/b2bmac"
	opniv1beta2 "github.com/rancher/opni/apis/v1beta2"
	"github.com/rancher/opni/pkg/resources"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ClusterIDHeader = "OpniClusterID"
)

type OpensearchDetailsResponse struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	ExternalURL string `json:"externalURL"`
}

func (p *Plugin) ConfigureRoutes(app *fiber.App) {
	storageBackend := p.storageBackend.Get()
	clusterMiddleware := cluster.New(storageBackend, ClusterIDHeader)

	v1 := app.Group("/logging/v1", clusterMiddleware.Handle)
	v1.Get("cluster", p.handleGetOpensearchDetails)
}

func (p *Plugin) handleGetOpensearchDetails(c *fiber.Ctx) error {
	lg := p.logger

	// Fetch the cluster ID which is set by the middleware
	id := cluster.AuthorizedID(c)

	// Get the external URL
	binding := &opniv1beta2.MulticlusterRoleBinding{}
	if err := p.k8sClient.Get(p.ctx, types.NamespacedName{
		Name:      OpensearchBindingName,
		Namespace: p.storageNamespace,
	}, binding); err != nil {
		lg.Error("error fetching opensearch details", "err", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Get the credentials
	labels := map[string]string{
		resources.OpniClusterID: id,
	}
	secrets := &corev1.SecretList{}
	if err := p.k8sClient.List(p.ctx, secrets, client.InNamespace(p.storageNamespace), client.MatchingLabels(labels)); err != nil {
		lg.Error("error fetching opensearch details", "err", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if len(secrets.Items) != 1 {
		lg.Error("failed to list creds")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return details
	response := OpensearchDetailsResponse{
		Username:    secrets.Items[0].Name,
		Password:    string(secrets.Items[0].Data["password"]),
		ExternalURL: binding.Spec.OpensearchExternalURL,
	}
	responseData, err := json.Marshal(response)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	sharedKeys := cluster.AuthorizedKeys(c)
	header, err := b2bmac.NewEncodedHeader([]byte(id), responseData, sharedKeys.ServerKey)
	if err != nil {
		lg.Error("error generating response auth header", "err", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	c.Response().Header.Add("Authorization", header)
	return c.Status(fiber.StatusOK).Send(responseData)
}
