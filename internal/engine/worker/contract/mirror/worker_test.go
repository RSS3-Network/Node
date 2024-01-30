package mirror_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/cockroachdb"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/database/dialer"
	source "github.com/rss3-network/node/internal/engine/source/arweave"
	worker "github.com/rss3-network/node/internal/engine/worker/contract/mirror"
	"github.com/rss3-network/node/provider/arweave"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestWorker_Arweave(t *testing.T) {
	t.Parallel()

	type arguments struct {
		task   *source.Task
		config *config.Module
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      *schema.Feed
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "Mirror Post",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkArweave,
					Block: arweave.Block{
						Timestamp: 1608772429,
					},
					Transaction: arweave.Transaction{
						ID:       "lW0AMDN2RgOeqULk-u6Tv0wfZWpx9MfkrmqQQU-Mvuo",
						Reward:   "27798031",
						Quantity: "0",
						Owner:    "x3EoIMrB1qOrAivgqa1fi3SwJQgZdVuLgZhnvB0NLVngTIfB9yytGQVr9jJ4w4zn6q80T_L8aU5yXVhB12cfoABwP6AYDbvLefXLCGFgUcEcwrQ_bhWmDpxEjU2sMu4qrlI2PVj_3y1-Cx1wXldVD3i2JT81BeXacj8f19ZpKFZHaSvLM7jHm5STqwNaZDhhNnwR9qwe5kKkjafvOVDF9ayp7NHLOsEq5s8rcaQKZMYR_yEbuSfE4jzuq5F3Cwcp18bzooOqgtpCi1LozAlWepYywrxuhc-TxxObar2FVZU7bHyuNMD_aLoZ7mYqHUDGwPocqMqZVa6zQPAm7J1mnyI-fab-3E9vlZ3ZEO-PD21EAo_ALidmC3jR35tzWTGx6y7D4W1J_biJc_iKBybmsOH5W0K31CMW1sO2P2FKndUVz8vzedYvNbjYJhFFqAaEcn2x4mj15n2d75_DyEqmyEMXZ6MAg2-hUixYky7Gk7hpokwbwWEu35FsEIh4VKoanxJPgZ7JCWzt7su5su891RVf0Aodvr8UR971o1xgGaTR5ChLBC4SjyswcRTzNIb5LiPoJK4Sp03wAQW7dGgbKo8W5XjLsI-KWqypEyRkQYy-GYzwIWZ-4_5qzARaZOkVS_mn2aZzbsxxZNWxtjpJVNOd8olr0Fo88nn1xnLu2b0",
						Data:     "eyJjb250ZW50Ijp7ImJvZHkiOiJTdGF0ZWZ1bCBXb3JrcyBpcyBleGNpdGVkIHRvIHB1YmxpY2x5IGludHJvZHVjZSBvdXIgZmlyc3QgcHJvamVjdCwgYSBsaW1pdGVkIHJ1biBwcmludCBib29rIGZvY3VzZWQgb24gdGhlIHBlb3BsZSBhbmQgdGVhbXMgdGhhdCBicm91Z2h0IHRoZSBFdGhlcmV1bSBCZWFjb24gQ2hhaW4gdG8gbGlmZS5cblxuIyMjIFdobyBpcyBTdGF0ZWZ1bCBXb3Jrcz9cblxuV2UgYXJlIGEgZ3JvdXAgb2YgcGFzc2lvbmF0ZSBsb25ndGltZSBjb21tdW5pdHkgbWVtYmVycyBpbnRlcmVzdGVkIGluIGV4cGxvcmluZyBhbmQgY3JlYXRpbmcgRXRoZXJldW0gY3VsdHVyZTogcmVhY2ggb3V0IHRvIFtAc3RhdGVmdWx3b3Jrc10oaHR0cHM6Ly90d2l0dGVyLmNvbS9TdGF0ZWZ1bFdvcmtzKSBvbiBUZWxlZ3JhbSBvciBUd2l0dGVyIHdpdGggcXVlc3Rpb25zLlxuXG4jIyMgQm9vayBEZXRhaWxzXG5cblRoaXMgYm9vayB3aWxsIGZlYXR1cmUgcHJvZmlsZXMsIHJlZmxlY3Rpb25zLCBhbmQgc3RvcmllcyBmcm9tIEV0aDIgaW1wbGVtZW50ZXJzIGFuZCByZXNlYXJjaGVycy4gV2Ugd2FudCB0byBzaGFyZSB0aGUgc3RvcmllcyBvZiB0aGUgcGVvcGxlIHRoYXQgbWFkZSB0aGlzIHBvc3NpYmxlLiBXZSBoYXZlIGJlZW4gd29ya2luZyBvbiB0aGlzIGZvciB0aGUgcGFzdCBmZXcgbW9udGhzLCBhbmQgaGF2ZSBtYWRlIHNpZ25pZmljYW50IHByb2dyZXNzIHRvd2FyZHMgZGVmaW5pbmcgYW5kIHJlYWxpemluZyB0aGUgdmlzaW9uLlxuXG5TbyBmYXIsIHdlIGhhdmUgZ2F0aGVyZWQgcHJvZmlsZXMgZnJvbSBvdmVyICoqMjUgRXRoMiBjb250cmlidXRvcnMqKiwgd2l0aCBtYW55IG1vcmUgdG8gY29tZS4gV2UncmUgYW50aWNpcGF0aW5nIGRlbGl2ZXJpZXMgb2YgdGhlIGJvb2sgKipHZW5lc2lzIEVkaXRpb24gaW4gUTEgMjAyMS4qKlxuXG5BcyBhIGxpbWl0ZWQgcnVuIGZpcnN0IHByaW50LCB0aGUgYm9vayB3aWxsIGJlIGEgaGlnaCBxdWFsaXR5IGNvbGxlY3RvcnMgaXRlbS4gV2UgYXJlIHBheWluZyBjbG9zZSBhdHRlbnRpb24gdG8gZXZlcnkgYXNwZWN0LCBmcm9tIHRoZSBvdmVyYWxsIGRlc2lnbiwgdG8gdGhlIGZlYXR1cmVzIGhpZ2hsaWdodGluZyB0aGUgdW5pcXVlbmVzcyBvZiBlYWNoIGNvcHkuIEZvciBleGFtcGxlLCB0aGUgYnVybiB0cmFuc2FjdGlvbiBoYXNoIG9mIHJlZGVtcHRpb24gYmVpbmcgbWFya2VkIG9uIHRoZSBwaHlzaWNhbCBwcmludC4gXG5cbiMjIyBXaHkgYSBCb29rP1xuXG5XZSB0aGluayB0aGUgYXJ0aWZhY3QgdGhhdCBmcmFtZXMgdGhlIHdvcmsgY2FuIGJlIGFzIGltcG9ydGFudCBhcyB0aGUgY29udGVudCBpdHNlbGYuIFNvIG11Y2ggb2Ygd2hhdCB3ZSBjb25zdW1lIGV4aXN0cyBlcGhlbWVyYWxseSBpbiBmZWVkcyBhbmQgY2hhdHM6IHRoZSBCZWFjb24gQm9vayBjYXB0dXJlcyBhIHNuYXBzaG90IG9mIHNlbnRpbWVudCwgcG9ydGVkIGludG8gb2JqZWN0LWZvcm0uIFRoaXMgYXJ0aWZhY3Qgd2lsbCBsaXZlIG9uIGFzIHNvbWV0aGluZyB0byBiZSByZWFkLCByZWZlcmVuY2VkLCBhbmQgY2FyZWZ1bGx5IHN0b3JlZCBsb25nIGludG8gdGhlIGZ1dHVyZS5cblxuIyMjIEhvdyBjYW4gSSBnZXQgb25lP1xuXG5IZXJlIGFyZSB0aGUgZ2VuZXJhbCBkZXRhaWxzLCBzdWJqZWN0IHRvIGNoYW5nZS4gXG5cbioqTk9URTogVGhlcmUgaXMgbm8gd2F5IHRvIGFjcXVpcmUgdGhlIGJvb2sgYXQgdGhpcyB0aW1lLiBUaGUgZGF0ZSBhbmQgcmVsZWFzZSB3aWxsIGJlIHdlbGwgcHVibGljaXplZCBoZXJlIGFuZCB3aXRoIG91ciBwYXJ0bmVyIG9yZ3MuKipcblxuVGhlIGJ1bGsgb2YgdGhlIEdlbmVzaXMgRWRpdGlvbiAofjEwMCBib29rcykgd2lsbCBiZSBkaXN0cmlidXRlZCBvbi1jaGFpbiwgc2ltaWxhciB0byBVbmlzb2Nrcy4gVGhlIG1haW4gZGlmZmVyZW5jZSBpcyB0aGF0IHdlIHBsYW4gdG8gdXNlIGEgQmFsYW5jZXIgU21hcnQgcG9vbCB0byBhbGxvdyBtYXhpbWFsIGFjY2VzcyBmb3IgdGhlIGVudGlyZSBFdGhlcmV1bSBjb21tdW5pdHkuIExlYXJuIG1vcmUgYWJvdXQgaG93IHRoZXNlIG1lY2hhbmlzbXMgZnVuY3Rpb24gdGhyb3VnaCBjYXNlIHN0dWRpZXMgbGlrZSB0aGUgW0JhbmtsZXNzIEJBUF0oaHR0cHM6Ly9iYW5rbGVzcy5zdWJzdGFjay5jb20vcC90aGUtdWx0aW1hdGUtZ3VpZGUtdG8tYmFsYW5jZXItc21hcnQpIGFuZCBbUEVSUF0oaHR0cHM6Ly9wZXJwZXR1YWxwcm90b2NvbC5tZWRpdW0uY29tL3doeS13ZS1jaG9zZS10by1kaXN0cmlidXRlLXBlcnAtdXNpbmctYS1iYWxhbmNlci1saXF1aWRpdHktYm9vdHN0cmFwcGluZy1wb29sLWFhYzdmMWFiNjE4MSkuIFxuXG5BIGZldyBvZiB0aGUgYmVuZWZpdHMgaW5jbHVkZTpcblxuLSBiZXR0ZXIgcHJpY2UgZGlzY292ZXJ5OiBzdGFydHMgaGlnaCA-IGxvdywgbWFraW5nIGl0IGhhcmRlciBmb3Igd2hhbGVzIHRvIGNvcm5lciBzdXBwbHkgZWFybHkgb25cbi0gZmFpciBkaXN0cmlidXRpb24gZm9yIGh1bWFucyAoYm90cyBoYXZlIG5vIGFkdmFudGFnZSlcbi0gZWZmaWNpZW50IGZvciBib290c3RyYXBwaW5nIGxpcXVpZGl0eSBmb3IgdGhlIHBvb2wgY3JlYXRvclxuXG5BIHBvcnRpb24gb2YgdGhlIEdlbmVzaXMgRWRpdGlvbiAofjEwIGJvb2tzKSB3aWxsIGJlIHBsZWRnZWQgdG8gcGFydG5lciBjb21tdW5pdHkgb3JncyB0byBnaXZlIGF3YXkgdG8gdGhlaXIgY29tbXVuaXRpZXMuIEV0aGVyZXVtIGN1bHR1cmUgaXMgbm90aGluZyBpZiBub3QgcGFydGljaXBhdG9yeSwgYW5kIHdlIHdhbnQgdG8gZnVsbHkgZW1icmFjZSB0aGlzIGFzcGVjdC4gV2Ugd2lsbCBiZSBhbm5vdW5jaW5nIG91ciBwYXJ0bmVycyBjbG9zZXIgdG8gdGhlIGNvbXBsZXRlIGxhdW5jaC5cblxuVGhlcmUgd2lsbCBiZSBvdGhlciBmdW4gd2F5cyB0byB3aW4gYm9va3MsIHN0YXkgdHVuZWQgZm9yIG1vcmUgZGV0YWlscyEgSWYgeW91IGVuZCB1cCBub3QgZ2V0dGluZyBhIGJvb2ssIGRvbid0IHdvcnJ5OiBkaWdpdGFsIHZlcnNpb25zIG9mIHRoZSBjb250ZW50IHdpbGwgYmUgYXZhaWxhYmxlIGZvciBmcmVlIHRvIGVuam95LCBoZXJlIG9uIHRoaXMgYmxvZy5cblxuRmluYWxseSwgYW4gaW1wb3J0YW50IGdvYWwgb2YgdGhpcyBwcm9qZWN0IGlzIHRvIHByb21vdGUgYW5kIHN1cHBvcnQgRXRoZXJldW0gUHVibGljIEdvb2RzLiBUaGVyZWZvcmUsIGFueSBuZXQgRVRIIGdlbmVyYXRlZCBieSB0aGUgR2VuZXNpcyBzYWxlIG5ldCBvZiBwcmludGluZywgc2hpcHBpbmcgYW5kIGNvbnRyaWJ1dG9yIGNvc3RzIHdpbGwgYmUgZG9uYXRlZCB0byB0aGUgW09wZW4gR3JhbnRzIEV0aDIgc3RyZWFtXShodHRwczovL29wZW5ncmFudHMuY29tL2dyYW50LzB4NTNlN2RhYThlM2FhMjNjZDMwYzc1YjJmNTk5YzMwM2JhZGExNzA2NCkuXG5cbi0tLVxuXG5Zb3UgbWF5IGhhdmUgbm90aWNlZCB0aGlzIHBvc3QgaXMgaG9zdGVkIG9uIFtNaXJyb3IueHl6XShodHRwczovL21pcnJvci54eXovKSwgYSBuZXcgd2ViMyBwcm90b2NvbCBmb3IgcHVibGlzaGluZyB0byBjb21tdW5pdGllcy4gV2UncmUgZXhjaXRlZCB0byBiZSBvbmUgb2YgdGhlIGZpcnN0IHByb2plY3RzIHRvIHB1Ymxpc2ggdXNpbmcgaXQhIENoZWNrIGJhY2sgbmV4dCB3ZWVrIGZvciBtb3JlIGluZm8gb24gd2hhdCBTdGF0ZWZ1bCBXb3JrcyBpcyBhbGwgYWJvdXQgLSB3ZSBjYW4ndCB3YWl0IHRvIHNoYXJlIG1vcmUgYWJvdXQgb3VyIHZpc2lvbi5cblxuTWVzc2FnZSB1cyBvbiBUZWxlZ3JhbSBvciBUd2l0dGVyIFtAc3RhdGVmdWx3b3Jrc10oaHR0cHM6Ly90d2l0dGVyLmNvbS9TdGF0ZWZ1bFdvcmtzKSB3aXRoIHF1ZXN0aW9ucy4g4p2k77iPIiwidGltZXN0YW1wIjoxNjA4NzY2MDcwLCJ0aXRsZSI6IldlJ3JlIG1ha2luZyBhIGJvb2sgYWJvdXQgRXRoMiEiLCJwdWJsaWNhdGlvbiI6InN0YXRlZnVsIn0sImRpZ2VzdCI6ImExNTFlZTFkZWNiMjAyOGE4YmI0ODI3N2Y2OTI4YzZmMzgzMTljMzI2MDFkYzFkYTFlZTgyYWNmY2FkMmU1MjUiLCJhdXRob3JzaGlwIjp7ImNvbnRyaWJ1dG9yIjoiMHg0QzBhNDY2REYwNjI4RkU4Njk5MDUxYjNBYzY1MDY2NTMxOTFjYzIxIiwic2lnbmluZ0tleSI6Ii0tLS0tQkVHSU4gRUMgUFVCTElDIEtFWS0tLS0tXHJcbk1Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMERBUWNEUWdBRWNMKytaNHdpdDFOWTYrb1VkaDc5bVFydHdzaVVcclxuQktVdXg5MHRpcGh6S0I0Zk9iWm1oK1pQYkJQNGhia3hCUGhHMzllMU5JeWdRZmY3Skc0YS9wNDczZz09XHJcbi0tLS0tRU5EIEVDIFBVQkxJQyBLRVktLS0tLVxyXG4iLCJzaWduYXR1cmUiOiI4akNxclJmZUxTNTZkbGhjaG10NUU2K0ltSTJhMWpBTzR6SzNtMThudkR3Q1N0K1ZOSlJ1dS8yQVJwQVpoRjlBd0F3OHo5MzlMaHp1d0dBVnZwRXVBdz09Iiwic2lnbmluZ0tleVNpZ25hdHVyZSI6IjB4ZjJjOWM0N2I1MjA2ZGIzNTQyOWMyNmYxOTVhMzJkMjk3OGM5YTU5Njc2ZjVlZmU5YWQyMzgxMDc5MDg3ODMwZTQ2ODU4YjMzY2NkMzNmN2M3MDhmZmFlMTc2NDk0YjBkMDhkZTNhOTVkMDU1YTg0NjE0MTE1ZDNhZmU0ZDlmYjgxYyIsInNpZ25pbmdLZXlNZXNzYWdlIjoiSSBhdXRob3JpemUgdGhlIHB1YmxpY2F0aW9uIG9mIGFydGljbGVzIG9uIHN0YXRlZnVsLm1pcnJvci54eXogZnJvbSB0aGlzIGRldmljZSB1c2luZzpcbi0tLS0tQkVHSU4gRUMgUFVCTElDIEtFWS0tLS0tXHJcbk1Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMERBUWNEUWdBRWNMKytaNHdpdDFOWTYrb1VkaDc5bVFydHdzaVVcclxuQktVdXg5MHRpcGh6S0I0Zk9iWm1oK1pQYkJQNGhia3hCUGhHMzllMU5JeWdRZmY3Skc0YS9wNDczZz09XHJcbi0tLS0tRU5EIEVDIFBVQkxJQyBLRVktLS0tLVxyXG4iLCJhbGdvcml0aG0iOnsibmFtZSI6IkVDRFNBIiwiaGFzaCI6IlNIQS0yNTYifX0sInZlcnNpb24iOiIxMi0yMS0yMDIwIn0",
						Tags: []arweave.Tag{
							{Name: "Q29udGVudC1UeXBl", Value: "YXBwbGljYXRpb24vanNvbg"},
							{Name: "QXBwLU5hbWU", Value: "TWlycm9yWFla"},
							{Name: "Q29udHJpYnV0b3I", Value: "MHg0QzBhNDY2REYwNjI4RkU4Njk5MDUxYjNBYzY1MDY2NTMxOTFjYzIx"},
						},
					},
				},
				config: &config.Module{
					IPFSGateways: []string{"https://ipfs.rss3.page/"},
				},
			},
			want: &schema.Feed{
				ID:       "lW0AMDN2RgOeqULk-u6Tv0wfZWpx9MfkrmqQQU-Mvuo",
				Network:  filter.NetworkArweave,
				Index:    0,
				From:     "Ky1c1Kkt-jZ9sY1hvLF5nCf6WWdBhIU5Un_BMYh-t3c",
				To:       "Ky1c1Kkt-jZ9sY1hvLF5nCf6WWdBhIU5Un_BMYh-t3c",
				Type:     filter.TypeSocialPost,
				Platform: lo.ToPtr(filter.PlatformMirror),
				Fee: &schema.Fee{
					Amount:  decimal.NewFromInt(27798031),
					Decimal: 12,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialPost,
						Tag:      filter.TagSocial,
						Platform: filter.PlatformMirror.String(),
						From:     "0x4C0a466DF0628FE8699051b3Ac6506653191cc21",
						To:       "Ky1c1Kkt-jZ9sY1hvLF5nCf6WWdBhIU5Un_BMYh-t3c",
						Metadata: &metadata.SocialPost{
							Title:         "We're making a book about Eth2!",
							Body:          "Stateful Works is excited to publicly introduce our first project, a limited run print book focused on the people and teams that brought the Ethereum Beacon Chain to life.\n\n### Who is Stateful Works?\n\nWe are a group of passionate longtime community members interested in exploring and creating Ethereum culture: reach out to [@statefulworks](https://twitter.com/StatefulWorks) on Telegram or Twitter with questions.\n\n### Book Details\n\nThis book will feature profiles, reflections, and stories from Eth2 implementers and researchers. We want to share the stories of the people that made this possible. We have been working on this for the past few months, and have made significant progress towards defining and realizing the vision.\n\nSo far, we have gathered profiles from over **25 Eth2 contributors**, with many more to come. We're anticipating deliveries of the book **Genesis Edition in Q1 2021.**\n\nAs a limited run first print, the book will be a high quality collectors item. We are paying close attention to every aspect, from the overall design, to the features highlighting the uniqueness of each copy. For example, the burn transaction hash of redemption being marked on the physical print. \n\n### Why a Book?\n\nWe think the artifact that frames the work can be as important as the content itself. So much of what we consume exists ephemerally in feeds and chats: the Beacon Book captures a snapshot of sentiment, ported into object-form. This artifact will live on as something to be read, referenced, and carefully stored long into the future.\n\n### How can I get one?\n\nHere are the general details, subject to change. \n\n**NOTE: There is no way to acquire the book at this time. The date and release will be well publicized here and with our partner orgs.**\n\nThe bulk of the Genesis Edition (~100 books) will be distributed on-chain, similar to Unisocks. The main difference is that we plan to use a Balancer Smart pool to allow maximal access for the entire Ethereum community. Learn more about how these mechanisms function through case studies like the [Bankless BAP](https://bankless.substack.com/p/the-ultimate-guide-to-balancer-smart) and [PERP](https://perpetualprotocol.medium.com/why-we-chose-to-distribute-perp-using-a-balancer-liquidity-bootstrapping-pool-aac7f1ab6181). \n\nA few of the benefits include:\n\n- better price discovery: starts high > low, making it harder for whales to corner supply early on\n- fair distribution for humans (bots have no advantage)\n- efficient for bootstrapping liquidity for the pool creator\n\nA portion of the Genesis Edition (~10 books) will be pledged to partner community orgs to give away to their communities. Ethereum culture is nothing if not participatory, and we want to fully embrace this aspect. We will be announcing our partners closer to the complete launch.\n\nThere will be other fun ways to win books, stay tuned for more details! If you end up not getting a book, don't worry: digital versions of the content will be available for free to enjoy, here on this blog.\n\nFinally, an important goal of this project is to promote and support Ethereum Public Goods. Therefore, any net ETH generated by the Genesis sale net of printing, shipping and contributor costs will be donated to the [Open Grants Eth2 stream](https://opengrants.com/grant/0x53e7daa8e3aa23cd30c75b2f599c303bada17064).\n\n---\n\nYou may have noticed this post is hosted on [Mirror.xyz](https://mirror.xyz/), a new web3 protocol for publishing to communities. We're excited to be one of the first projects to publish using it! Check back next week for more info on what Stateful Works is all about - we can't wait to share more about our vision.\n\nMessage us on Telegram or Twitter [@statefulworks](https://twitter.com/StatefulWorks) with questions. ❤️",
							PublicationID: "a151ee1decb2028a8bb48277f6928c6f38319c32601dc1da1ee82acfcad2e525",
							ContentURI:    "ar://lW0AMDN2RgOeqULk-u6Tv0wfZWpx9MfkrmqQQU-Mvuo",
							Target:        nil,
							Timestamp:     1608766070,
						},
					},
				},
				Status:    true,
				Timestamp: 1608772429,
			},
			wantError: require.NoError,
		},
		{
			name: "Mirror Revise 01",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkArweave,
					Block: arweave.Block{
						Timestamp: 1609360471,
					},
					Transaction: arweave.Transaction{
						ID:       "CKsUVyAvgDeMyHjTn4kAwLMgv81nCH4xt6dY5xZcYbE",
						Reward:   "17827441",
						Quantity: "0",
						Owner:    "x3EoIMrB1qOrAivgqa1fi3SwJQgZdVuLgZhnvB0NLVngTIfB9yytGQVr9jJ4w4zn6q80T_L8aU5yXVhB12cfoABwP6AYDbvLefXLCGFgUcEcwrQ_bhWmDpxEjU2sMu4qrlI2PVj_3y1-Cx1wXldVD3i2JT81BeXacj8f19ZpKFZHaSvLM7jHm5STqwNaZDhhNnwR9qwe5kKkjafvOVDF9ayp7NHLOsEq5s8rcaQKZMYR_yEbuSfE4jzuq5F3Cwcp18bzooOqgtpCi1LozAlWepYywrxuhc-TxxObar2FVZU7bHyuNMD_aLoZ7mYqHUDGwPocqMqZVa6zQPAm7J1mnyI-fab-3E9vlZ3ZEO-PD21EAo_ALidmC3jR35tzWTGx6y7D4W1J_biJc_iKBybmsOH5W0K31CMW1sO2P2FKndUVz8vzedYvNbjYJhFFqAaEcn2x4mj15n2d75_DyEqmyEMXZ6MAg2-hUixYky7Gk7hpokwbwWEu35FsEIh4VKoanxJPgZ7JCWzt7su5su891RVf0Aodvr8UR971o1xgGaTR5ChLBC4SjyswcRTzNIb5LiPoJK4Sp03wAQW7dGgbKo8W5XjLsI-KWqypEyRkQYy-GYzwIWZ-4_5qzARaZOkVS_mn2aZzbsxxZNWxtjpJVNOd8olr0Fo88nn1xnLu2b0",
						Data:     "eyJjb250ZW50Ijp7ImJvZHkiOiJPbmUgdGhpbmcgSSd2ZSBiZWVuIHdvcmtpbmcgb24gb3ZlciB0aGUgaG9saWRheXMgaXMgdGhlIGlkZWEgb2YgYXBwbHlpbmcgW2NvaG9ydC1iYXNlZCBsZWFybmluZ10oaHR0cHM6Ly93d3cubm90aW9uLnNvL0xlYXJuLWFib3V0LVdlcy1hbmQtR2FnYW4tcy1zdGVhbHRoLXN0YXJ0dXAtYmY4YWU3ODlmZGVkNDc1M2IwZjU0YTg1Y2U1MzE1YzApIHRvIGEgY29tbXVuaXR5IG9yIHNvY2lhbCB0b2tlbi4gXG5cbiFbT3VyIGxhc3QgUGVyc29uYWwgVGFsa2VucyBtZWV0dXAsIGluIE5vdmVtYmVyXShodHRwczovL2ltYWdlcy5taXJyb3ItbWVkaWEueHl6L3B1YmxpY2F0aW9uLWltYWdlcy9hOGNmZDFhZC1iZWUyLTQwYmItODhmYi1lMWEzYmFjMDZhYjAuanBlZylcblxuSSB0aGluayB3ZSBjYW4gZG8gaXQgYnkgaG9zdGluZyB0aGUgY29ob3J0IG9uIEtpY2tiYWNrIGFuZCBnZXR0aW5nIHBhcnRpY2lwYW50cyB0byBzdGFrZSBbJHhKb29uIHRva2Vuc10oaHR0cHM6Ly9ibG9ja3Njb3V0LmNvbS9wb2EveGRhaS90b2tlbnMvMHg1ZkU5ODg1MjI2Njc3RjNFYjVDOWFkOGFCNmM0MjFCNEVBMzg1MzVkL3Rva2VuLXRyYW5zZmVycyksIHNpbWlsYXIgdG8gS2lja2JhY2sncyBmaXRuZXNzIGNoYWxsZW5nZSBpbiBEZWNlbWJlci4gXG5cbkJ1dCBhIGZldyBxdWVzdGlvbnMgcmVtYWluOlxuXG4tIFdoYXQgc2hvdWxkIHRoZSB0YXNrIGJlP1xuXG4tIEhvdyBtdWNoIHNob3VsZCBwZW9wbGUgc3Rha2U_XG5cbi0gSG93IGRvZXMgdGhpcyBmcmFtZXdvcmsgaml2ZSB3aXRoIG5vbi10b2tlbmlzZWQgY29ob3J0LWJhc2VkIGxlYXJuaW5nP1xuXG4tIENhbiB3ZSB1c2UgdG9rZW5pc2VkIG1lZGlhIHBsYXRmb3JtcyBsaWtlIE1pcnJvciBpbiB0aGlzIGV4cGVyaW1lbnQ_XG5cbkkgd291bGQgbG92ZSB0byBoZWFyIGZyb20geW91LCBETSBtZSBhdCBbQGpvb25pYW4gb24gVHdpdHRlcl0oaHR0cHM6Ly90d2l0dGVyLmNvbS9qb29uaWFuKSBvciBwaW5nIG1lIG9uIGpvb25pYW5AY3J5cHRvZ3JhcGhpYy5tZWRpYS4gIiwidGltZXN0YW1wIjoxNjA5MzYwMzkyLCJ0aXRsZSI6IkNvaG9ydC1iYXNlZCBsZWFybmluZyBmb3IgY29tbXVuaXR5IHRva2VucyIsInB1YmxpY2F0aW9uIjoiam9vbmlhbiJ9LCJkaWdlc3QiOiJhMDk0YmIxZjQzNzYyMGFjYmJiY2U2ZjIyOWYyNDk0N2FiOGM2ZmMzN2YxZGRmYTFhZTIxNjcxMmUzOGUzYmNjIiwiYXV0aG9yc2hpcCI6eyJjb250cmlidXRvciI6IjB4MzFiNEMyOTJiNDYzOWEzMkEwYUNBNzJCMDQ1MTQ5OUI0NjRjNThjYiIsInNpZ25pbmdLZXkiOiJ7XCJjcnZcIjpcIlAtMjU2XCIsXCJleHRcIjp0cnVlLFwia2V5X29wc1wiOltcInZlcmlmeVwiXSxcImt0eVwiOlwiRUNcIixcInhcIjpcIk92WlZfeEdhMDZ0bEtxQndDaWVFb2ZYM25qNmpPblNqZ3FFRzJESFFoT2dcIixcInlcIjpcIlZGUjlNOVJBX1RjbllfR1JKNUtCd2JheFE1alk0Z0NybXYydzhhS0xtbmdcIn0iLCJzaWduYXR1cmUiOiJBSHdPd1RNbEdvRmM4Q3V3b0lHcDI3V1lTZ0ZEYzAyU3VFSk9JR25ZMXNJN3FMR05pMFhaNDVMWkFHekFUOGgyMS9NVjhuSHhqZmphcTRpVXNLZ0lHQT09Iiwic2lnbmluZ0tleVNpZ25hdHVyZSI6IjB4NmI5ZmU5YjI4NzI2Y2ZkMGQ4MjVhYjViYTZkMzQxODA2MTUyZTY4MzU3YTM0MTI5NWFmNzgwN2I2Y2NlMTMwYzQzZDdhZDcyNDEyNWRmNjE2MWZhYzdiNGM2MjM0ZjJmMmVlMGRmZTcyY2ZkY2ZjZGU0YmY0ZWNiMTU3MDVhOTcxYyIsInNpZ25pbmdLZXlNZXNzYWdlIjoiSSBhdXRob3JpemUgdGhlIHB1YmxpY2F0aW9uIG9mIGFydGljbGVzIG9uIGpvb25pYW4ubWlycm9yLnh5eiBmcm9tIHRoaXMgZGV2aWNlIHVzaW5nOlxue1wiY3J2XCI6XCJQLTI1NlwiLFwiZXh0XCI6dHJ1ZSxcImtleV9vcHNcIjpbXCJ2ZXJpZnlcIl0sXCJrdHlcIjpcIkVDXCIsXCJ4XCI6XCJPdlpWX3hHYTA2dGxLcUJ3Q2llRW9mWDNuajZqT25TamdxRUcyREhRaE9nXCIsXCJ5XCI6XCJWRlI5TTlSQV9UY25ZX0dSSjVLQndiYXhRNWpZNGdDcm12Mnc4YUtMbW5nXCJ9IiwiYWxnb3JpdGhtIjp7Im5hbWUiOiJFQ0RTQSIsImhhc2giOiJTSEEtMjU2In19LCJ2ZXJzaW9uIjoiMTItMjEtMjAyMCJ9",
						Tags: []arweave.Tag{
							{Name: "Q29udGVudC1UeXBl", Value: "YXBwbGljYXRpb24vanNvbg"},
							{Name: "QXBwLU5hbWU", Value: "TWlycm9yWFla"},
							{Name: "Q29udHJpYnV0b3I", Value: "MHgzMWI0QzI5MmI0NjM5YTMyQTBhQ0E3MkIwNDUxNDk5QjQ2NGM1OGNi"},
							{Name: "Q29udGVudC1EaWdlc3Q", Value: "YTA5NGJiMWY0Mzc2MjBhY2JiYmNlNmYyMjlmMjQ5NDdhYjhjNmZjMzdmMWRkZmExYWUyMTY3MTJlMzhlM2JjYw"},
							{Name: "T3JpZ2luYWwtQ29udGVudC1EaWdlc3Q", Value: ""},
						},
					},
				},
				config: &config.Module{
					IPFSGateways: []string{"https://ipfs.rss3.page/"},
				},
			},
			want: &schema.Feed{
				ID:       "CKsUVyAvgDeMyHjTn4kAwLMgv81nCH4xt6dY5xZcYbE",
				Network:  filter.NetworkArweave,
				Index:    0,
				From:     "Ky1c1Kkt-jZ9sY1hvLF5nCf6WWdBhIU5Un_BMYh-t3c",
				To:       "Ky1c1Kkt-jZ9sY1hvLF5nCf6WWdBhIU5Un_BMYh-t3c",
				Type:     filter.TypeSocialRevise,
				Platform: lo.ToPtr(filter.PlatformMirror),
				Fee: &schema.Fee{
					Amount:  decimal.NewFromInt(17827441),
					Decimal: 12,
				},

				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialRevise,
						Tag:      filter.TagSocial,
						Platform: filter.PlatformMirror.String(),
						From:     "0x31b4C292b4639a32A0aCA72B0451499B464c58cb",
						To:       "Ky1c1Kkt-jZ9sY1hvLF5nCf6WWdBhIU5Un_BMYh-t3c",
						Metadata: &metadata.SocialPost{
							Title:         "Cohort-based learning for community tokens",
							Body:          "One thing I've been working on over the holidays is the idea of applying [cohort-based learning](https://www.notion.so/Learn-about-Wes-and-Gagan-s-stealth-startup-bf8ae789fded4753b0f54a85ce5315c0) to a community or social token. \n\n![Our last Personal Talkens meetup, in November](https://images.mirror-media.xyz/publication-images/a8cfd1ad-bee2-40bb-88fb-e1a3bac06ab0.jpeg)\n\nI think we can do it by hosting the cohort on Kickback and getting participants to stake [$xJoon tokens](https://blockscout.com/poa/xdai/tokens/0x5fE9885226677F3Eb5C9ad8aB6c421B4EA38535d/token-transfers), similar to Kickback's fitness challenge in December. \n\nBut a few questions remain:\n\n- What should the task be?\n\n- How much should people stake?\n\n- How does this framework jive with non-tokenised cohort-based learning?\n\n- Can we use tokenised media platforms like Mirror in this experiment?\n\nI would love to hear from you, DM me at [@joonian on Twitter](https://twitter.com/joonian) or ping me on joonian@cryptographic.media. ",
							PublicationID: "a094bb1f437620acbbbce6f229f24947ab8c6fc37f1ddfa1ae216712e38e3bcc",
							ContentURI:    "ar://CKsUVyAvgDeMyHjTn4kAwLMgv81nCH4xt6dY5xZcYbE",
							Timestamp:     1609360392,
						},
					},
				},
				Status:    true,
				Timestamp: 1609360471,
			},
			wantError: require.NoError,
		},
	}

	driver := database.DriverCockroachDB
	partition := true

	container, dataSourceName, err := createContainer(context.Background(), driver, partition)
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, gnomock.Stop(container))
	})

	// Dial the database.
	databaseClient, err := dialer.Dial(context.Background(), &config.Database{
		Driver:    driver,
		URI:       dataSourceName,
		Partition: &partition,
	})
	require.NoError(t, err)

	err = databaseClient.Migrate(context.Background())
	require.NoError(t, err)

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			instance, err := worker.NewWorker(testcase.arguments.config, databaseClient)
			require.NoError(t, err)

			matched, err := instance.Match(ctx, testcase.arguments.task)
			testcase.wantError(t, err)
			require.True(t, matched)

			feed, err := instance.Transform(ctx, testcase.arguments.task)
			testcase.wantError(t, err)

			//t.Log(string(lo.Must(json.MarshalIndent(feed, "", "\x20\x20"))))

			require.Equal(t, testcase.want, feed)
		})
	}
}

