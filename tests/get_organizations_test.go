/*
Copyright (c) 2021 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"                      // nolint
	. "github.com/onsi/gomega"                         // nolint
	. "github.com/onsi/gomega/ghttp"                   // nolint
	. "github.com/openshift-online/ocm-sdk-go/testing" // nolint
	"github.com/openshift-online/ocm-support-cli/pkg/organization"
	"github.com/openshift-online/ocm-support-cli/pkg/types"
)

var _ = Describe("List orgs", func() {
	var ctx context.Context

	BeforeEach(func() {
		// Create a context:
		ctx = context.Background()
	})

	When("Config file doesn't exist", func() {
		It("Fails", func() {
			getResult := NewCommand().
				Args(
					"get", "orgs", "Red Hat",
				).Run(ctx, ocmSupportBinary)
			Expect(getResult.ExitCode()).ToNot(BeZero())
			Expect(getResult.ErrString()).To(ContainSubstring("Not logged in"))
		})
	})

	When("Config file doesn't contain valid credentials", func() {
		It("Fails", func() {
			getResult := NewCommand().
				ConfigString(`{}`).
				Args(
					"get", "orgs",
				).Run(ctx, ocmCliBinary)
			Expect(getResult.ExitCode()).ToNot(BeZero())
			Expect(getResult.ErrString()).To(ContainSubstring("Not logged in"))
		})
	})

	When("Config file contains valid credentials", func() {
		var ssoServer *Server
		var apiServer *Server
		var config string

		BeforeEach(func() {
			// Create the servers:
			ssoServer = MakeTCPServer()
			apiServer = MakeTCPServer()

			// Create the token:
			accessToken := MakeTokenString("Bearer", 15*time.Minute)

			// Prepare the server:
			ssoServer.AppendHandlers(
				RespondWithAccessToken(accessToken),
			)

			// Login:
			result := NewCommand().
				Args(
					"login",
					"--client-id", "my-client",
					"--client-secret", "my-secret",
					"--token-url", ssoServer.URL(),
					"--url", apiServer.URL(),
				).
				Run(ctx, ocmCliBinary)
			Expect(result.ExitCode()).To(BeZero())
			config = result.ConfigString()
		})

		AfterEach(func() {
			// Close the servers:
			ssoServer.Close()
			apiServer.Close()
		})

		It("Get organizations from EbsAccountID", func() {
			mockOrganizations := []organization.Organization{
				{
					Name:         "My Org 1",
					EbsAccountID: "1",
					Meta: types.Meta{
						ID:   "123",
						HREF: "/api/accounts_mgmt/v1/organizations/123",
					},
				},
				{
					Name:         "My Org 2",
					EbsAccountID: "1",
					Meta: types.Meta{
						ID:   "456",
						HREF: "/api/accounts_mgmt/v1/organizations/456",
					},
				},
			}
			orgs, err := SerializeOrganizations(mockOrganizations)
			Expect(err).To(BeNil())
			// Prepare the server:
			apiServer.AppendHandlers(
				RespondWithJSON(
					http.StatusOK,
					`{
						"kind": "OrganizationList",
						"page": 1,
						"size": 2,
						"total": 2,
						"items": `+
						string(orgs)+
						`}`,
				),
			)

			// Run the command:
			result := NewCommand().
				ConfigString(config).
				Args("get", "orgs", "1").
				Run(ctx, ocmSupportBinary)
			Expect(result.ExitCode()).To(BeZero())
			Expect(result.ErrString()).To(BeEmpty())
			lines := result.OutLines()
			receivedOrgs, err := DeserializeOrganizations(lines)
			Expect(err).To(BeNil())
			Expect(reflect.DeepEqual(mockOrganizations, receivedOrgs)).To(BeTrue())
		})

		It("Get matching organization from list of organizations ", func() {
			matchingOrgName := "My Org 1"
			mockOrganizations := []organization.Organization{
				{
					Name:         matchingOrgName,
					EbsAccountID: "1",
					Meta: types.Meta{
						ID:   "123",
						HREF: "/api/accounts_mgmt/v1/organizations/123",
					},
				},
			}
			orgs, err := SerializeOrganizations(mockOrganizations)
			Expect(err).To(BeNil())
			// Prepare the server:
			apiServer.AppendHandlers(
				RespondWithJSON(
					http.StatusOK,
					`{
						"kind": "OrganizationList",
						"page": 1,
						"size": 1,
						"total": 1,
						"items": `+
						string(orgs)+
						`}`,
				),
			)

			// Run the command:
			result := NewCommand().
				ConfigString(config).
				Args("get", "orgs", matchingOrgName).
				Run(ctx, ocmSupportBinary)
			Expect(result.ExitCode()).To(BeZero())
			Expect(result.ErrString()).To(BeEmpty())
			lines := result.OutLines()
			receivedOrgs, err := DeserializeOrganizations(lines)
			fmt.Println(len(receivedOrgs))
			Expect(err).To(BeNil())
			Expect(len(receivedOrgs)).To(Equal(1))
			Expect(receivedOrgs[0].Name).To(Equal(matchingOrgName))
		})
	})
})

func SerializeOrganizations(orgs []organization.Organization) ([]byte, error) {
	orgBytes, err := json.Marshal(orgs)
	if err != nil {
		return nil, err
	}
	return orgBytes, nil
}

func DeserializeOrganizations(result []string) ([]organization.Organization, error) {
	line := strings.Join(result, "")
	var receivedOrgs []organization.Organization
	err := json.Unmarshal([]byte(line), &receivedOrgs)
	if err != nil {
		return nil, err
	}
	return receivedOrgs, nil
}