func createContainer(ctx context.Context, driver database.Driver, partition bool) (container *gnomock.Container, dataSourceName string, err error) {
	config := config.Database{
		Driver:    driver,
		Partition: &partition,
	}

	switch driver {
	case database.DriverCockroachDB:
		preset := cockroachdb.Preset(
			cockroachdb.WithDatabase("test"),
			cockroachdb.WithVersion("v23.1.8"),
		)

		// Use a health check function to wait for the database to be ready.
		healthcheckFunc := func(ctx context.Context, container *gnomock.Container) error {
			config.URI = formatContainerURI(container)

			client, err := dialer.Dial(ctx, &config)
			if err != nil {
				return err
			}

			transaction, err := client.Begin(ctx)
			if err != nil {
				return err
			}

			defer lo.Try(transaction.Rollback)

			return nil
		}

		container, err = gnomock.Start(preset, gnomock.WithContext(ctx), gnomock.WithHealthCheck(healthcheckFunc))
		if err != nil {
			return nil, "", err
		}

		return container, formatContainerURI(container), nil
	default:
		return nil, "", fmt.Errorf("unsupported driver: %s", driver)
	}
}

func formatContainerURI(container *gnomock.Container) string {
	return fmt.Sprintf(
		"postgres://root@%s:%d/%s?sslmode=disable",
		container.Host,
		container.DefaultPort(),
		"test",
	)
}
